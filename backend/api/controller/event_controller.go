package controller

import(
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"ton-event-bot/domain"
)


type EventController struct {
	EventUsecase domain.EventUsecase	
}

func (tc *EventController) Create(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "Ok",
	})
}
func (tc *EventController) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "Ok",
	})

}
func (tc *EventController) Update(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "Ok",
	})
}
func (tc *EventController) GetById(c *gin.Context) {
	var data domain.Event
	id := c.Params.ByName("id")
	eventId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	data, err = tc.EventUsecase.GetById(c, eventId)
        if err != nil {
                c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
                return
        }
        c.JSON(http.StatusOK, data)
}
func (tc *EventController) GetAll(c *gin.Context) {
	var data []domain.Event
	data, err := tc.EventUsecase.GetAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

