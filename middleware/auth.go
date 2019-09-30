package middleware

import (
	"os"
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

	// Workaround to get around the gorm hook and get the password
	if !db.Table("users").Select("email, password, uuid, accesslevel").Where("email = ?", loginParams.Username).Scan(&user).RecordNotFound() {

	if model.CheckPasswordHash(loginParams.Password, user.Password) {

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"uuid": user.UUID,
			"acl": user.Accesslevel,
			"time":  time.Now().Unix(),
		})

		tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET")))

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

	var ip model.IP
	// Update attempts for client IP
	db.Where(model.IP{Address: c.ClientIP()}).FirstOrInit(&ip, model.IP{Address: c.ClientIP()})
	ip.Attempt++
	db.Save(&ip)

	c.JSON(400, UnsignedResponse{
		Message: "Invalid username or/and password",
	})
}

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("Invalid header value given")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("Incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}

func parseToken(jwtToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, OK := token.Method.(*jwt.SigningMethodHMAC); !OK {
			return nil, errors.New("bad signed method received")
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		return nil, errors.New("Invalid JWT token")
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
			Message: "Invalid jwt token",
		})
		return
	}

	Claims, OK := token.Claims.(jwt.MapClaims)
	if !OK {
		c.AbortWithStatusJSON(http.StatusInternalServerError, UnsignedResponse{
			Message: "Unable to parse claims",
		})
		return
	}

	var tokenTime = Claims["time"].(float64)

	// The token expire after 30 minutes
	if (tokenTime + 1800) < float64(time.Now().Unix()) {
		c.AbortWithStatusJSON(403, UnsignedResponse{
			Message: "Token has expired",
		})
		return
	}

	c.Set("uuid", Claims["uuid"].(string))

	c.Next()
}

// ACLCheck : Check if the user is admin or not
func ACLCheck(c *gin.Context) {
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
			Message: "Invalid jwt token",
		})
		return
	}

	Claims, OK := token.Claims.(jwt.MapClaims)
	if !OK {
		c.AbortWithStatusJSON(http.StatusInternalServerError, UnsignedResponse{
			Message: "Unable to parse claims",
		})
		return
	}

	var acl = Claims["acl"].(string)

	if acl != "1" {
		c.AbortWithStatusJSON(403, UnsignedResponse{
			Message: "You are not authorized to perform this action",
		})
		return
	}

	c.Next()
}
