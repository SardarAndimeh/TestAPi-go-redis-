package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(server *gin.Engine) {

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Homepage"})
	})
	server.POST("/user", createUser)
	server.GET("/user", showUser)

}
