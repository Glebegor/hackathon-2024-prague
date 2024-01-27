package controller

import(
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
        

	data, err := tc.EventUsecase.GetById(c)
        if err != nil {
                c.JSON(http.StatusInternalServerError, map[string]interface{}{
                        "message": "Error: Cant get event.",
                })
                fmt.Print(err.Error())
                return
        }
        c.JSON(http.StatusOK, data)
}
func (tc *EventController) GetAll(c *gin.Context) {
	var data []domain.Event
	data, err := tc.EventUsecase.GetAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error: Cant get all events.", 
		})
		fmt.Print(err.Error())
		return
	}
	c.JSON(http.StatusOK, data)
}

