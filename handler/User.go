package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-vote/model"
)

// GetUsers : fetch all the users from the database
func GetUsers(c *gin.Context) {

	var users []model.User

	getAllAbstract(c, &users)

}

// GetUser : fetch one user by his id
func GetUser(c *gin.Context) {

	var user model.User

	getAbstract(c, &user)
}

// PostUser : create a new user
func PostUser(c *gin.Context) {

	var user model.User

	postAbstract(c, &user)
}


