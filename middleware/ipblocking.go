package middleware

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "github.com/go-vote/provider"
    "github.com/go-vote/model"
)

// IPFirewall : block banned IPs
func IPFirewall() gin.HandlerFunc {
    var db = provider.GetDB()
    var blocked model.IP
    return func(c *gin.Context) {
        if !db.Where("address = ?", c.ClientIP()).Find(&blocked).RecordNotFound() {
            c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
                "status":  http.StatusForbidden,
                "message": "Permission denied",
            })
            return
        }
    }
}
