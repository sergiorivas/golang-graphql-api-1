package main

import (
	"fmt"
	"log"
	"os"
	"test_articles/api"
	"test_articles/api/resolvers"
	"test_articles/middleware"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func connectToDB() *gorm.DB {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	return db
}

func loadSchema() string {
	s, err := api.GetSchema("./api/schema.graphql")
	if err != nil {
		panic(err)
	}

	return s
}

func graphQLHandler(r *relay.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		r.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	loadEnv()

	var db *gorm.DB
	db = connectToDB()
	defer db.Close()

	db.LogMode(!(os.Getenv("PRODUCTION") != "" && os.Getenv("PRODUCTION") == "TRUE")) //For debugging

	// Preparing schema
	rootResolver := resolvers.RootResolver{}
	schema := graphql.MustParseSchema(loadSchema(), &rootResolver)

	r := gin.New()
	r.Use(middleware.AddDatabase(db))
	r.Use(middleware.GinContextToContextMiddleware())
	r.Use(middleware.AddUser())
	r.POST("/graphql", graphQLHandler(&relay.Handler{Schema: schema}))

	r.Run(":3001")
}
