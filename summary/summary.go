package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

//PreviewImage represents a preview image for a page
type PreviewImage struct {
	URL       string `json:"url,omitempty"`
	SecureURL string `json:"secureURL,omitempty"`
	Type      string `json:"type,omitempty"`
	Width     int    `json:"width,omitempty"`
	Height    int    `json:"height,omitempty"`
	Alt       string `json:"alt,omitempty"`
}

//PageSummary represents summary properties for a web page
type PageSummary struct {
	Type        string          `json:"type,omitempty"`
	URL         string          `json:"url,omitempty"`
	Title       string          `json:"title,omitempty"`
	SiteName    string          `json:"siteName,omitempty"`
	Description string          `json:"description,omitempty"`
	Author      string          `json:"author,omitempty"`
	Keywords    []string        `json:"keywords,omitempty"`
	Icon        *PreviewImage   `json:"icon,omitempty"`
	Images      []*PreviewImage `json:"images,omitempty"`
}

const headerCORS = "Access-Control-Allow-Origin"

//SummaryHandler handles requests for the page summary API.
//This API expects one query string parameter named `url`,
//which should contain a URL to a web page. It responds with
//a JSON-encoded PageSummary struct containing the page summary
//meta-data.
func SummaryHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
	w.Header().Add(headerCORS, "*")
	URLURL := r.URL.Query().Get("url")
	err := errors.New("StatusBadRequest")
	if len(URLURL) == 0 {
		fmt.Println(err)
	}

	fetched, fetchErr := fetchHTML(URLURL)
	extracted, exErr := extractSummary(URLURL, fetched)

	if fetchErr != nil {
		fmt.Printf("Error in fetching %v\n", fetchErr)
	}

	if exErr != nil {
		fmt.Printf("Error in extracting summary %v\n", exErr)
	}

	defer fetched.Close()

	enc := json.NewEncoder(w)
	w.Header().Add("Content-Type", "application/json")

	if err2 := enc.Encode(extracted); err2 != nil {
		fmt.Printf("error encoding struct into JSON%v\n", err2)
	}

}

//fetchHTML fetches `pageURL` and returns the body stream or an error.
//Errors are returned if the response status code is an error (>=400),
//or if the content type indicates the URL is not an HTML page.
// PASSED
func fetchHTML(pageURL string) (io.ReadCloser, error) {
	pgURLResp, err := http.Get(pageURL)

	ctype := pgURLResp.Header.Get("Content-Type")

	if pgURLResp.StatusCode >= 400 {
		nfErr := errors.New("Not Found URL")
		return nil, nfErr
	}

	if !strings.HasPrefix(ctype, "text/html") {
		return nil, errors.New("Non-HTMLkkkk URL")
	}

	return pgURLResp.Body, err
}

//extractSummary tokenizes the `htmlStream` and populates a PageSummary
//struct with the page's summary meta-data.
func extractSummary(pageURL string, htmlStream io.ReadCloser) (*PageSummary, error) {

	tokenizer := html.NewTokenizer(htmlStream)
	ps := &PageSummary{}
	var imgs []*PreviewImage
	imgInd := -1
	for {
		next := tokenizer.Next()
		curToken := tokenizer.Token()
		if next == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				break
			}
			log.Fatalf("error tokenizing HTML: %v", tokenizer.Err())
		} else if (next == html.StartTagToken || next == html.SelfClosingTagToken) && curToken.Data == "meta" && len(curToken.Attr) == 2 {
			attrs := curToken.Attr
			field, content := "", ""
			if attrs[0].Key != "content" {
				field = attrs[0].Val
				content = attrs[1].Val
			} else {
				field = attrs[1].Val
				content = attrs[0].Val
			}
			if field == "og:type" {
				ps.Type = content
			} else if field == "og:url" {
				ps.URL = content
			} else if field == "og:title" {
				ps.Title = content
			} else if field == "og:site_name" {
				ps.SiteName = content
			} else if (field == "og:description") || (field == "description" && ps.Description == "") {
				ps.Description = content
			} else if field == "author" {
				ps.Author = content
			} else if field == "keywords" {
				kw := strings.Split(content, ",")
				for i, word := range kw {
					kw[i] = strings.TrimSpace(word)
				}
				ps.Keywords = kw
			} else if field == "og:image" {
				imgInd++
				imgs = append(imgs, &PreviewImage{URL: content})
			} else if strings.HasPrefix(field, "og:image") {
				if field == "og:image:secure_url" {
					imgs[imgInd].SecureURL = content
				} else if field == "og:image:type" {
					imgs[imgInd].Type = content
				} else if field == "og:image:width" {
					imgs[imgInd].Width, _ = strconv.Atoi(content)
				} else if field == "og:image:height" {
					imgs[imgInd].Height, _ = strconv.Atoi(content)
				} else if field == "og:image:alt" {
					imgs[imgInd].Alt = content
				}
			}
		} else if next == html.EndTagToken && tokenizer.Token().Data == "head" {
			break
		} else if (next == html.StartTagToken || next == html.SelfClosingTagToken) && curToken.Data == "link" {
			iconAttrs := curToken.Attr
			ps.Icon = &PreviewImage{}
			for i := 1; i < len(iconAttrs); i++ {
				cur := iconAttrs[i].Key
				if cur == "href" {
					curLink := iconAttrs[i].Val
					if strings.HasPrefix(curLink, "/") {
						curLink = strings.TrimSuffix(pageURL, "/test.html") + curLink
					}
					ps.Icon.URL = curLink
				} else if cur == "type" {
					ps.Icon.Type = iconAttrs[i].Val
				} else if cur == "sizes" && iconAttrs[i].Val != "any" {
					px := strings.Split(iconAttrs[i].Val, "x")
					ps.Icon.Height, _ = strconv.Atoi(px[0])
					ps.Icon.Width, _ = strconv.Atoi(px[1])
				}
			}
		} else if next == html.StartTagToken && curToken.Data == "title" && ps.Title == "" {
			tokenizer.Next()
			titleVal := tokenizer.Token()
			ps.Title = titleVal.Data
		}

	}

	for i := 0; i < len(imgs); i++ {
		igURL := imgs[i].URL
		if strings.HasPrefix(igURL, "/") {
			igURL = strings.TrimSuffix(pageURL, "/test.html") + igURL
		}
		imgs[i].URL = igURL
	}
	ps.Images = imgs
	return ps, nil
}
