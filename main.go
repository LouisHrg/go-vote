package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-vote/handler"
)

func init() {
	var port = "8080"
	if len(port) == 0 {
		log.Panic("no given port")
	}
}

func main() {

	r := gin.Default()

	r.GET("/", TestEndpoint)
	r.GET("/users", handler.GetUsers)
	r.GET("/users/:id", handler.GetUser)
	//r.POST("/users", handler.PostUser)

	r.GET("/surveys", handler.GetSurveys)
	r.GET("/surveys/:id", handler.GetSurvey)

	r.Run()
}

// TestEndpoint les potes
func TestEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"response": "Welcome to go-vote API",
	})
}
