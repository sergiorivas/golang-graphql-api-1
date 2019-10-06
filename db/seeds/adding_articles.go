package main

import (
	"fmt"
	"log"
	"os"
	"test_articles/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
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

func main() {
	loadEnv()

	var db *gorm.DB
	db = connectToDB()
	populate(db)
	defer db.Close()

	db.LogMode(!(os.Getenv("PRODUCTIO") != "" && os.Getenv("PRODUCTION") == "TRUE")) //For debugging
}

func populate(db *gorm.DB) {
	a1 := models.Article{
		Title: "Title 1",
	}
	db.FirstOrCreate(&a1, a1)

	a2 := models.Article{
		Title: "Title 2",
	}
	db.FirstOrCreate(&a2, a2)

	c1 := models.Comment{
		Content:   "comment a",
		ArticleID: a1.ID,
	}
	db.FirstOrCreate(&c1, c1)

	c2 := models.Comment{
		Content:   "comment b",
		ArticleID: a2.ID,
	}
	db.FirstOrCreate(&c2, c2)

	c3 := models.Comment{
		Content:   "comment c",
		ArticleID: a1.ID,
	}
	db.FirstOrCreate(&c3, c3)

}
