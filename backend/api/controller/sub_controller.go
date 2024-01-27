package controller

import (
	"net/http"
	"strconv"
	"ton-event-bot/domain"

	"github.com/gin-gonic/gin"
)

type SubController struct {
	SubUsecase domain.SubUsecase
}

func (tc *SubController) Create(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "Ok",
	})
}
func (tc *SubController) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "Ok",
	})

}
func (tc *SubController) Update(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "Ok",
	})
}
func (tc *SubController) GetById(c *gin.Context) {
	var data domain.Sub
	id := c.Params.ByName("id")
	subId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	data, err = tc.SubUsecase.GetById(c, subId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}
func (tc *SubController) GetAll(c *gin.Context) {
	var data []domain.Sub
	data, err := tc.SubUsecase.GetAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}
