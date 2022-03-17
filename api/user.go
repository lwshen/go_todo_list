package api

import (
	"github.com/gin-gonic/gin"
	"todo_list/service"
)

func UserRegister(c *gin.Context) {
	var userRegister service.UserService
	if err := c.ShouldBind(&userRegister); err == nil {
		res := userRegister.Register()
		c.JSON(200, res)
	} else {
		c.JSON(400, err.Error())
	}
}
