package main

import (
	"redis/demo/db"
	"redis/demo/route"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Redis
	db.InitRedis()
	server := gin.Default()

	route.Routes(server)
	server.Run(":8080")

}
