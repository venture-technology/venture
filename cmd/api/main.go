package main

import (
	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/cmd/api/gateway"

	_ "github.com/lib/pq"
)

func main() {

	r := gin.Default()

	v1 := r.Group("api/v1")
	v1.Use(gateway.SetHeaders())
	gateway.NewGateway(r, v1).Setup()

}
