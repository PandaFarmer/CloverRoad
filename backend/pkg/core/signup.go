// controllers/public.go

package core

import (
	"log"
	"net/http"

	"github.com/PandaFarmer/CloverRoad/backend/pkg/common/models"
	"github.com/gin-gonic/gin"
)

// Signup creates a user in db
func (h handler) Signup(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Println(err)

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "invalid json",
		})
		c.Abort()

		return
	}

	hashed_password, err := HashAndSalt([]byte(user.Password))
	if err != nil {
		log.Println(err.Error())

		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "error hashing password",
		})
		c.Abort()

		return
	}

	user.Password = hashed_password

	if result := h.DB.Create(&user); result.Error != nil {
		log.Println(result.Error)
		c.JSON(500, gin.H{
			"msg": "error creating user",
		})
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	c.JSON(http.StatusBadRequest, user)
}
