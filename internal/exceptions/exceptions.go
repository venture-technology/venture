package exceptions

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InvalidBodyContentResponseError(err error) gin.H {
	return gin.H{
		"error": "invalid body content",
	}
}

func InternalServerResponseError(err error, msg string) gin.H {
	return gin.H{
		"error": fmt.Sprintf("internal server error - %s: %s", msg, err.Error()),
	}
}

func TypeServerResponseError(msg string) gin.H {
	return gin.H{
		"error": fmt.Sprintf("type error: %s", msg),
	}
}
