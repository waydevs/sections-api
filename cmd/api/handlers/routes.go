package handlers

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/waydevs/sections-api/internal/designpatters"
)

const (
	designPattersGroup   = "designpatters"
	desingPatternIDParam = "id"
)

type DesignPatternService interface {
	GetByID(ctx context.Context, id string) (designpatters.DesignPattern, error)
	Create(ctx context.Context, designPattern designpatters.DesignPattern) (designpatters.DesignPattern, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, designPattern designpatters.DesignPattern) (designpatters.DesignPattern, error)
}

func DesignPatternRoutes(router *gin.Engine, service DesignPatternService) *gin.Engine {
	group := router.Group(designPattersGroup)

	handler := NewDesignPatternsHandler(service)
	group.GET(fmt.Sprintf("/:%s", desingPatternIDParam), handler.GetPatternByID)
	group.POST("", handler.CreatePattern)
	group.DELETE(fmt.Sprintf("/:%s", desingPatternIDParam), handler.DeletePattern)
	group.PUT("", handler.UpdatePattern)

	return router
}
