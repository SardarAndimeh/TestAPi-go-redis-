package route

import (
	"net/http"
	"redis/demo/db"
	"redis/demo/model"

	"github.com/gin-gonic/gin"
)

func createUser(c *gin.Context) {
	var user model.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "saved user successfully"})

}

func showUser(c *gin.Context) {
	key := "user:2"
	data, err := db.Rdb.HGetAll(db.Ctx, key).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not get the user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "got it", "user": data})
}
