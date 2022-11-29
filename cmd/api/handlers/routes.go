package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/waydevs/sections-api/internal/designpatters"
)

const (
	designPattersGroup = "designpatters"
)

type DesignPatternGetter interface {
	GetPattern(key string) designpatters.DesignPatternDTO
}

func DesignPatternRoutes(router *gin.Engine, service DesignPatternGetter) *gin.Engine {
	group := router.Group(designPattersGroup)

	handler := NewDesignPatternsHandler(service)
	// deberiamos recibir el id como parametro pero para el ejemplo lo omitimos
	group.GET("", handler.GetPatternByID)

	return router
}
