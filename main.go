package main

import (
	"example/ginference-server/routes/modelroutes"
	"example/ginference-server/routes/userroutes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/users", userroutes.GetAllUsers)
	router.GET("/users/id/:id", userroutes.GetUserByID)
	router.GET("/users/name/:name", userroutes.GetUserByName)
	router.GET("/models", modelroutes.GetAllModels)
	router.GET("/models/id/:id", modelroutes.GetModelByID)
	router.GET("/models/name/:name", modelroutes.GetModelByName)
	router.GET("/models/username/:username", modelroutes.GetModelsByUsername)
	router.Run("localhost:8080")
}
