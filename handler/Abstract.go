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

	id := c.Param("id")

	if err := db.Set("gorm:auto_preload", true).First(object, id).Error; err != nil {
		c.JSON(404, "Not found")
	} else {
		c.JSON(200, object)
	}
}

func postAbstract(c *gin.Context, object interface{}) {

}
