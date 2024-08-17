package main

import (
	"time"

	"github.com/ei-sugimoto/gortfolio/ent"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000, http://front:3000"}, // 許可するオリジンを指定
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	client := ent.Migrate()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	r.GET("/article", func(c *gin.Context) {
		items, err := GetArticle(client, c)

		if err != nil {
			c.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
		}
		c.JSON(200, items)
	})

	r.Run(":8080")

}
