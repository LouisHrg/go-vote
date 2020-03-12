package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-vote/model"
)

// GetUsers : fetch all the users from the database
func GetUsers(c *gin.Context) {

	var users []model.User

	getAll(c, &users)

}

// GetUser : fetch one user by his id
func GetUser(c *gin.Context) {

	var user model.User

	get(c, &user)
}

// GetCurrentUser : get current user info
func GetCurrentUser(c *gin.Context) {
	var user model.User

 	uuid, _ := c.Get("uuid")

	if err := db.Set("gorm:auto_preload", true).Where("uuid = ?", uuid).First(&user).Error; err != nil {
		c.JSON(404, gin.H{
			"code": 404,
			"message": "Ressource not found",
		})
	} else {
		c.JSON(200, user)
	}
}

// PostUser : create a new user
func PostUser(c *gin.Context) {

	var user model.User
	user.Accesslevel = 2
	post(c, &user)
}

// PatchUser : update an user
func PatchUser(c *gin.Context) {

	var user model.User
	var patchStruct model.User

	patch(c, &user, &patchStruct)
}

// DeleteUser : delete an user
func DeleteUser(c *gin.Context) {

	var user model.User

	delete(c, &user)
}

// PromoteUser : promote user ACL
func PromoteUser(c *gin.Context) {

	var user model.User

	uuid := c.Param("uuid")

	if err := db.Set("gorm:auto_preload", true).Where("uuid = ?", uuid).First(&user).Error; err != nil {
		c.JSON(404, gin.H{
			"code": 404,
			"message": "Ressource not found",
		})
	} else {

		if user.Accesslevel == 1 {
			user.Accesslevel = 2
		} else {
			user.Accesslevel = 1
		}

		db.Save(&user)

		c.JSON(200, gin.H{
			"code": 200,
			"data": user,
		})
	}
}
