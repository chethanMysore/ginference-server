package routes

import (
	"example/ginference-server/routes/modelroutes"
	"example/ginference-server/routes/userroutes"

	docs "example/ginference-server/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init() *gin.Engine {
	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("/api/v1")
	{
		// User routes '/users*'
		usr := v1.Group("/users")
		{
			usr.GET("/", userroutes.GetAllUsers)
			usr.GET("/id/:id", userroutes.GetUserByID)
			usr.GET("/name/:name", userroutes.GetUserByName)
			usr.GET("/username/:username", userroutes.GetUserByUserName)
			usr.POST("/create", userroutes.CreateNewUser)
			usr.PUT("/edit", userroutes.EditUser)
		}
		// Model routes '/models*'
		mod := v1.Group("/models")
		{
			mod.GET("/", modelroutes.GetAllModels)
			mod.GET("/id/:id", modelroutes.GetModelByID)
			mod.GET("/name/:name", modelroutes.GetModelByName)
			mod.GET("/username/:username", modelroutes.GetModelsByUsername)
			mod.POST("/create", modelroutes.CreateNewModel)
			mod.PUT("/edit", modelroutes.EditModel)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
