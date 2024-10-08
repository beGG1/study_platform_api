package main

import (
	"github.com/beGG1/study_platform_api/database"
	"github.com/beGG1/study_platform_api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
