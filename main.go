package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ayushrakesh/go-rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port is not found in the environment")
	}
	fmt.Println("Port: ", portString)

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("Database not found!")
	}

	conn, er := sql.Open("postgres", dbUrl)
	fmt.Println(conn)
	if er != nil {
		log.Fatal("Can't connect to the database!")
	}

	db := database.New(conn)
	apiCfg := apiConfig{
		DB: db,
	}

	go startScraping(db, 10, time.Minute)

	router := chi.NewRouter()

	router.Use(
		cors.Handler(cors.Options{
			AllowedOrigins: []string{"http://*", "https://*"},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders: []string{"*"},
			ExposedHeaders: []string{"Link"},
			// AllowedCredentials: false,
			MaxAge: 300,
		}),
	)

	v1Router := chi.NewRouter()

	v1Router.Get("/readiness", handlerReadiness)
	v1Router.Get("/error", handlerError)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.getUserByAPIKey)

	v1Router.Post("/feeds", apiCfg.handlerCreateFeed)
	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)

	v1Router.Post("/feed_follows", apiCfg.handlerCreateFeedFollow)
	v1Router.Get("/feed_follows", apiCfg.handlerGetFeedFollow)
	v1Router.Delete("/feed_follows/{feedFollowID}", apiCfg.handlerDeleteFeedFollow)

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Running on port: ", portString)
}
