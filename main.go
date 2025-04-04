package main

import (
	"log"
	"github.com/abdulalikhan/Dawn-News-API/controllers"
	docs "github.com/abdulalikhan/Dawn-News-API/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Dawn News Web Scraper API
// @version 1.0
// @license.name Apache 2.0
// @description An API that fetches news articles from Dawn.com
// @host localhost:3000
// @contact.name Abdul Ali Khan
// @BasePath /api/v1
func main() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := r.Group("/api/v1")
	{
		v1.GET("/latest-news", controllers.FetchLatestNews)
		v1.GET("/article/:id", controllers.FetchArticle)
		v1.GET("/date/:date", controllers.FetchNewsOnDate)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	if err := r.Run(":3000"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
