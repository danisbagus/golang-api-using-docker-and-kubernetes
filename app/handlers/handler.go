package handler

import (
	"fmt"
	"time"
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

	data := new(Credential)
	err := c.Bind(data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "can't bind struct",
		})
		c.Abort() 
	}

	db, err:= config.DBInit()
	_user := models.Users{}

	if err != nil {
		c.JSON(404,err)
	}

	if err := db.Where("username = ?", data.Username).First(&_user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}

	 defer db.Close()
	 
	if data.Username != _user.Username {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "wrong username or password",
		})
		c.Abort() 

	} else {
		if data.Password != _user.Password {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "wrong username or password",
			})
			c.Abort() 
		}
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = _user.ID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	sign := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// token, err := sign.SignedString([]byte(os.Getenv("API_SECRET")))
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

