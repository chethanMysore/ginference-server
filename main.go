package main

import (
	"example/ginference-server/config/devconfig"
	"example/ginference-server/routes"
)

func main() {
	router := routes.Init()
	router.Run(devconfig.APIDomainURI)
}
