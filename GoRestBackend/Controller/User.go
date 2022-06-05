package controller

import (
	"net/http"
	"time"

	model "github.com/NirmalVatsyayan/GoRestBackend/Database/Model"
	input "github.com/NirmalVatsyayan/GoRestBackend/Input"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
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

	count, err := model.CountUser(register_input.Email)
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email already registered",
		})
		return
	}

	var user model.User
	user.ID = uuid.NewV4().String()
	user.FirstName = register_input.FirstName
	user.LastName = register_input.LastName
	user.Email = register_input.Email
	user.Password = register_input.Password
	user.Created = time.Now()
	user.Updated = time.Now()

	err = model.InsertUser(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error in registration",
		})
		return
	}

	c.JSON(http.StatusOK, register_input)

}

func UserUpdate(c *gin.Context) {

	var register_input input.Register
	err := c.BindJSON(&register_input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid form data",
		})
		return
	}

	var user model.User
	_ = model.GetUser(&user, register_input.Email)

	updateData := make(map[string]interface{})
	updateData["FirstName"] = register_input.FirstName
	updateData["LastName"] = register_input.LastName

	_ = model.UpdateUser(&user, &updateData)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
