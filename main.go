package main

import (
	"github.com/beGG1/study_program_api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
