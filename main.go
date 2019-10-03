package main

import (
	"log"
	"net/http"
	"test_articles/api"
	"test_articles/api/resolvers"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func main() {
	s, _ := api.GetSchema("./api/schema.graphql")
	schema := graphql.MustParseSchema(s, &resolvers.QueryResolver{})
	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":3001", nil))
}
