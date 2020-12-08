package main

import (
	"login/routers"
	"login/utils"
)

func main() {
	router := routers.InitRoute()
	port := utils.EnvVar("SERVER_PORT", ":8080")
	// fmt.Println(port)
	router.Run(port)
}
