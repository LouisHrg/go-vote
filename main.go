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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}


func initRouter() *gin.Engine {

	r := gin.Default()

	r.Use(CORSMiddleware())

	r.Use(gin.Recovery())

	r.Use(gin.Logger())

	r.Use(middleware.IPFirewall())

	r.POST("/login", middleware.LoginHandler)
    r.GET("/", handler.BasicResponse)
	r.POST("/users", handler.PostUser)
	auth := r.Group("/")
	auth.Use(middleware.JwtTokenCheck)
	{
		auth.GET("/users", handler.GetUsers)
		auth.GET("/users/:uuid", handler.GetUser)

		auth.GET("/login/me", handler.GetCurrentUser)

		auth.PATCH("/responses/:uuid/vote", handler.Vote)
		auth.DELETE("/responses/:uuid/vote", handler.DeleteVote)

		auth.GET("/surveys", handler.GetSurveys)
		auth.GET("/surveys/:uuid", handler.GetSurvey)
		auth.POST("/surveys", handler.PostSurvey)

		auth.PATCH("/responses/:uuid", handler.PatchResponse)

		// Admin only protected routes
		admin := auth.Group("/")
		admin.Use(middleware.ACLCheck)

		
		admin.PATCH("/users/:uuid", handler.PatchUser)

		admin.DELETE("/users/:uuid", handler.DeleteUser)
		admin.PATCH("/users/:uuid/promote", handler.PromoteUser)

		admin.DELETE("/surveys/:uuid", handler.DeleteSurvey)
	}

	return r
}

func main() {

	r := initRouter()

	r.Run()
}
