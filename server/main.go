package main

import (
	"github.com/gin-gonic/gin"
	"github.com/goameer030/732121104037/server/controllers"
	"github.com/goameer030/732121104037/server/initializers"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	engine := gin.Default()

	engine.GET("/ping", controllers.Ping)
	engine.GET("/categories/:category/product", controllers.GetProductsOnCategory)

	engine.Run()
}
