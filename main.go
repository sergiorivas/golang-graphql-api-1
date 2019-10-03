package main

import (
	"log"
	"net/http"
	"test_articles/api"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type query struct{}

func (_ *query) Hello() string { return "Hello, world!12" }

func main() {
	s, _ := api.GetSchema("./api/schema.graphql")
	schema := graphql.MustParseSchema(s, &query{})
	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":3001", nil))
}
