package main

import (
	"fmt"
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/pustaka?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db connection error")
	}

	fmt.Println("Database connecction succeed")

	db.AutoMigrate(&book.Book{})

	// book := book.Book{}
	// book.Title = "Man Tiger"
	// book.Price = 90000
	// book.Rating = 5
	// book.Description = "This is a book about tiger-like man"

	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("Error creating book record")
	// }

	// Fetch data
	var books []book.Book
	err = db.Debug().First(&books, 1).Error
	if err != nil {
		fmt.Println("Error creating book record")
	}
	for _, v := range books {
		fmt.Println("Title : ", v.Title)
	}

	router := gin.Default()
	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id", handler.BooksHandler)
	v1.GET("query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	router.Run()
}
