package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/venture-technology/venture/cmd/api/admin"
	v1 "github.com/venture-technology/venture/cmd/api/v1"
	"github.com/venture-technology/venture/internal/infra"
)

func main() {

	r := gin.Default()

	v1Api := r.Group("api/v1")
	admApi := r.Group("admin")

	app := infra.NewApplication(r, v1Api, admApi)

	v1.NewV1(app).Setup()
	admin.NewAdmin(app).Setup()

	r.Run(fmt.Sprintf(":%d", app.Config.Server.Port))

}
