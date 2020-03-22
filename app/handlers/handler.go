package handler

import (
	"fmt"
	"go-project/config"
	"go-project/models"

	"net/http"

	"github.com/gin-gonic/gin"

	jwt "github.com/dgrijalva/jwt-go"
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

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
  }

func Login(c *gin.Context){

	user := new(Credential)
	err := c.Bind(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "can't bind struct",
		})
	}
	 
	if user.Username != "danisbagus" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "wrong username or password",
		})
	} else {
		if user.Password != "12345" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "wrong username or password",
			})
		}
	}

	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	token, err := sign.SignedString([]byte("secret"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
    
}

