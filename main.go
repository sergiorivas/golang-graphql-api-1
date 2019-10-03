package main

import (
	"log"
	"net/http"
	"test_articles/api"
	"test_articles/api/resolvers"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=srivas dbname=test sslmode=disable password=")
	if err != nil {
		panic(err)
	}

	s, err := api.GetSchema("./api/schema.graphql")
	if err != nil {
		panic(err)
	}

	rootResolver := resolvers.QueryResolver{DB: db}
	schema := graphql.MustParseSchema(s, &rootResolver)
	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":3001", nil))

}
