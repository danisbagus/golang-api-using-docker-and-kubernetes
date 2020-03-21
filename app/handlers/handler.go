package handler

import (
	"fmt"
	"go-project/config"
	"go-project/models"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {

	db, err:= config.DBInit()
	_user := []models.Users{}

	if err != nil {
		c.JSON(404,err)
	}

	if err := db.Find(&_user).Error; err != nil {
	   c.AbortWithStatus(404)
	   fmt.Println(err)
	} else {
	   c.JSON(200, _user)
	}
	defer db.Close()

}


