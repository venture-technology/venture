package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/exceptions"
	"github.com/venture-technology/venture/internal/service"
	"github.com/venture-technology/venture/models"
)

type ResponsibleController struct {
	responsibleservice *service.ResponsibleService
}

func NewResponsibleController(responsibleservice *service.ResponsibleService) *ResponsibleController {
	return &ResponsibleController{
		responsibleservice: responsibleservice,
	}
}

func (ct *ResponsibleController) RegisterRoutes(router *gin.Engine) {
	api := router.Group("api/v1/responsible")

	api.POST("/", ct.CreateResponsible)
	api.GET("/:cpf", ct.GetResponsible)
	api.PATCH("/:cpf", ct.UpdateResponsible)
	api.DELETE("/:cpf", ct.DeleteResponsible)
	api.POST("/:cpf/card", ct.RegisterCreditCard)
}

func (ct *ResponsibleController) CreateResponsible(c *gin.Context) {
	var input models.Responsible

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	_, err := ct.responsibleservice.GetResponsible(c, &input.CPF)

	// verificando se o usuário existe
	if err == nil {
		log.Print("error to create responsible, responsible already exists")
		c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "user already exists"))
		return
	}

	cust, err := ct.responsibleservice.CreateCustomer(c, &input)
	if err != nil {
		log.Printf("error to create customer at stripe: %s", err.Error())
		c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "an error occured when creating customer at stripe"))
		return
	}

	input.CustomerId = cust.ID

	if !input.IsCreditCardEmpty() {
		paymentMethod, err := ct.responsibleservice.CreatePaymentMethod(c, &input.CreditCard.CardToken)

		if err != nil {
			log.Printf("error to create payment method at stripe: %s", err.Error())
			c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "an error occured when create payment method"))
			return
		}

		input.PaymentMethodId = paymentMethod.ID

		_, err = ct.responsibleservice.AttachPaymentMethod(c, &input.CustomerId, &input.PaymentMethodId, input.CreditCard.Default)

		if err != nil {
			log.Printf("error to create payment method at stripe: %s", err.Error())
			c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "an error occured when create payment method"))
			return
		}
	}

	err = ct.responsibleservice.CreateResponsible(c, &input)
	if err != nil {
		log.Printf("error to create responsible: %s", err.Error())
		c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "an error occured when creating responsible"))
		return
	}

	log.Print("responsible create was successful")

	c.JSON(http.StatusCreated, input)
}

func (ct *ResponsibleController) GetResponsible(c *gin.Context) {
	cpf := c.Param("cpf")

	responsible, err := ct.responsibleservice.GetResponsible(c, &cpf)
	if err != nil {
		log.Printf("error while found responsible: %s", err.Error())
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "responsible not found"))
		return
	}

	c.JSON(http.StatusOK, responsible)
}

func (ct *ResponsibleController) UpdateResponsible(c *gin.Context) {

	cpf := c.Param("cpf")

	var input models.Responsible

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	input.CPF = cpf

	currentResponsible, err := ct.responsibleservice.GetResponsible(c, &input.CPF)
	if err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "internal server error at get current user"))
		return
	}

	_, err = ct.responsibleservice.UpdateCustomer(c, currentResponsible)
	if err != nil {
		log.Printf("customer update stripe error: %s", err.Error())
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "customer update stripe error"))
		return
	}

	err = ct.responsibleservice.UpdateResponsible(c, currentResponsible, &input)
	if err != nil {
		log.Printf("responsible update error: %s", err.Error())
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "responsible update error"))
		return
	}

	log.Print("infos updated")

	c.JSON(http.StatusOK, gin.H{"message": "updated w successfully"})
}

func (ct *ResponsibleController) DeleteResponsible(c *gin.Context) {

	cpf := c.Param("cpf")

	responsible, err := ct.responsibleservice.GetResponsible(c, &cpf)
	if err != nil {
		log.Printf("get customerid error: %s", err.Error())
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "get customerid error"))
		return
	}

	_, err = ct.responsibleservice.DeleteCustomer(c, responsible.CustomerId)
	if err != nil {
		log.Printf("delete customerid error: %s", err.Error())
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "delete customerid error"))
		return
	}

	err = ct.responsibleservice.DeleteResponsible(c, &cpf)
	if err != nil {
		log.Printf("error whiling deleted school: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "error to deleted school"})
		return
	}

	c.SetCookie("token", "", -1, "/", c.Request.Host, false, true)

	log.Printf("deleted your account --> %v", cpf)

	c.JSON(http.StatusOK, gin.H{"message": "responsible deleted w successfully"})
}

func (ct *ResponsibleController) RegisterCreditCard(c *gin.Context) {

	cpf := c.Param("cpf")

	var input models.CreditCard

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	input.CPF = cpf

	log.Print(input)

	paymentMethod, err := ct.responsibleservice.CreatePaymentMethod(c, &input.CardToken)
	if err != nil {
		log.Printf("error to create payment method: %s", err.Error())
		c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "an error occured when register payment method in stripe"))
		return
	}

	responsible, err := ct.responsibleservice.GetResponsible(c, &cpf)
	if err != nil {
		log.Printf("error to get customer id: %s", err.Error())
		c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "an error occured when getting customer id"))
		return
	}

	_, err = ct.responsibleservice.AttachPaymentMethod(c, &responsible.CustomerId, &paymentMethod.ID, false) // <- esse false é um mock
	if err != nil {
		log.Printf("attach card in customer error: %s", err.Error())
		c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "an error occured when attaching card in customer"))
		return
	}

	if input.Default {

		_, err := ct.responsibleservice.UpdatePaymentMethodDefault(c, &responsible.CustomerId, &paymentMethod.ID)
		if err != nil {
			log.Printf("change default card for customer error: %s", err.Error())
			c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "an error occured when changing default card for customer"))
			return
		}

		err = ct.responsibleservice.SaveCreditCard(c, &input.CPF, &input.CardToken, &paymentMethod.ID)
		if err != nil {
			log.Printf("error to register card: %s", err.Error())
			c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "an error occured when register credit card"))
			return
		}

	}

	c.JSON(http.StatusOK, gin.H{"message": "card attached in customer"})
}
