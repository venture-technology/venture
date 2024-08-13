package gateway

import "github.com/gin-gonic/gin"

func NewGateway(group *gin.RouterGroup) {

}

func ConfigureHeaders(c *gin.Context) {
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.Header("charset", "utf-8")
}
