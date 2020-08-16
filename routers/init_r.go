package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hunzo/go-gin-swagger/controller"
)

func InitRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/users/:username", controller.GetAll())
	r.POST("/somepost/", controller.SomePost())

	return r

}
