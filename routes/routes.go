package routes

import (
	config "example/ginference-server/config/devconfig"
	"example/ginference-server/controllers/authcontroller"
	"example/ginference-server/controllers/modelcontroller"
	"example/ginference-server/controllers/usercontroller"
	"example/ginference-server/middlewares"
	"time"

	docs "example/ginference-server/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init() *gin.Engine {
	router := gin.Default()

	cnf := cors.DefaultConfig()
	cnf.AllowOrigins = config.CORSAllowedDomains
	// cnf.AllowAllOrigins = config.CORSAllowAllOrigins
	cnf.AllowMethods = config.CORSAllowedMethods
	cnf.AllowHeaders = config.CORSAllowedHeaders
	cnf.ExposeHeaders = config.CORSExposedHeaders
	cnf.AllowCredentials = config.CORSAllowCredentials
	cnf.MaxAge = time.Duration(config.CORSMaxAge) * time.Hour
	router.Use(cors.New(cnf))

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("/api/v1")
	auth := v1.Group("/auth")
	{
		auth.POST("/register", authcontroller.Register)
		auth.GET("/login", authcontroller.Login)
	}
	v1.Use(middlewares.JWTAuthMiddleware())
	{
		// User routes '/users*'
		usr := v1.Group("/users")
		{
			usr.GET("/all", usercontroller.GetAllUsers)
			usr.GET("/id/:id", usercontroller.GetUserByID)
			usr.GET("/name/:name", usercontroller.GetUserByName)
			usr.GET("/username/:username", usercontroller.GetUserByUserName)
			usr.GET("/auth/id/:id", usercontroller.GetUserRoleByID)
			usr.GET("/auth/user", usercontroller.GetAuthUser)
			usr.PUT("/edit", usercontroller.EditUser)
		}
		// Model routes '/models*'
		mod := v1.Group("/models")
		{
			mod.GET("/all", modelcontroller.GetAllModels)
			mod.GET("/id/:id", modelcontroller.GetModelByID)
			mod.GET("/name/:name", modelcontroller.GetModelByName)
			mod.GET("/username/:username", modelcontroller.GetModelsByUsername)
			mod.POST("/create", modelcontroller.CreateNewModel)
			mod.PUT("/edit", modelcontroller.EditModel)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
