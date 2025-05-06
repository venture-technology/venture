package controller

import (
	"bytes"
	"fmt"
	"io"
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

func (wh *WebhookController) PostV1WebhookSignatureEvents(httpContext *gin.Context) {
	bodyBytes, err := io.ReadAll(httpContext.Request.Body)
	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao ler o corpo da requisição"})
		return
	}
	httpContext.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	var eventWrapper agreements.EventWrapper

	if err := httpContext.BindJSON(&eventWrapper); err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid JSON"})
		return
	}
	httpContext.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	infra.App.Logger.Infof(fmt.Sprintf("Received webhook event: %v", eventWrapper))

	usecase := usecase.NewWebhookSignatureEventsUseCase(
		&infra.App.Repositories,
		infra.App.Logger,
		infra.App.Adapters,
	)

	_, err = usecase.Execute(httpContext, eventWrapper)
	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, exceptions.InternalServerResponseError(err, "error handling webhook event"))
		return
	}

	httpContext.Header("Content-Type", "text/plain")
	httpContext.String(http.StatusOK, "Hello API Event Received")
}

func (wh *WebhookController) PostV1WebhookPaymentsEvents(httpContext *gin.Context) {
	httpContext.JSON(http.StatusOK, gin.H{"message": "Hello API Event Received"})
}
