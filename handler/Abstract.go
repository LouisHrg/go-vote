package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-vote/provider"
	"strconv"
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

	if err := db.Set("gorm:auto_preload", true).Where("uuid = ?", uuid).First(object).Error; err != nil {
		c.JSON(404, gin.H{
			"code": 404,
			"message": "Ressource not found",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"data": object,
		})
	}
}

func postAbstract(c *gin.Context, object interface{}) {

	c.ShouldBindJSON(&object)

	if err := db.Create(object).Error; err != nil {
		c.SecureJSON(400, gin.H{
			"code": 400,
			"message": "Bad request",
			"errors": err,
		})
	} else{
		c.JSON(201, gin.H{
			"code": 201,
			"message": "Ressource created",
			"data": object,
		})
	}
}


// BasicResponse : homeapge of the api
func BasicResponse(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"message": "Success",
	})
}
