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
		c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "esse responsável já existe"))
		return
	}

	cust, err := rh.responsibleUseCase.CreateCustomer(c, &input)
	if err != nil {
		log.Printf("error to create customer at stripe: %s", err.Error())
		c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "aconteceu algum erro ao tentar criar o cliente na stripe"))
		return
	}

	input.CustomerId = cust.ID

	if !input.IsCreditCardEmpty() {
		paymentMethod, err := rh.responsibleUseCase.CreatePaymentMethod(c, &input.CreditCard.CardToken)

		if err != nil {
			log.Printf("error to create payment method at stripe: %s", err.Error())
			c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "aconteceu algum erro ao tentar registrar método de pagamento na stripe"))
			return
		}

		input.PaymentMethodId = paymentMethod.ID

		_, err = rh.responsibleUseCase.AttachPaymentMethod(c, &input.CustomerId, &input.PaymentMethodId, input.CreditCard.Default)

		if err != nil {
			log.Printf("error to create payment method at stripe: %s", err.Error())
			c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "aconteceu algum erro ao tentar registrar método de pagamento na stripe"))
			return
		}
	}

	err = rh.responsibleUseCase.Create(c, &input)
	if err != nil {
		log.Printf("error to create responsible: %s", err.Error())
		c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "erro inesperado ao tentar criar responsável"))
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
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "responsavel não encontrado"))
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
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "erro interno de servidor ao tentar buscar o responsável atual"))
		return
	}

	_, err = rh.responsibleUseCase.UpdateCustomer(c, currentResponsible)
	if err != nil {
		log.Printf("customer update stripe error: %s", err.Error())
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "erro ao tentar atualizar as informações do responsável na stripe"))
		return
	}

	err = rh.responsibleUseCase.Update(c, currentResponsible, &input)
	if err != nil {
		log.Printf("responsible update error: %s", err.Error())
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "erro ao tentar atualizar as informações do responsável"))
		return
	}

	log.Print("infos updated")

	c.JSON(http.StatusNoContent, http.NoBody)

}

func (rh *ResponsibleHandler) Delete(c *gin.Context) {

	cpf := c.Param("cpf")

	// buscando customerid do responsible
	responsible, err := rh.responsibleUseCase.Get(c, &cpf)
	if err != nil {
		log.Printf("get customerid error: %s", err.Error())
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "ao tentar buscar a chave do cliente no stripe"))
		return
	}

	_, err = rh.responsibleUseCase.DeleteCustomer(c, responsible.CustomerId)
	if err != nil {
		log.Printf("delete customerid error: %s", err.Error())
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "erro ao deletar cliente na stripe"))
		return
	}

	err = rh.responsibleUseCase.Delete(c, &cpf)
	if err != nil {
		log.Printf("error whiling deleted responsible: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "erro ao deletar responsável"})
		return
	}

	c.SetCookie("token", "", -1, "/", c.Request.Host, false, true)

	log.Printf("deleted your account --> %v", cpf)

	c.JSON(http.StatusNoContent, http.NoBody)

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
		c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "erro ao buscar método de pagamento na stripe"))
		return
	}

	input.PaymentMethodId = paymentMethod.ID

	// get user to get customer id
	responsible, err := rh.responsibleUseCase.Get(c, &input.CPF)
	if err != nil {
		log.Printf("error to get customer id: %s", err.Error())
		c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "erro ao buscar chave do cliente na stripe"))
		return
	}

	input.CustomerId = responsible.CustomerId

	_, err = rh.responsibleUseCase.AttachPaymentMethod(c, &input.CustomerId, &input.PaymentMethodId, input.CreditCard.Default)

	if err != nil {
		log.Printf("error to create payment method at stripe: %s", err.Error())
		c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "erro ao criar método de pagamento no stripe"))
		return
	}

	err = rh.responsibleUseCase.SaveCard(c, &input.CPF, &input.CreditCard.CardToken, &paymentMethod.ID)
	if err != nil {
		log.Printf("error to register card: %s", err.Error())
		c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "erro ao registrar cartão na stripe"))
		return
	}

	c.JSON(http.StatusNoContent, http.NoBody)

}
