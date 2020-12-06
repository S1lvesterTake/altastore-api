package main

import (
	"altastore-api/infrastructure/persistence/repository/db"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "*"},
		ExposeHeaders:    []string{"Accept", "Content-Length", "Content-Type", "Authorization", "Accept:Encoding"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	v1 := router.Group("/api/v1") // initial route
	db := db.DBInit()             // initial db configuration

	// Health check
	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "I'm well",
		})
	})

	//login route
	loginRoute(v1, db)

	///// end of route ////

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal(fmt.Sprintf("PORT must be set [%s]", port))
	}

	router.Run(":" + port)
}

func loginRoute(route *gin.RouterGroup, db *gorm.DB) {
	handler := LoginHandler(db)

	v1 := route.Group("/login")
	{
		v1.POST("", handler.LoginHandler)
	}
}
