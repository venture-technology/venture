package exceptions

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func InvalidBodyContentResponseError(err error) gin.H {
	return gin.H{
		"error": "conteúdo do body inválido",
	}
}

func InternalServerResponseError(err error, msg string) gin.H {
	return gin.H{
		"error": fmt.Sprintf("erro interno de servidor - %s: %s", msg, err.Error()),
	}
}

func TypeServerResponseError(msg string) gin.H {
	return gin.H{
		"error": fmt.Sprintf("type error: %s", msg),
	}
}

func NotParamErrorResponse(param string) gin.H {
	return gin.H{
		"error": fmt.Sprintf("O paramêtro '%s' não foi encontrado", param),
	}
}

func NotFoundObjectErrorResponse(obj string) gin.H {
	return gin.H{
		"error": fmt.Sprintf("%s não foi encontrado", obj),
	}
}
