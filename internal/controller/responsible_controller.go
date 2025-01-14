package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/entity"
	"github.com/venture-technology/venture/internal/exceptions"
	"github.com/venture-technology/venture/internal/infra"
	"github.com/venture-technology/venture/internal/usecase"
	"github.com/venture-technology/venture/pkg/utils"
)

type ResponsibleController struct {
}

func NewResponsibleController() *ResponsibleController {
	return &ResponsibleController{}
}

func (rh *ResponsibleController) Create(c *gin.Context) {
	var input entity.Responsible
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	usecase := usecase.NewCreateResponsibleUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
		infra.App.Config,
	)

	input.Password = utils.MakeHash(input.Password)

	err := usecase.CreateResponsible(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "erro ao tentar criar responsável"))
		return
	}

	c.JSON(http.StatusCreated, input)
}

func (rh *ResponsibleController) Get(c *gin.Context) {
	cpf := c.Param("cpf")

	usecase := usecase.NewGetResponsibleUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	responsible, err := usecase.GetResponsible(cpf)
	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "responsavel não encontrado"))
		return
	}

	c.JSON(http.StatusOK, responsible)
}

func (rh *ResponsibleController) PatchV1UpdateResponsible(c *gin.Context) {
	cpf := c.Param("cpf")
	var data map[string]interface{}
	if err := c.BindJSON(data); err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	usecase := usecase.NewUpdateResponsibleUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	err := usecase.UpdateResponsible(cpf, data)
	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "erro ao tentar atualizar as informações do responsável na stripe"))
		return
	}

	c.JSON(http.StatusNoContent, http.NoBody)
}

func (rh *ResponsibleController) Delete(c *gin.Context) {
	cpf := c.Param("cpf")

	// buscando customerid do responsible
	responsible, err := rh.responsibleUseCase.Get(c, &cpf)
	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "ao tentar buscar a chave do cliente no stripe"))
		return
	}

	_, err = rh.responsibleUseCase.DeleteCustomer(c, responsible.CustomerId)
	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InternalServerResponseError(err, "erro ao deletar cliente na stripe"))
		return
	}

	err = rh.responsibleUseCase.Delete(c, &cpf)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "erro ao deletar responsável"})
		return
	}

	c.SetCookie("token", "", -1, "/", c.Request.Host, false, true)
	c.JSON(http.StatusNoContent, http.NoBody)
}

func (rh *ResponsibleController) SaveCard(c *gin.Context) {
	var input entity.Responsible

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	paymentMethod, err := rh.responsibleUseCase.CreatePaymentMethod(c, &input.CreditCard.CardToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "erro ao buscar método de pagamento na stripe"))
		return
	}

	input.PaymentMethodId = paymentMethod.ID

	// get user to get customer id
	responsible, err := rh.responsibleUseCase.Get(c, &input.CPF)
	if err != nil {
		c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "erro ao buscar chave do cliente na stripe"))
		return
	}

	input.CustomerId = responsible.CustomerId

	_, err = rh.responsibleUseCase.AttachPaymentMethod(c, &input.CustomerId, &input.PaymentMethodId, input.CreditCard.Default)

	if err != nil {
		c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "erro ao criar método de pagamento no stripe"))
		return
	}

	err = rh.responsibleUseCase.SaveCard(c, &input.CPF, &input.CreditCard.CardToken, &paymentMethod.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "erro ao registrar cartão na stripe"))
		return
	}

	c.JSON(http.StatusNoContent, http.NoBody)
}

func (rh *ResponsibleController) CalculatePrice(c *gin.Context) {
	var input entity.MapPrice
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	usecase := usecase.NewCalculatePriceUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
		infra.App.Adapters,
	)

	origin := fmt.Sprintf(
		"%s,%s,%s",
		input.Origin.Street,
		input.Origin.Number,
		input.Origin.ZIP,
	)

	destination := fmt.Sprintf(
		"%s,%s,%s",
		input.Destination.Street,
		input.Destination.Number,
		input.Destination.ZIP,
	)

	value, err := usecase.CalculatePrice(origin, destination, input.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"valor": value})
}
