package main

import (
	db "github.com/NirmalVatsyayan/GoRestBackend/Database/MongoConn"
	router "github.com/NirmalVatsyayan/GoRestBackend/Router"
)

func init() {
	db.InitializeMongoConn("localhost", 27017, "", "DemoRestBackend", "")

}

func StartServer() {
	routes := router.RouteDispatcher()
	routes.Run(":10000")
}

func main() {
	StartServer()
}
