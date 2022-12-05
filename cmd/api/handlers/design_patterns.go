package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DesignPatternsHandler struct {
	service DesignPatternGetter
}

func NewDesignPatternsHandler(service DesignPatternGetter) DesignPatternsHandler {
	return DesignPatternsHandler{
		service: service,
	}
}

func (s DesignPatternsHandler) GetPatternByID(c *gin.Context) {
	// aca la key la deberiamos obtener como un parametro de la url
	response := s.service.GetPattern("some-key")

	c.JSON(http.StatusOK, gin.H{
		"message": response.Title,
	})
}
