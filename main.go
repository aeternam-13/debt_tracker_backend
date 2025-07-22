package main

import (
	"debt_tracker/database"
	"debt_tracker/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")
}
