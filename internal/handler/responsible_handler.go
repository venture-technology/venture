package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/exceptions"
	"github.com/venture-technology/venture/internal/usecase/responsible"
	"github.com/venture-technology/venture/pkg/utils"
)

type ResponsibleHandler struct {
	responsibleUseCase *responsible.ResponsibleUseCase
}

func NewResponsibleHandler(responsibleUseCase *responsible.ResponsibleUseCase) *ResponsibleHandler {
	return &ResponsibleHandler{
		responsibleUseCase: responsibleUseCase,
	}
}

func (rh *ResponsibleHandler) Create(c *gin.Context) {

	var input entity.Responsible

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	input.Password = utils.HashPassword(input.Password)

	// fazendo get para verificar se o usuário existe
	_, err := rh.responsibleUseCase.Get(c, &input.CPF)

	if err == nil {
		log.Print("error to create responsible, responsible already exists")
		c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "user already exists"))
		return
	}

	cust, err := rh.responsibleUseCase.CreateCustomer(c, &input)
	if err != nil {
		log.Printf("error to create customer at stripe: %s", err.Error())
		c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "an error occured when creating customer at stripe"))
		return
	}

	input.CustomerId = cust.ID

	if !input.IsCreditCardEmpty() {
		paymentMethod, err := rh.responsibleUseCase.CreatePaymentMethod(c, &input.CreditCard.CardToken)

		if err != nil {
			log.Printf("error to create payment method at stripe: %s", err.Error())
			c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "an error occured when create payment method"))
			return
		}

		input.PaymentMethodId = paymentMethod.ID

		_, err = rh.responsibleUseCase.AttachPaymentMethod(c, &input.CustomerId, &input.PaymentMethodId, input.CreditCard.Default)

		if err != nil {
			log.Printf("error to create payment method at stripe: %s", err.Error())
			c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "an error occured when create payment method"))
			return
		}
	}

	err = rh.responsibleUseCase.Create(c, &input)
	if err != nil {
		log.Printf("error to create responsible: %s", err.Error())
		c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "an error occured when creating responsible"))
		return
	}

	log.Print("responsible create was successful")

	c.JSON(http.StatusCreated, input)

}

func (rh *ResponsibleHandler) Get(c *gin.Context) {

	cpf := c.Param("cpf")

	responsible, err := rh.responsibleUseCase.Get(c, &cpf)
	if err != nil {
		log.Printf("error while found responsible: %s", err.Error())
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "responsible not found"))
		return
	}

	c.JSON(http.StatusOK, responsible)

}

func (rh *ResponsibleHandler) Update(c *gin.Context) {

	cpf := c.Param("cpf")

	var input entity.Responsible

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	input.CPF = cpf

	currentResponsible, err := rh.responsibleUseCase.Get(c, &input.CPF)
	if err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "internal server error at get current user"))
		return
	}

	_, err = rh.responsibleUseCase.UpdateCustomer(c, currentResponsible)
	if err != nil {
		log.Printf("customer update stripe error: %s", err.Error())
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "customer update stripe error"))
		return
	}

	err = rh.responsibleUseCase.Update(c, currentResponsible, &input)
	if err != nil {
		log.Printf("responsible update error: %s", err.Error())
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "responsible update error"))
		return
	}

	log.Print("infos updated")

	c.JSON(http.StatusOK, gin.H{"message": "updated w successfully"})

}

func (rh *ResponsibleHandler) Delete(c *gin.Context) {

	cpf := c.Param("cpf")

	// buscando customerid do responsible
	responsible, err := rh.responsibleUseCase.Get(c, &cpf)
	if err != nil {
		log.Printf("get customerid error: %s", err.Error())
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "get customerid error"))
		return
	}

	_, err = rh.responsibleUseCase.DeleteCustomer(c, responsible.CustomerId)
	if err != nil {
		log.Printf("delete customerid error: %s", err.Error())
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "delete customerid error"))
		return
	}

	err = rh.responsibleUseCase.Delete(c, &cpf)
	if err != nil {
		log.Printf("error whiling deleted school: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "error to deleted school"})
		return
	}

	c.SetCookie("token", "", -1, "/", c.Request.Host, false, true)

	log.Printf("deleted your account --> %v", cpf)

	c.JSON(http.StatusOK, gin.H{"message": "responsible deleted w successfully"})

}

func (rh *ResponsibleHandler) SaveCard(c *gin.Context) {

	var input entity.Responsible

	if err := c.BindJSON(&input); err != nil {
		log.Printf("error to parsed body: %s", err.Error())
		c.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	paymentMethod, err := rh.responsibleUseCase.CreatePaymentMethod(c, &input.CreditCard.CardToken)
	if err != nil {
		log.Printf("error to create payment method: %s", err.Error())
		c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "an error occured when register payment method in stripe"))
		return
	}

	input.PaymentMethodId = paymentMethod.ID

	// get user to get customer id
	responsible, err := rh.responsibleUseCase.Get(c, &input.CPF)
	if err != nil {
		log.Printf("error to get customer id: %s", err.Error())
		c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "an error occured when getting customer id"))
		return
	}

	input.CustomerId = responsible.CustomerId

	_, err = rh.responsibleUseCase.AttachPaymentMethod(c, &input.CustomerId, &input.PaymentMethodId, input.CreditCard.Default)

	if err != nil {
		log.Printf("error to create payment method at stripe: %s", err.Error())
		c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "an error occured when create payment method"))
		return
	}

	err = rh.responsibleUseCase.SaveCard(c, &input.CPF, &input.CreditCard.CardToken, &paymentMethod.ID)
	if err != nil {
		log.Printf("error to register card: %s", err.Error())
		c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "an error occured when register credit card"))
		return
	}

	c.JSON(http.StatusOK, responsible.CreditCard)

}
