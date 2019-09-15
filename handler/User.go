package handler

import (
	"github.com/go-vote/model"
	"github.com/go-vote/provider"
	"github.com/biezhi/gorm-paginator/pagination"
	"strconv"
	"github.com/gin-gonic/gin"
)


var db = provider.GetDB()

func GetUsers(c *gin.Context) {

		var users []model.User

		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "3"))


		paginator := pagination.Paging(&pagination.Param{
		    DB:      db,
		    Page:    page,
		    Limit:   limit,
		    OrderBy: []string{"id desc"},
		    ShowSQL: false,
		}, &users)

		c.JSON(200, paginator)
}

func GetUser(c *gin.Context) {

		var user model.User

		id := c.Param("id")

		if err := db.First(&user, id).Error; err != nil {
			c.JSON(404, "Not found")
		} else{
			c.JSON(200, user)
		}
}

func PostUser(c *gin.Context) {

		content := c.PostForm("content")

		if err := db.Create(&model.User{Content: content}).Error; err != nil {
			c.JSON(404, "Not found")
		} else{

		c.JSON(201, gin.H{
			"status": 	"created",
			"type": 		"user",
			"content":  content,
		})
		}
}
