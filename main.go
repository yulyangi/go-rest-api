package main

import (
	"example.com/go-rest-api/db"
	"example.com/go-rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	db.CreateTables()

	router := gin.Default()
	routes.RegisterRoutes(router)

	router.Run("localhost:8080")
}
