package routes

import (
	"example/ginference-server/routes/modelroutes"
	"example/ginference-server/routes/userroutes"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()

	// User routes '/users*'
	router.GET("/users", userroutes.GetAllUsers)
	router.GET("/users/id/:id", userroutes.GetUserByID)
	router.GET("/users/name/:name", userroutes.GetUserByName)
	router.GET("/users/username/:username", userroutes.GetUserByUserName)
	router.POST("/users/create", userroutes.CreateNewUser)
	router.PUT("/users/edit", userroutes.EditUser)

	// Model routes '/models*'
	router.GET("/models", modelroutes.GetAllModels)
	router.GET("/models/id/:id", modelroutes.GetModelByID)
	router.GET("/models/name/:name", modelroutes.GetModelByName)
	router.GET("/models/username/:username", modelroutes.GetModelsByUsername)
	router.POST("/models/create", modelroutes.CreateNewModel)
	router.PUT("/models/edit", modelroutes.EditModel)

	return router
}
