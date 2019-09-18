package main

import (
	"log"
	"os"
	"io"

	"github.com/gin-gonic/gin"

	"github.com/go-vote/middleware"
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

	gin.DisableConsoleColor()

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	authware, err := middleware.AuthMiddleware()

	if err != nil {
	  log.Fatal("JWT Error:" + err.Error())
	}

	r.POST("/login", authware.LoginHandler)

  r.GET("/", handler.BasicResponse)
	r.GET("/users", handler.GetUsers)
	r.GET("/users/:uuid", handler.GetUser)
	r.POST("/users", handler.PostUser)

	r.GET("/surveys", handler.GetSurveys)
	r.GET("/surveys/:uuid", handler.GetSurvey)

	r.Run()
}
