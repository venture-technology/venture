package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venture-technology/venture/internal/domain/service/agreements"
	"github.com/venture-technology/venture/internal/exceptions"
	"github.com/venture-technology/venture/internal/infra"
	"github.com/venture-technology/venture/internal/usecase"
)

type WebhookController struct {
}

func NewWebhookController() *WebhookController {
	return &WebhookController{}
}

func (wh *WebhookController) PostV1WebhookEvents(httpContext *gin.Context) {
	bodyReceived := httpContext.PostForm("json")
	if bodyReceived == "" {
		httpContext.JSON(http.StatusBadRequest, exceptions.InvalidBodyContentResponseError(fmt.Errorf("empty body")))
		return
	}

	var event agreements.EventWrapper
	if err := json.Unmarshal([]byte(bodyReceived), &event); err != nil {
		infra.App.Logger.Infof(fmt.Sprintf("unmarshall error: %s", err.Error()))
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao desserializar JSON"})
		return
	}

	usecase := usecase.NewWebhookEventsUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
	)

	_, err := usecase.Execute(event)
	if err != nil {
		infra.App.Logger.Infof(fmt.Sprintf("usecase error: %s", err.Error()))
		httpContext.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "error handling webhook event"))
		return
	}

	requestBody, err := httpContext.GetRawData()
	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, exceptions.InvalidBodyContentResponseError(err))
		return
	}

	infra.App.Logger.Infof(fmt.Sprintf("event: %v", event))
	infra.App.Logger.Infof(fmt.Sprintf("requestParams: %s", string(requestBody)))

	httpContext.Header("Content-Type", "text/plain")
	httpContext.String(http.StatusOK, "Hello API Event Received")
}
