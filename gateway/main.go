package main

import (
	// "NWHCP/NWHCP-server/gateway/handlers"
	// "NWHCP/NWHCP-server/gateway/models/users"
	// "NWHCP/NWHCP-server/gateway/sessions"

	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"nwhcp/nwhcp-server/gateway/handlers"
	"nwhcp/nwhcp-server/gateway/models/users"
	"nwhcp/nwhcp-server/gateway/sessions"
	"os"
	"strings"
	"sync/atomic"
	"time"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	// "github.com/nwhcp-server/gateway/handlers"
	// "github.com/nwhcp-server/gateway/models/users"
	// "github.com/nwhcp-server/gateway/sessions"
)

type Director func(r *http.Request)

func CustomDirector(target []*url.URL, signingKey string, sessionStore sessions.Store) Director {
	var counter int32 = 0
	return func(r *http.Request) {
		log.Println("hello")
		targ := target[counter%int32(len(target))]
		atomic.AddInt32(&counter, 1)
		state := &handlers.SessionState{}
		_, err := sessions.GetState(r, signingKey, sessionStore, state)
		if err != nil {
			r.Header.Del("X-User")
			log.Printf("Error getting state: %v", err)
			return
		}
		user, _ := json.Marshal(state.User)
		r.Header.Add("X-User", string(user))
		r.Header.Set("X-User", string(user))
		r.Host = targ.Host
		r.URL.Host = targ.Host
		r.URL.Scheme = targ.Scheme
	}
}

func main() {
	addr := os.Getenv("ADDR")
	cert := os.Getenv("TLSCERT")
	key := os.Getenv("TLSKEY")
	sess := os.Getenv("SESSIONKEY")
	redisAddr := os.Getenv("REDISADDR")
	server2addr := os.Getenv("SERVER2ADDR")
	// meetingAddr := os.Getenv("MEETINGADDR")
	// orgsAddr := os.Getenv("ORGSADDR")
	dsn := os.Getenv("DSN")

	if len(addr) == 0 {
		addr = ":443"
	}

	if len(cert) == 0 || len(key) == 0 {
		fmt.Fprintln(os.Stderr, "Either the key or certificate was not found")
		os.Exit(1)
	}

	rclient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	dur, err2 := time.ParseDuration("24h")
	if err2 != nil {
		log.Fatal(err)
	}

	handler := handlers.Handler{
		SessionKey:   sess,
		SessionStore: sessions.NewRedisStore(rclient, dur),
		UserStore:    users.GetNewStore(db),
	}

	orgsAddress := strings.Split(server2addr, ",")
	var oUrls []*url.URL
	for _, cur := range orgsAddress {
		curURL, err := url.Parse(cur)
		if err != nil {
			fmt.Printf("Error parsing URL addr: %v", err)
		}
		oUrls = append(oUrls, curURL)
	}

	// meetingProxy := &httputil.ReverseProxy{Director: meetingDirector}
	// orgsProxy := &httputil.ReverseProxy{Director: orgsDirector}
	orgsProxy := &httputil.ReverseProxy{Director: CustomDirector(oUrls, handler.SessionKey, handler.SessionStore)}
	mux := mux.NewRouter()

	mux.HandleFunc("/api/v1/users", handler.UsersHandler)
	mux.HandleFunc("/api/v1/sessions", handler.SessionsHandler)
	mux.HandleFunc("/api/v1/getuser/", handler.GetUserInfoHandler)

	apiEndpoint := "/api/v2"
	// mux.Handle("/api2/v1/orgs", orgsProxy)
	mux.Handle(apiEndpoint+"/orgs/{id}", orgsProxy)
	// mux.Handle("/api2/v1/search", orgsProxy)
	mux.Handle(apiEndpoint+"/getuser/", orgsProxy)

	newMux := handlers.NewPreflight(mux)
	log.Printf("server is listening at http://%s", addr)
	log.Fatal(http.ListenAndServeTLS(addr, cert, key, newMux))
}
