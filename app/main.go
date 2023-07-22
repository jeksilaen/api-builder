package main

import (
	"github.com/gin-gonic/gin"
	db "github.com/jeksilaen/api-builder/db"
	middlewares "github.com/jeksilaen/api-builder/middlewares"
	handler "github.com/jeksilaen/api-builder/modules/user/handlers"
)

func main() {
	err := db.InitDB()
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.Use(middlewares.SetJSONContentTypeMiddleware())

	handler.InitUserHttpHandler(router)

	router.Run("localhost:8080")
}
