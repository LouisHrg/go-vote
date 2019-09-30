package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-vote/model"
	"github.com/go-vote/provider"
)

// GetSurveys : fetch all the surveys from the database
func GetSurveys(c *gin.Context) {

	var surveys []model.Survey

	getAll(c, &surveys)

}

// GetSurvey : fetch one survey by his id
func GetSurvey(c *gin.Context) {

	var survey model.Survey

	get(c, &survey)
}

// PostSurvey : create a new user
func PostSurvey(c *gin.Context) {

	var survey model.Survey

	post(c, &survey)
}


// PatchResponse : update a new user
func PatchResponse(c *gin.Context) {

	var response model.Response
	var responsePatch model.Response

	patch(c, &response, &responsePatch)
}

// Vote : Add user to the response ressource
func Vote(c *gin.Context) {

	var response model.Response
	var user model.User

	userUUID := c.MustGet("uuid").(string)

	uuid := c.Param("uuid")

	provider.GetDB().Model(response).Where("uuid = ?", uuid).First(&response)
	provider.GetDB().Model(user).Where("uuid = ?", userUUID).First(&user)

	if err := provider.GetDB().Model(response).Association("Users").Append(user).Error; err != nil {
		c.JSON(404, gin.H{
			"code": 404,
			"message": "Ressource not found",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"data": response,
		})
	}
}

// DeleteVote : Remove user from the response ressource
func DeleteVote(c *gin.Context) {

  var response model.Response
  var user model.User

  userUUID := c.MustGet("uuid").(string)

  uuid := c.Param("uuid")

  provider.GetDB().Model(response).Where("uuid = ?", uuid).First(&response)
  provider.GetDB().Model(user).Where("uuid = ?", userUUID).First(&user)

  if err := provider.GetDB().Model(response).Association("Users").Delete(user).Error; err != nil {
  	c.JSON(404, gin.H{
  		"code": 404,
  		"message": "Ressource not found",
  	})
  } else {
  	c.JSON(200, gin.H{
  		"code": 200,
  		"data": response,
  	})
  }

}


// DeleteSurvey : delete an user
func DeleteSurvey(c *gin.Context) {

	var survey model.Survey

	delete(c, &survey)
}
