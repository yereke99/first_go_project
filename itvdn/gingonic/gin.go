package gingonic

import (
	"fmt"
	"net/http"
	"time"
	"log"
	"errors"
	"os"
	"os/signal"
	"context"
	"syscall"
	"github.com/gin-gonic/gin"
)

func Run() {

	handler := gin.New()

	handler.Use(gin.BasicAuth(gin.Accounts{
		"test": "test",
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}), gin.Logger(), gin.Recovery(), Logger())

	book := handler.Group("/book")
	{
		book.GET("/:id", HandleGetBook)
		book.POST("/", HandleAddBook)
		book.PUT("/:id", HandleUpdateBook)
		book.DELETE("/:id", HandleDeleteBook)
	}

	handler.GET("/books/", HandleGetBooks)

	s := &http.Server{
		Addr:           ":8082",
		Handler:        handler,          // if nil use default http.DefaultServeMux
		ReadTimeout:    10 * time.Second, // max duration reading entire request
		WriteTimeout:   10 * time.Second, // max timing write response
		IdleTimeout:    15 * time.Second, // max time wait for the next request
		MaxHeaderBytes: 1 << 20,          // 2^20 or 128kbytes
	}

	go func() {
		log.Printf("Listening on http://%s\n", s.Addr)
		log.Fatal(s.ListenAndServe())
	}()

	graceful(s, 5 * time.Second)
}

func graceful(hs *http.Server, timeout time.Duration) {
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	log.Printf("\nShutdown with timeout: %s\n", timeout)

	if err := hs.Shutdown(ctx); err != nil {
		log.Printf("Error: %v\n", err)
	} else {
		log.Println("Server stopped")
	}
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context){
		c.Header("Content-Type", "application/json")

		log.Printf("\ncustom loger server [gin] method [%s]  connection from [%v]\n", c.Request.Method, c.Request.RemoteAddr)

		c.Next()
	}
}

func HandleGetBooks(c *gin.Context) {
	books := bookStore.GetBooks()

	c.JSON(http.StatusOK, books)
}

func HandleGetBook(c *gin.Context) {
	bookid := c.Params.ByName("id")

	book := bookStore.FindBookById(bookid)

	if book == nil {
		error := fmt.Sprintf("Book with id %s not found", bookid)

		c.JSON(http.StatusNotFound, error)

		return
	}

	c.JSON(http.StatusOK, book)

}

func HandleAddBook(c *gin.Context) {

	var book Book

	err := c.BindJSON(&book)

	if err != nil {
		error := fmt.Sprintf("Bad request. %v", err)
		c.JSON(http.StatusBadRequest, error)
		return
	}


	err = bookStore.AddBook(book)
	if err != nil {
		error := fmt.Sprintf("Bad request. %v", err)
		c.JSON(http.StatusBadRequest, error)
		return
	}

	HandleGetBooks(c)
}

func HandleUpdateBook(c *gin.Context) {
	bookid := c.Params.ByName("id")

	var book Book

	err := c.BindJSON(&book)
	if err != nil {
		error := fmt.Sprintf("Bad request. %v", err)
		c.JSON(http.StatusBadRequest, error)
		return
	}

	book.Id = bookid

	err = bookStore.SetBook(book)

	if err != nil {
		error := fmt.Sprintf("Bad request. %v", err)
		c.JSON(http.StatusNotFound, error)
		return
	}

	HandleGetBook(c)
}

func HandleDeleteBook(c *gin.Context) {

	bookid := c.Params.ByName("id")
	err := bookStore.DelBook(bookid)

	if err != nil {
		error := fmt.Sprintf("%v", err)
		c.JSON(http.StatusNotFound, error)

		return
	}

	HandleGetBooks(c)
}

// BOOK

type Book struct {
	Id     string `json:"id"`
	Author string `json:"author"`
	Name   string `json:"name"`
}

type BookStore struct {
	books []Book
}

var bookStore = BookStore{
	books: make([]Book, 0),
}

func (s BookStore) GetBooks() []Book {
	return s.books
}

func (s BookStore) FindBookById(id string) *Book {
	for _, book := range s.books {
		if book.Id == id {
			return &book
		}
	}

	return nil
}

func (s *BookStore) AddBook(book Book) error{
	bk := s.FindBookById(book.Id)
	if bk != nil {
		return errors.New(fmt.Sprintf("Book with id %s already exists", book.Id))

	}
	s.books = append(s.books, book)

	return nil
}


func (s *BookStore) SetBook(book Book) error{
	for i, bk := range s.books {
		if bk.Id == book.Id {

			s.books[i] = book

			return nil
		}
	}

	return errors.New(fmt.Sprintf("There is no book with id %s", book.Id))
}

func (s *BookStore) DelBook(id string) error {
	for i, bk := range s.books {
		if bk.Id == id {

			s.books = append(s.books[:i], s.books[i+1:]...)

			return nil
		}
	}

	return errors.New(fmt.Sprintf("There is no book with id %s", id))
}