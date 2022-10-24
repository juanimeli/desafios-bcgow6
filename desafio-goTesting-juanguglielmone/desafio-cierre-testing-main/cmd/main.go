package main

import (
	"os"

	"github.com/bootcamp-go/desafio-cierre-testing/cmd/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.MapRoutes(r)

	if r.Run(":18085") != nil {
		os.Exit(1)
	}

}
