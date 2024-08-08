package auth

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/exceptions"
	"github.com/venture-technology/venture/internal/service"
	"github.com/venture-technology/venture/models"
)

type AuthRequests struct {
	responsibleservice *service.ResponsibleService
	schoolservice      *service.SchoolService
}

func NewAuth(responsibleservice *service.ResponsibleService, schoolservice *service.SchoolService) *AuthRequests {
	return &AuthRequests{
		responsibleservice: responsibleservice,
		schoolservice:      schoolservice,
	}
}

func (auth *AuthRequests) AuthRoutes(router *gin.Engine) {
	authApi := router.Group("/auth")

	authApi.POST("/responsible", auth.AuthResponsible)
	authApi.POST("/school", auth.AuthSchool)
	authApi.POST("/driver", auth.AuthDriver)
}

func (auth *AuthRequests) AuthResponsible(c *gin.Context) {
	var input models.Responsible

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}
	log.Printf("doing login --> %s", input.Email)

	responsible, err := auth.responsibleservice.AuthResponsible(c, &input)
	if err != nil {
		log.Printf("wrong email or password: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong email or password"})
		return
	}

	jwt, err := auth.responsibleservice.CreateTokenJWTResponsible(c, responsible)

	log.Printf("token returned --> %v", jwt)

	if err != nil {
		log.Printf("error to create jwt token: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "error to create jwt token"})
		return
	}

	c.SetCookie("token", jwt, 3600, "/", c.Request.Host, false, true)

	c.JSON(http.StatusAccepted, gin.H{
		"responsible": responsible,
		"token":       jwt,
	})
}

func (auth *AuthRequests) AuthSchool(c *gin.Context) {
	var input models.School

	log.Printf("doing login --> %s", input.Email)

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"erro": "invalid body content"})
		return
	}

	school, err := auth.schoolservice.AuthSchool(c, &input)

	if err != nil {
		log.Printf("wrong email or password: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong email or password"})
		return
	}

	jwt, err := auth.schoolservice.CreateTokenJWTSchool(c, school)

	log.Printf("token returned --> %v", jwt)

	if err != nil {
		log.Printf("error to create jwt token: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "error to create jwt token"})
		return
	}

	c.SetCookie("token", jwt, 3600, "/", c.Request.Host, false, true)

	c.JSON(http.StatusAccepted, gin.H{
		"school": school,
		"token":  jwt,
	})
}

func (auth *AuthRequests) AuthDriver(c *gin.Context) {
}
