package users

import (
	"net/http"

	"github.com/PandaFarmer/CloverRoad/backend/pkg/common/models"
	"github.com/PandaFarmer/CloverRoad/backend/pkg/core"
	"github.com/gin-gonic/gin"
)

type UpdateUserRequestBody struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`
	// FirstName   string `json:"first_name"`
	// Surname	string `json:"surname`
	FullName    string `json:"full_name"`
	Password    string `json:"password"`
	IsSuperUser bool   `json:"is_super_user"`
}

func (h handler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	body := UpdateUserRequestBody{}

	// getting request's body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user models.User

	if result := h.DB.First(&user, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	// user.Email = body.Email
	// user.UserName = body.UserName
	// user.FullName = body.FullName
	user.Password, _ = core.HashAndSalt([]byte(body.Password))
	// user.IsSuperUser = body.IsSuperUser

	// save user
	h.DB.Save(&user)

	c.Status(http.StatusOK)
}
