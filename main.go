package main

import (
	"example/ginference-server/config/devconfig"
	"example/ginference-server/routes"
)

// @title           Swagger ginference-server API
// @version         1.0
// @description     This is a GO REST server for sentix inference.
// @termsOfService  http://swagger.io/terms/

// @contact.name   chethanMysore
// @contact.url    https://chethanmysore.github.io/portfolio/#contact
// @contact.email  willishardrock94@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description "Type 'Bearer TOKEN' to correctly set the API Key"
// @authorizationurl http://localhost:8080/api/v1/auth/login
// @tokenUrl http://localhost:8080/api/v1/auth/login

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	router := routes.Init()
	router.Run(devconfig.APIDomainURI)
}
