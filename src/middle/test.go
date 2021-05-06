package middle

import (
	"github.com/gin-gonic/gin"
)

// A middleware uesd to test, if the query contains middle parameter, continue, else abort the request.
func PrintMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		middle := c.Query("middle")
		if middle == "" {
			c.JSON(400, gin.H{
				"status": "缺少middle参数",
			})
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}
