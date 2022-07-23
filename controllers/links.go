package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"linkShortener/models"
	"net/http"
	"strconv"
)

func checkAuth(authToken string) bool {
	result := false
	if viper.GetString("Token") == authToken {
		result = true
	}
	return result
}

func GetURL(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if !checkAuth(auth) {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}
	shortId := c.Param("id")
	var shortIdUint uint64
	shortIdUint, _ = strconv.ParseUint(shortId, 10, 64)
	var url = models.FetchUrlCached(uint(shortIdUint))
	c.Redirect(http.StatusMovedPermanently, url)
}

func ShortenURL(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	if !checkAuth(auth) {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
	}

	url := c.PostForm("url")
	shortId := models.SaveUrl(url)
	c.IndentedJSON(http.StatusOK, gin.H{"short": viper.GetString("BaseUrl") + strconv.FormatUint(uint64(shortId), 10)})
	return
}
