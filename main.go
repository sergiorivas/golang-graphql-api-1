package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"test_articles/api"
	"test_articles/api/resolvers"

	"github.com/joho/godotenv"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)
	db, err := gorm.Open("postgres", connectionString)
	defer db.Close()

	db.LogMode(!(os.Getenv("PRODUCTIO") != "" && os.Getenv("PRODUCTION") == "TRUE")) //For debugging

	if err != nil {
		panic(err)
	}

	s, err := api.GetSchema("./api/schema.graphql")
	if err != nil {
		panic(err)
	}

	rootResolver := resolvers.QueryResolver{DB: db}
	schema := graphql.MustParseSchema(s, &rootResolver)
	http.Handle("/graphql", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":3001", nil))

}
