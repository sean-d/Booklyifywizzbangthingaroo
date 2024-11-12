package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sean-d/Booklyifywizzbangthingaroo/db"
	"github.com/sean-d/Booklyifywizzbangthingaroo/routes"
)

func main() {
	db.InitDB()
	db.CreateTables()
	server := gin.Default() // returns a pointer

	routes.RegisterRoutes(server) // since gin.Default returns a pointer, no need explicitly pass pointer.
	err := server.Run(":8080")

	if err != nil {
		panic(err)
	}
}
