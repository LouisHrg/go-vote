package middleware

import (
	"net/http"
	"errors"
	"time"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-vote/model"
	"github.com/go-vote/provider"
)

var db = provider.GetDB()

type UnsignedResponse struct {
	Message interface{} `json:"message"`
}

type SignedResponse struct {
	Jwt   string `json:"token"`
}

func index(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "index"})
}

// LoginHandler : The login handler, set it for a /login route for example
func LoginHandler(c *gin.Context) {

	type login struct {
		Username string `json:"email,omitempty"`
		Password string `json:"password,omitempty"`
	}

	loginParams := login{}

	c.ShouldBindJSON(&loginParams)

	type Result struct {
	  Email 			string
	  UUID 				string
	  Accesslevel string
	  Password  	string
	}

	var user Result

	// Workaround to get around the hook and get the password
	if !db.Table("users").Select("email, password, uuid, accesslevel").Where("email = ?", loginParams.Username).Scan(&user).RecordNotFound() {

	if model.CheckPasswordHash(loginParams.Password, user.Password) {

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"uuid": user.UUID,
			"acl": user.Accesslevel,
			"time":  time.Now().Unix(),
		})

		tokenStr, err := token.SignedString([]byte("supersaucysecret"))

		if err != nil {
			c.JSON(500, UnsignedResponse{
				Message: err.Error(),
			})
			return
		}

		c.JSON(200, SignedResponse{
			Jwt:   tokenStr,
		})
		return
		}

	}

	c.JSON(400, UnsignedResponse{
		Message: "bad username",
	})
}

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("bad header value given")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}

func parseToken(jwtToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, OK := token.Method.(*jwt.SigningMethodHMAC); !OK {
			return nil, errors.New("bad signed method received")
		}
		return []byte("supersaucysecret"), nil
	})

	if err != nil {
		return nil, errors.New("bad jwt token")
	}

	return token, nil
}
// JwtTokenCheck : the middleware, register it for a group of routes
func JwtTokenCheck(c *gin.Context) {
	jwtToken, err := extractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: err.Error(),
		})
		return
	}

	token, err := parseToken(jwtToken)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: "bad jwt token",
		})
		return
	}

	Claims, OK := token.Claims.(jwt.MapClaims)
	if !OK {
		c.AbortWithStatusJSON(http.StatusInternalServerError, UnsignedResponse{
			Message: "unable to parse claims",
		})
		return
	}

	var tokenTime = Claims["time"].(float64)

	// The token expire after 30 minutes
	if (tokenTime + 1800) < float64(time.Now().Unix()) {
		c.AbortWithStatusJSON(403, UnsignedResponse{
			Message: "token has expired",
		})
		return
	}

	c.Next()
}

func privateACLCheck(c *gin.Context) {
	jwtToken, err := extractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: err.Error(),
		})
		return
	}

	token, err := parseToken(jwtToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: "bad jwt token",
		})
		return
	}

	claims, OK := token.Claims.(jwt.MapClaims)
	if !OK {
		c.AbortWithStatusJSON(http.StatusInternalServerError, UnsignedResponse{
			Message: "unable to parse claims",
		})
		return
	}

	claimedUID, OK := claims["user"].(string)
	if !OK {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: "no user property in claims",
		})
		return
	}

	uid := c.Param("uid")
	if claimedUID != uid {
		c.AbortWithStatusJSON(http.StatusBadRequest, UnsignedResponse{
			Message: "token uid does not match resource uid",
		})
		return
	}

	c.Next()
}
