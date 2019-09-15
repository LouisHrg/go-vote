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

	r.Run()
}

func TestEndpoint(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "SALUT",
		})
}
