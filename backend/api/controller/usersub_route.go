package controller

import (
	"net/http"
	"strconv"
	"ton-event-bot/domain"

	"github.com/gin-gonic/gin"
)

type UsersubController struct {
	UsersubUsecase domain.UsersubUsecase
}

func (tc *UsersubController) Create(c *gin.Context) {
	var input domain.Usersub
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	if err := tc.UsersubUsecase.Create(c, &input); err != nil {
		c.JSON(http.StatusBadGateway, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "ok",
	})
}
func (tc *UsersubController) Delete(c *gin.Context) {
	id := c.Params.ByName("user_id")
	userId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	id = c.Params.ByName("id_channel")
	usersubId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	var userIdstr domain.ChangePersonRequest
	if err := c.BindJSON(&userIdstr); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if err := tc.UsersubUsecase.Delete(c, usersubId, userId); err != nil {
		c.JSON(http.StatusBadGateway, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "ok",
	})

}
func (tc *UsersubController) Update(c *gin.Context) {
	id := c.Params.ByName("user_id")
	userId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	id = c.Params.ByName("id_channel")
	usersubId, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	var input domain.UpdateUsersub
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	if err := tc.UsersubUsecase.Update(c, usersubId, userId, &input); err != nil {
		c.JSON(http.StatusBadGateway, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "ok",
	})
}
func (tc *UsersubController) GetById(c *gin.Context) {
	var data domain.Usersub
	id := c.Params.ByName("user_id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	id = c.Params.ByName("id_channel")
	usersubId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	data, err = tc.UsersubUsecase.GetById(c, usersubId, userId)
	if err != nil {
		c.JSON(http.StatusBadGateway, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}
func (tc *UsersubController) GetAll(c *gin.Context) {
	id := c.Params.ByName("user_id")
	subId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	data, err := tc.UsersubUsecase.GetAll(c, subId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}
