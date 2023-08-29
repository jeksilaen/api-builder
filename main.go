package main

import (
	// "log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	db "github.com/jeksilaen/api-builder/db"
	middlewares "github.com/jeksilaen/api-builder/middlewares"
	collectionHandler "github.com/jeksilaen/api-builder/modules/collection/handlers"
	requestHandler "github.com/jeksilaen/api-builder/modules/request/handlers"
	userHandler "github.com/jeksilaen/api-builder/modules/user/handlers"
	// "github.com/joho/godotenv"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	err := db.InitDB()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	router.Use(cors.New(corsConfig))

	router.Use(middlewares.SetJSONContentTypeMiddleware())

	userHandler.InitUserHttpHandler(router)
	collectionHandler.InitCollectionHttpHandler(router)
	requestHandler.InitRequestHttpHandler(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run("0.0.0.0:" + port)
}
