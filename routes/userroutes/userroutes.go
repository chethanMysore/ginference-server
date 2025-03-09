package userroutes

import (
	"net/http"

	"example/ginference-server/data"

	"github.com/gin-gonic/gin"
)

// GET (/users)
func GetAllUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data.RegisteredUsers)
}

// GET (/users/id/:id)
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := data.RegisteredUsers.FindByUUID(id)
	if err == "" {
		c.IndentedJSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusNotFound, err)
	}
}

// GET (/users/name/:name)
func GetUserByName(c *gin.Context) {
	name := c.Param("name")
	user, err := data.RegisteredUsers.FindByName(name)
	if err == "" {
		c.IndentedJSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusNotFound, err)
	}
}
