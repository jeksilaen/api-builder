package main

import (
	// "log"
	"os"

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

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})

	// router.Use(cors.Default())
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
