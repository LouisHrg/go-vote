package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-vote/model"
)

// GetSurveys : fetch all the surveys from the database
func GetSurveys(c *gin.Context) {

	var surveys []model.Survey

	getAllAbstract(c, &surveys)

}

// GetSurvey : fetch one survey by his id
func GetSurvey(c *gin.Context) {

	var survey model.Survey

	getAbstract(c, &survey)
}
