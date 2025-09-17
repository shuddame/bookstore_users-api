package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shuddame/bookstore_users-api/domain/users"
	"github.com/shuddame/bookstore_users-api/services"
	"github.com/shuddame/bookstore_users-api/utils/errors"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}
func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		restErr := errors.NewBadRequestError("invalid user id")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, result)
}
func FindUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
func UpdateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
func DeleteUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me")
}
