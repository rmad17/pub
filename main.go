
package main

import (
  "github.com/gin-gonic/gin"

  "pub/controllers"
  "pub/models"

)

var router *gin.Engine

func main() {
    models.ConnectDatabase()

    router := gin.Default()

    api := router.Group("/api")

	api.POST("/register", controllers.Register)

    router.Run()

}
