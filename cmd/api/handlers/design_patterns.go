package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/waydevs/sections-api/internal/designpatters"
)

type DesignPatternsHandler struct {
	service DesignPatternService
}

func NewDesignPatternsHandler(service DesignPatternService) DesignPatternsHandler {
	return DesignPatternsHandler{
		service: service,
	}
}

func (s DesignPatternsHandler) GetPatternByID(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param(desingPatternIDParam)

	response, err := s.service.GetByID(ctx, id)

	if err != nil {
		httpCode := http.StatusInternalServerError

		if errors.Is(err, designpatters.ErrDesignPatternNotFound) {
			httpCode = http.StatusNotFound
		}

		c.JSON(httpCode, Response{
			Status:  httpCode,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: "",
		Data:    response,
	})
}

func (s DesignPatternsHandler) CreatePattern(c *gin.Context) {
	ctx := c.Request.Context()

	var request designpatters.DesignPattern
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	response, err := s.service.Create(ctx, request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusCreated, Response{
		Status:  http.StatusCreated,
		Message: "",
		Data:    response,
	})
}

func (s DesignPatternsHandler) DeletePattern(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param(desingPatternIDParam)

	err := s.service.Delete(ctx, id)

	if err != nil {
		httpCode := http.StatusInternalServerError

		if errors.Is(err, designpatters.ErrDesignPatternNotFound) {
			httpCode = http.StatusNotFound
		}

		c.JSON(httpCode, Response{
			Status:  httpCode,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: "Design pattern deleted successfully",
		Data:    nil,
	})
}

func (s DesignPatternsHandler) UpdatePattern(c *gin.Context) {
	ctx := c.Request.Context()

	var request designpatters.DesignPattern
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	response, err := s.service.Update(ctx, request)

	if err != nil {
		httpCode := http.StatusInternalServerError

		if errors.Is(err, designpatters.ErrDesignPatternNotFound) {
			httpCode = http.StatusNotFound
		}

		c.JSON(httpCode, Response{
			Status:  httpCode,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: "",
		Data:    response,
	})
}
