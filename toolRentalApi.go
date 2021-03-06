package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getHello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// TODO: Implement a DI container for route handlers
// TODO: Implement an ORM
// TODO: implement authentication
// TODO: Implement Query Layer
// TODO: Implement Command Layer
// TODO: Implement GraphQL
// TODO: Implement GRPC

func main() {
	fmt.Println("Hello, World!")
	router := gin.Default()
	router.GET("/hello", getHello)
}
