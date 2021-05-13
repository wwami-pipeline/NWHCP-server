package main

import (
	"NWHCP/NWHCP-server/gateway/handlers"
	"NWHCP/NWHCP-server/gateway/models/users"
	"NWHCP/NWHCP-server/gateway/sessions"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
)

type Director func(r *http.Request)

func CustomDirector2(target *url.URL) Director {
	return func(r *http.Request) {
		r.Header.Add("X-Forwarded-Host", r.Host)
		r.Host = target.Host
		r.URL.Host = target.Host
		r.URL.Scheme = target.Scheme
	}
}

func main() {
	addr := os.Getenv("ADDR")
	cert := os.Getenv("TLSCERT")
	key := os.Getenv("TLSKEY")
	sess := os.Getenv("SESSIONKEY")
	redisAddr := os.Getenv("REDISADDR")
	SUMMARYADDR := os.Getenv("SUMMARYADDR")
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

	mux := http.NewServeMux()

	// orgsDirector := func(r *http.Request) {
	// 	addresses := strings.Split(orgsAddr, ", ")
	// 	serv := addresses[0]
	// 	if len(addresses) > 1 {
	// 		rand.Seed(time.Now().UnixNano())
	// 		serv = addresses[rand.Intn(len(addresses))]
	// 	}
	// 	r.Header.Del("X-User")
	// 	state := &handlers.SessionState{}
	// 	sid, _ := sessions.GetSessionID(r, handler.SessionKey)
	// 	err := handler.SessionStore.Get(sid, &state)
	// 	if err == nil {
	// 		json, _ := json.Marshal(state.User)
	// 		r.Header.Set("X-User", string(json))
	// 	}
	// 	r.Host = serv
	// 	r.URL.Host = serv
	// 	r.URL.Scheme = "http"
	// }

	// meetingProxy := &httputil.ReverseProxy{Director: meetingDirector}
	// orgsProxy := &httputil.ReverseProxy{Director: orgsDirector}
	summaryURL := &url.URL{Scheme: "http", Host: SUMMARYADDR}
	summaryProxy := &httputil.ReverseProxy{Director: CustomDirector2(summaryURL)}

	mux.HandleFunc("/api/v1/users", handler.UsersHandler)
	mux.HandleFunc("/api/v1/sessions", handler.SessionsHandler)
	mux.HandleFunc("/api/v1/getuser/", handler.GetUserInfoHandler)
	mux.Handle("/api/v1/summary", summaryProxy)
	// mux.Handle("/meeting", meetingProxy)
	// mux.Handle("/meeting/", meetingProxy)
	// mux.Handle("/user/", meetingProxy)
	// mux.Handle("/orgs", orgsProxy)
	// mux.Handle("/orgs/", orgsProxy)
	// mux.Handle("/search", orgsProxy)

	newMux := handlers.NewPreflight(mux)
	log.Printf("server is listening at http://%s", addr)
	log.Fatal(http.ListenAndServeTLS(addr, cert, key, newMux))
}
