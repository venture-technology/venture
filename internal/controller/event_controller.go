package controller

import "github.com/gin-gonic/gin"

type EventController struct {
}

func NewEventController() *EventController {
	return &EventController{}
}

func (eh *EventController) PostV1CreateEvent(httpContext *gin.Context) {
}

func (eh *EventController) GetV1ListEvents(httpContext *gin.Context) {
}

func (eh *EventController) PatchV1UpdateEvent(httpContext *gin.Context) {
}

func (eh *EventController) DeleteV1DeleteEvent(httpContext *gin.Context) {
}

func (eh *EventController) GetV1ListEventsByIntegrator(httpContext *gin.Context) {
}
