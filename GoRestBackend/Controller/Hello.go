package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"Status": "OK",
	})

}

func UserDetail(c *gin.Context) {

	userid := c.Param("userid")
	utm_source := c.Query("utm_source")

	c.JSON(http.StatusOK, gin.H{
		"UserID":    userid,
		"UTMSource": utm_source,
	})

}
