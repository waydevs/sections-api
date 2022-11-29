package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/waydevs/sections-api/cmd/api/handlers"
	"github.com/waydevs/sections-api/internal/designpatters"
	"github.com/waydevs/sections-api/internal/platform/repository"
)

func main() {
	r := gin.Default()

	dbConn, err := repository.NewRepository()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dbDesignPattern := repository.NewDesignPattern(dbConn)

	designPatternService := designpatters.NewDesignPattersService(dbDesignPattern)

	r = handlers.DesignPatternRoutes(r, designPatternService)

	r.Run()
}
