package users

import (
	"net/http"
	"time"

	"github.com/PandaFarmer/CloverRoad/backend/pkg/common/models"
	"github.com/PandaFarmer/CloverRoad/backend/pkg/core"
	"github.com/gin-gonic/gin"
)

type AddUserRequestBody struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	Email       string `json:"email"`
	UserName    string `json:"user_name"`
	FullName    string `json:"full_name"`
	Password    string `json:"password"`
	IsSuperUser bool   `json:"is_super_user"`
}

func (h handler) AddUser(c *gin.Context) {
	body := AddUserRequestBody{}

	// parse body, attach to AddUserRequestBody struct
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user models.User

	user.Id = body.Id
	user.Email = body.Email
	user.UserName = body.UserName
	// user.FirstName = body.FirstName
	// user.Surname = body.Surname
	user.FullName = body.FullName
	user.Password = core.HashAndSalt([]byte(body.Password))
	user.DateJoined = time.Now()
	user.IsSuperUser = body.IsSuperUser

	// insert new db entry
	if result := h.DB.Create(&user); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &user)
}
