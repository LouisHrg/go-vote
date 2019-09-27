package main

import (
	"log"
	"os"
	"io"


	"github.com/gin-gonic/gin"
	"github.com/go-vote/middleware"

	"github.com/go-vote/handler"
	"github.com/asaskevich/govalidator"

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

	govalidator.SetFieldsRequiredByDefault(false)

	f, _ := os.Create("gin.log")

	gin.DefaultWriter = io.MultiWriter(f)

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(middleware.IPFirewall())

	r.POST("/login", middleware.LoginHandler)

  r.GET("/", handler.BasicResponse)

	auth := r.Group("/")
	auth.Use(middleware.JwtTokenCheck)
	{
		auth.GET("/users", handler.GetUsers)
		auth.GET("/users/:uuid", handler.GetUser)
		auth.POST("/users", handler.PostUser)
		auth.DELETE("/users/:uuid", handler.DeleteUser)
		auth.PUT("/users/:uuid", handler.PutUser)

		auth.PATCH("/users/:uuid/promote", handler.PromoteUser)

		auth.GET("/surveys", handler.GetSurveys)
		auth.GET("/surveys/:uuid", handler.GetSurvey)
		auth.POST("/surveys", handler.PostSurvey)
	}


	r.Run()
}
