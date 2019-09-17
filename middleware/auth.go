package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/appleboy/gin-jwt/v2"
	"github.com/go-vote/model"
	"time"

)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var identityKey = "uuid"


// AuthMiddleware : The authmiddlare handling jwt token security
func AuthMiddleware()  (*jwt.GinJWTMiddleware, error){
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: payloadFunc,
		IdentityHandler: identityHandler,
		Authenticator: authenticator,
		Authorizator: authorizator,
		Unauthorized: unauthorized,
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc: time.Now,
	})
}


func unauthorized (c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
}


func authorizator (data interface{}, c *gin.Context) bool {
			if v, ok := data.(*model.User); ok && v.Email == "admin" {
				return true
			}

			return false
}


func authenticator (c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password

			if (userID == "admin@gmail.com" && password == "admin") || (userID == "test@gmail.com" && password == "test") {
				return &model.User{
					Email:  userID,
					Lastname:  "Bo-Yi",
					Firstname: "Wu",
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		}

func identityHandler (c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &model.User{
		Email: claims[identityKey].(string),
	}
}


func payloadFunc (data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					identityKey: v.Email,
				}
			}
			return jwt.MapClaims{}
		}
