package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-vote/provider"
	"strconv"
	"fmt"
)

const pageDefault string = "1"
const limitDefault string = "10"

var db = provider.GetDB()

func getAllAbstract(c *gin.Context, objects interface{}) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", pageDefault))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", limitDefault))

	usersCol := provider.GetAll(objects, page, limit)

	c.JSON(200, usersCol)
}

func getAbstract(c *gin.Context, object interface{}) {

	uuid := c.Param("uuid")
	fmt.Println(uuid)
	if err := db.Set("gorm:auto_preload", true).Where("uuid", uuid).First(&object).Error; err != nil {
		c.JSON(404, "Not found")
	} else {
		c.JSON(200, object)
	}
}

func postAbstract(c *gin.Context, object interface{}) {
	if err := db.Create(object).Error; err != nil {
		c.JSON(500, "Server Error")
	} else{
		c.JSON(201, object)
	}
}


// BasicResponse : homeapge of the api
func BasicResponse(c *gin.Context) {
	c.JSON(200, gin.H{
		"response": "Welcome to go-vote API",
	})
}
