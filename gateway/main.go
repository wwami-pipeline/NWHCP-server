package main

// packages used
import (
	"context"
	"crypto/tls"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"nwhcp/nwhcp-server/gateway/handlers"
	"nwhcp/nwhcp-server/gateway/models/orgs"
	"nwhcp/nwhcp-server/gateway/models/users"
	"nwhcp/nwhcp-server/gateway/sessions"
	"os"
	"strings"
	"sync/atomic"
	"time"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// executable code

// type: function which unpacks an http request
type Director func(r *http.Request)

// returns a function which unpacks an http request
// takes in a target URL, a string for a signinKey, and the sessions interface
// gets sessionState, recognizes user, sets headings, host, scheme etc.
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

// how to retrieve db env vars
// if env not found, return default option
func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func main() {
	// db
	mongoAddr := getenv("MONGO_ADDR", "mongodb://127.0.0.1:27017")
	mongoDb := getenv("MONGO_DB", "mongodb")
	mongoCol := getenv("MONGO_COL", "organization")

	// session info
	redisAddr := getenv("REDIS_ADDR", "127.0.0.1:6379")
	redisPass := getenv("REDIS_PASS", "")
	redisTls := getenv("REDIS_TLS", "")
	sess := getenv("REDIS_SESSIONKEY", "key")

	dsn := getenv("MYSQL_DSN", "root@tcp(127.0.0.1)/mydatabase")

	server2addr := getenv("SERVER2_ADDR", "http://organizations:5000")
	internalPort := getenv("INTERNAL_PORT", ":90")

	// mongodb driver boilerplate
	clientOptions := options.Client().ApplyURI(mongoAddr)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoSession, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Error dialing dbaddr: ", err)
	} else {
		fmt.Println("MongoDb Connect Success!")
	}

	orgStore, err := orgs.NewOrgStore(mongoSession, mongoDb, mongoCol)

	// auth, contex, cors, session, orgs routes
	hctx := &handlers.HandlerContext{
		OrgStore: orgStore,
	}

	// redis
	rclient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPass,
		DB:       0,
	})

	// active user? initalize in-session database cache-ing
	if len(redisTls) > 0 {
		rclient = redis.NewClient(&redis.Options{
			TLSConfig: &tls.Config{
				MinVersion:         tls.VersionTLS12,
				ServerName:         redisTls,
				InsecureSkipVerify: true,
			},
			Addr:     redisAddr,
			Password: redisPass,
			DB:       0,
		})
	}

	// error handling
	err = rclient.Set("key", "value", 0).Err()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Redis Connect Success!")
	}

	// mysql
	if len(dsn) == 0 {
		dsn = "127.0.0.1"
	}
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("MySQL Connect Success!")
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

	// get orgs links
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

	// mux no longer being maintained
	router := mux.NewRouter()

	// routes
	router.HandleFunc("/api/v1/users", handler.UsersHandler)
	router.HandleFunc("/api/v1/sessions", handler.SessionsHandler)
	router.HandleFunc("/api/v1/sessions/{id}", handler.SpecificSessionHandler)

	apiEndpoint := "/api/v2"
	router.Handle(apiEndpoint+"/orgs/{id}", orgsProxy)
	router.Handle(apiEndpoint+"/getuser/", orgsProxy)

	apiEndpoint3 := "/api/v3"
	router.HandleFunc(apiEndpoint3+"/search", hctx.SearchOrgsHandler)
	router.HandleFunc(apiEndpoint3+"/orgs", hctx.GetAllOrgs)
	router.HandleFunc(apiEndpoint3+"/orgs/{id}", hctx.SpecificOrgHandler)

	mux2 := http.NewServeMux()
	mux2.HandleFunc(apiEndpoint3+"/pipeline-db/truncate", hctx.DeleteAllOrgsHandler)
	mux2.HandleFunc(apiEndpoint3+"/pipeline-db/poporgs", hctx.InsertOrgs)
	go serve(mux2, internalPort)

	// get all data from mongodb
	router.HandleFunc(apiEndpoint3+"/orgs-all", AllDataHandler)

	addr := ":8080"
	log.Printf("server is listening at %s...", addr)

	// log.Fatal(http.ListenAndServe(addr, router))
	log.Fatal(http.ListenAndServe(addr, handlers.NewPreflight(router)))

}

func serve(router *http.ServeMux, addr string) {
	log.Fatal(http.ListenAndServe(addr, handlers.NewPreflight(router)))
	// log.Printf("server is listening at %s...", addr)
}

// serve all data from mongodb
func AllDataHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(AllData())
}

// download all data from mongodb
func AllData() []byte {
	client, err := mongo.NewClient(options.Client().ApplyURI(getenv("MONGO_ADDR", "mongodb://127.0.0.1:27017")))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	db := client.Database("mongodb")
	surveys := db.Collection("surveys")
	cursor, err := surveys.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var data []bson.M
	if err = cursor.All(ctx, &data); err != nil {
		log.Fatal(err)
	}
	jsonByte, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	return jsonByte
}
