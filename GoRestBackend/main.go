package main

import (
	router "github.com/NirmalVatsyayan/GoRestBackend/Router"
)

func init() {

}

func StartServer() {
	routes := router.RouteDispatcher()
	routes.Run(":10000")
}

func main() {
	StartServer()
}
