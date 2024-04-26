package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/sysutil"
	. "github.com/kmou424/yukabot/global/logger"
)

var accessKey string

func init() {
	if key := sysutil.Getenv("ACCESS_KEY", ""); key != "" {
		accessKey = key
	} else {
		Logger.Warn("ACCESS_KEY is not set, auth will be disabled, it's not safe")
	}
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if accessKey != "" {
			if c.Request.Header.Get("X-YukaBot-Access-Key") != accessKey {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access denied"})
				return
			}
		}

		c.Next()
	}
}
