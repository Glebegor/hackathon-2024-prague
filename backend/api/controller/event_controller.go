package controller

import(
	// "net/http"
	"github.com/gin-gonic/gin"
	"ton-event-bot/domain"
)


type EventController struct {
	EventUsecase domain.EventUsecase	
}

func (tc *EventController) Create(c *gin.Context) {}
func (tc *EventController) Delete(c *gin.Context) {}
func (tc *EventController) Update(c *gin.Context) {}
func (tc *EventController) GetById(c *gin.Context) {}
func (tc *EventController) GetAll(c *gin.Context) {}
