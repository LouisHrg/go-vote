package main

import (
	"log"
	"os"
	"io"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/go-vote/handler"
	"github.com/go-vote/middleware"
	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load()
	if err != nil {
	  log.Printf("Error loading .env file")
	}

	var port = "8080"
	if len(port) == 0 {
		log.Panic("no given port")
	}
}

// Init libs, loggerS...
func initLibs() {

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	gin.DisableConsoleColor()

	govalidator.SetFieldsRequiredByDefault(false)

}

func initRouter() *gin.Engine {

	r := gin.Default()

	r.Use(gin.Recovery())

	r.Use(gin.Logger())

	r.Use(middleware.IPFirewall())

	r.POST("/login", middleware.LoginHandler)
    r.GET("/", handler.BasicResponse)

	auth := r.Group("/")
	auth.Use(middleware.JwtTokenCheck)
	{
		auth.GET("/users", handler.GetUsers)
		auth.GET("/users/:uuid", handler.GetUser)

		auth.GET("/surveys", handler.GetSurveys)
		auth.GET("/surveys/:uuid", handler.GetSurvey)
		auth.POST("/surveys", handler.PostSurvey)

		// Admin only protected routes
		admin := auth.Group("/")
		admin.Use(middleware.ACLCheck)

		admin.POST("/users", handler.PostUser)
		admin.PUT("/users/:uuid", handler.PutUser)

		admin.DELETE("/users/:uuid", handler.DeleteUser)
		admin.PATCH("/users/:uuid/promote", handler.PromoteUser)

	}

	return r
}

func main() {

	r := initRouter()

	r.Run()
}
