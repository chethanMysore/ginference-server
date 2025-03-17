package modelroutes

import (
	"net/http"

	"example/ginference-server/data"

	"github.com/gin-gonic/gin"
)

// GET (/models)
func GetAllModels(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data.SubscribedModels)
}

// GET (/models/id/:id)
func GetModelByID(c *gin.Context) {
	id := c.Param("id")
	model, err := data.SubscribedModels.FindByUUID(id)
	if err == nil {
		c.IndentedJSON(http.StatusOK, model)
	} else {
		c.JSON(http.StatusNotFound, err)
	}
}

// GET (/models/name/:name)
func GetModelByName(c *gin.Context) {
	name := c.Param("name")
	model, err := data.SubscribedModels.FindByName(name)
	if err == nil {
		c.IndentedJSON(http.StatusOK, model)
	} else {
		c.JSON(http.StatusNotFound, err)
	}
}

// GET (/models/username/:username)
func GetModelsByUsername(c *gin.Context) {
	username := c.Param("username")
	user, err := data.RegisteredUsers.FindByName(username)
	if err == nil {
		model, err := data.SubscribedModels.FindByUser(user.UserID.String())
		if err == nil {
			c.IndentedJSON(http.StatusOK, model)
		} else {
			c.JSON(http.StatusNotFound, err)
		}
	} else {
		c.JSON(http.StatusNotFound, err)
	}
}
