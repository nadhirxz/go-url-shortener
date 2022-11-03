package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nadhirxz/go-url-shortener/model"
	"github.com/nadhirxz/go-url-shortener/utils"
)

func Init() {
	router := gin.Default()

	router.GET("/shorts", func(c *gin.Context) {
		shorts, err := model.GetShorts()

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}

		c.IndentedJSON(http.StatusOK, shorts)
	})

	router.POST("/short", func(c *gin.Context) {
		var short model.Short

		err := c.BindJSON(&short)

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
			return
		}

		short.Short = utils.RandomString(20)

		fmt.Println(short)

		err = model.CreateShort(short)

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{
			"url":   short.Url,
			"short": utils.GenerateFullURL(short.Short),
		})
	})

	router.GET("/r/:url", func(c *gin.Context) {
		url := c.Param("url")

		short, err := model.GetShortByURL(url)

		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
			return
		}

		c.Redirect(http.StatusTemporaryRedirect, short.Url)
	})

	router.Run("localhost:" + os.Getenv("PORT"))
}
