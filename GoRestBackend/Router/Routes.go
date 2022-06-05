package router

import (
	controller "github.com/NirmalVatsyayan/GoRestBackend/Controller"
	"github.com/gin-gonic/gin"
)

func RouteDispatcher() *gin.Engine {

	gin.SetMode(gin.DebugMode)
	//gin.SetMode(gin.ReleaseMode) only for production

	// create gin engine
	router := gin.New()

	// add IP of load balancer
	router.SetTrustedProxies([]string{"127.0.0.1"})

	// add middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// create apis
	apisV1 := router.Group("/v1")
	apisV1.GET("/hello", controller.Hello)

	// get query param and url param
	apisV1.GET("/user/:userid", controller.UserDetail)
	apisV1.POST("/register", controller.UserRegister)

	// return engine with route defined
	return router

}
