package main

import (
	"fmt"

	"github.com/hunzo/go-gin-swagger/docs"
	"github.com/hunzo/go-gin-swagger/routers"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// gin-swagger middleware
// swagger embed files

func main() {
	fmt.Println("TEST")
	docs.SwaggerInfo.Title = "TEST Swagger "
	docs.SwaggerInfo.Description = "My Swagger"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	// docs.SwaggerInfo.Schemes = []string{"http", "https"}

	s := routers.InitRouter()

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	s.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	s.Run()

}
