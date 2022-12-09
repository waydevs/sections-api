package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/waydevs/sections-api/cmd/api/handlers"
	"github.com/waydevs/sections-api/internal/designpatters"
	"github.com/waydevs/sections-api/internal/platform/repository"
)

const (
	// Momentaneamente dejemoslo asi, pero en un futuro lo cambiaremos por una variable de entorno.
	mongoURI = "mongodb://localhost:27017"
)

func main() {
	r := gin.Default()

	dbConn, err := repository.NewClient(mongoURI)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer dbConn.Close()

	db := repository.NewDatabase(dbConn)

	desigPatternsRepositroy := repository.NewDesignPatterns(db)
	designPatternsService := designpatters.NewService(desigPatternsRepositroy)

	r = handlers.DesignPatternRoutes(r, designPatternsService)

	r.Run()
}
