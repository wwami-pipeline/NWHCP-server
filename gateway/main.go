package main

import (
	"NWHCP-server/gateway/handlers"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"nwhcp/nwhcp-server/gateway/handlers"
	"os"
	"strings"
	"sync/atomic"
	"time"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Director func(r *http.Request)

func CustomDirector(target []*url.URL, signingKey string, sessionStore handlers.SessionStore) Director {
	var counter int32 = 0
	return func(r *http.Request) {
		log.Println("hello")
		targ := target[counter%int32(len(target))]
		atomic.AddInt32(&counter, 1)
		state := &handlers.SessionState{}
		_, err := handlers.GetState(r, signingKey, sessionStore, state)
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

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func main() {
	// start mongo server to get organization information
	mongoAddr := getenv("MONGO_ADDR", "mongodb://127.0.0.1:27017")
	// mongoDb := getenv("MONGO_DB", "mongodb")
	// mongoCol := getenv("MONGO_COL", "organization")

	// start Redis server to get cache
	redisAddr := getenv("REDIS_ADDR", "127.0.0.1:6379")
	redisPass := getenv("REDIS_PASS", "")
	redisTls := getenv("REDIS_TLS", "")
	// sess := getenv("REDIS_SESSIONKEY", "key")

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

	// orgStore, _ := orgs.NewOrgStore(mongoSession, mongoDb, mongoCol)

	// org details
	// hctx := &handlers.HandlerContext{
	// 	OrgStore: orgStore,
	// }

	// redis
	rclient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPass,
		DB:       0,
	})

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

	err = rclient.Set("key", "value", 0).Err()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Redis Connect Success!")
	}
	// are we using this with mongo?
	// mysql
	if len(dsn) == 0 {
		dsn = "127.0.0.1"
	}
	// db, err := sql.Open("mysql", dsn)
	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	fmt.Println("MySQL Connect Success!")
	// }

	// dur, err2 := time.ParseDuration("24h")
	// if err2 != nil {
	// 	log.Fatal(err)
	// }

	// handler for redis cache?
	// handler := handlers.Handler{
	// 	SessionKey: sess,
	// 	// ThisSessionStore: handlers.NewRedisStore(rclient, dur),
	// 	// UserStore:    users.GetNewStore(db),
	// }

	// get URLS for orgs?
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
	// orgsProxy := &httputil.ReverseProxy{Director: CustomDirector(oUrls, handler.SessionKey, handler.SessionStore)}
	// mux router
	router := mux.NewRouter()

	// not in use
	// router.HandleFunc("/api/v1/users", handler.UsersHandler)
	// router.HandleFunc("/api/v1/sessions", handler.SessionsHandler)
	// router.HandleFunc("/api/v1/sessions/{id}", handler.SpecificSessionHandler)

	// not in use
	// apiEndpoint := "/api/v2"
	// router.Handle(apiEndpoint+"/orgs/{id}", orgsProxy)
	// authorization?
	// router.Handle(apiEndpoint+"/getuser/", orgsProxy)

	// not in use -routes implemented but not connected to mongoDB
	apiEndpoint3 := "/api/v3"
	// oc := handlers.NewOrganizationController(mongoSession)
	// router.HandleFunc(apiEndpoint3+"/search", hctx.GetOrgByID)
	// router.HandleFunc(apiEndpoint3+"/orgs", hctx.GetAllOrgs)
	// router.HandleFunc("/orgs", oc.CreateOrganization)
	// router.HandleFunc("/orgs/{id}", oc.GetOrgByID)

	// users
	uc := handlers.NewUserController(mongoSession)
	oc := handlers.NewOrganizationController(mongoSession)
	router.HandleFunc("/allUsers", uc.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", uc.GetUserByID).Methods("GET")
	router.HandleFunc("/users", uc.CreateUser).Methods("POST")
	router.HandleFunc("/deleteUsers/{id}", uc.DeleteUserByID).Methods("DELETE")

	// Routes for User Orgs --Need to refactor; UserRouter and UserHandler funcs
	router.HandleFunc("/users/{id}/orgs/{orgsid}/favoritedOrgs", uc.AddOrgToFavorite).Methods("PUT")
	router.HandleFunc("/users/{id}/deleteOrg/{orgsid}/favoritedOrgs", uc.DeleteOrgFavorite).Methods("DELETE")
	router.HandleFunc("/users/{id}/orgs/{orgId}/pathwayPrograms", uc.AddToPathwayOrganizations).Methods("PUT")
	router.HandleFunc("/users/{id}/deleteOrg/{orgId}/pathwayOrganizations", uc.DeleteFromPathwayOrganizations).Methods("DELETE")
	router.HandleFunc("/users/{id}/orgs/{orgId}/academicOrganizations", uc.AddToAcademicOrganizations).Methods("PUT")
	router.HandleFunc("/users/{id}/deleteOrg/{orgId}/academicOrganizations", uc.DeleteFromAcademicOrganizations).Methods("DELETE")
	router.HandleFunc("/users/{id}/orgs/{orgId}/completedOrganizations", uc.AddToCompletedOrganizations).Methods("PUT")
	router.HandleFunc("/users/{id}/deleteOrg/{orgId}/completedOrganizations", uc.DeleteFromCompletedOrganizations).Methods("DELETE")
	router.HandleFunc("/users/{id}/orgs/{orgId}/inProgressOrganizations", uc.AddToinProgressOrganizations).Methods("PUT")
	router.HandleFunc("/users/{id}/deleteOrg/{orgId}/inProgressOrganizations", uc.DeleteFrominProgressOrganizations).Methods("DELETE")

	// beginning routes refactor and Notes resource handlers 03/06

	// reduce notes function
	// remove trailing slash endpoint after development of user accounts
	http.HandleFunc("/notes/", notesRouter)
	http.HandleFunc("/notes", notesRouter)

	// Routes for User Notes

	// router.HandleFunc("/notes", uc.CreateNote).Methods("POST")
	router.HandleFunc("users/{userId}/notes/{noteId}/orgs/{orgId}/notes/{noteId}", uc.AddNoteID).Methods("PUT")
	// router.HandleFunc("/allNotes", uc.GetNotes).Methods("GET")
	// router.HandleFunc("/notes/{id}", uc.GetNoteByID).Methods("GET")
	// router.HandleFunc("/deleteNote/{id}", uc.DeleteNoteByID).Methods("DELETE")

	// organizations
	router.HandleFunc("/organizations", oc.CreateOrganization)
	router.HandleFunc("/organizations/{id}", oc.GetOrgByID)
	router.HandleFunc("/allOrganizations", oc.GetOrganizations)
	router.HandleFunc("/deleteOrg/{id}", oc.DeleteOrganizationByID)

	// not sure
	mux2 := http.NewServeMux()
	// mux2.HandleFunc(apiEndpoint3+"/pipeline-db/truncate", hctx.DeleteAllOrgsHandler)
	// mux2.HandleFunc(apiEndpoint3+"/pipeline-db/poporgs", hctx.InsertOrgs)
	go serve(mux2, internalPort)

	// get all data from mongodb
	// in use
	router.HandleFunc(apiEndpoint3+"/orgs-all", AllDataHandler)

	addr := ":8080"
	log.Printf("server is listening at %s...", addr)

	// log.Fatal(http.ListenAndServe(addr, router))
	log.Fatal(http.ListenAndServe(addr, handlers.NewPreflight(router)))

}

// move to handlers
func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("asset not found\n"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Running NWHCP API v4\n"))
}

// -Notes Router
func notesRouter(w http.ResponseWriter, r *http.Request) {

	// start mongo server to get organization information
	mongoAddr := getenv("MONGO_ADDR", "mongodb://127.0.0.1:27017")

	// start Redis server to get cache
	redisAddr := getenv("REDIS_ADDR", "127.0.0.1:6379")
	redisPass := getenv("REDIS_PASS", "")
	redisTls := getenv("REDIS_TLS", "")

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

	// orgStore, _ := orgs.NewOrgStore(mongoSession, mongoDb, mongoCol)

	// org details
	// hctx := &handlers.HandlerContext{
	// 	OrgStore: orgStore,
	// }

	// redis
	rclient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPass,
		DB:       0,
	})

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

	err = rclient.Set("key", "value", 0).Err()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Redis Connect Success!")
	}

	uc := handlers.NewUserController(mongoSession)
	// default request path; trim '/' at the end
	path := strings.TrimSuffix(r.URL.Path, "/")

	// compare path to '/notes/'
	if path == "/notes" {
		switch r.Method {
		case http.MethodPost:
			uc.CreateNote(r, n)
			return
		default:
			handlers.PostError(w, http.StatusMethodNotAllowed)
		}
	}

	// grab resource ID from params
	path = strings.TrimPrefix(path, "/users/")

	if !primitive.IsValidObjectID(path) {
		handlers.PostError(w, http.StatusNotFound)
		return
	}

	// parse ID into separate variable
	id := primitive.ObjectIDFromHex(path)

	if path == "/notes/" {
		switch r.Method {
		case http.MethodPost:
			CreateNote(w, r, id)
			return
		default:
			handlers.PostError(w, http.StatusMethodNotAllowed)
		}
	}

}

// are we using this? or code on line 189?
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

// used in AllDataHandler
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
