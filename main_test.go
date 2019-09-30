package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHomeRoute(t *testing.T) {
	router := initRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	body := gin.H{
		"code":200,
		"message":"Success",
	}

   	assert.Equal(t, body["message"], "Success")
	assert.Equal(t, body["code"], w.Code)
}

func TestUsersRouteUnLogged(t *testing.T) {
	router := initRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)
	router.ServeHTTP(w, req)

	body := gin.H{
		"message": "Invalid header value given",
	}
   	assert.Equal(t, body["message"], "Invalid header value given")
	assert.Equal(t, 400, w.Code)
}


func TestSurveysRouteUnLogged(t *testing.T) {
	router := initRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/surveys", nil)
	router.ServeHTTP(w, req)

	body := gin.H{
		"message": "Invalid header value given",
	}
   	assert.Equal(t, body["message"], "Invalid header value given")
	assert.Equal(t, 400, w.Code)
}