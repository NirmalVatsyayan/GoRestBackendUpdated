package controller

import (
	"net/http"

	input "github.com/NirmalVatsyayan/GoRestBackend/Input"
	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {

	var register_input input.Register

	err := c.BindJSON(&register_input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid form data",
		})
		return
	}

	c.JSON(http.StatusOK, register_input)

}
