package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"ton-event-bot/domain"

	"github.com/gin-gonic/gin"
)

type SubController struct {
	SubUsecase domain.SubUsecase
}

func (tc *SubController) Create(c *gin.Context) {
	var input domain.Sub
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	if err := tc.SubUsecase.Create(c, &input); err != nil {
		c.JSON(http.StatusBadGateway, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "ok",
	})
}
func (tc *SubController) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	subId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	var userIdstr domain.ChangePersonRequest
	if err := c.BindJSON(&userIdstr); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	fmt.Println(subId)
	userId, err := strconv.Atoi(userIdstr.ChangePerson)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if err := tc.SubUsecase.Delete(c, subId, userId); err != nil {
		c.JSON(http.StatusBadGateway, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "ok",
	})

}
func (tc *SubController) Update(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "ok",
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
