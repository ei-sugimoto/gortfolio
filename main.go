package main

import (
	"github.com/ei-sugimoto/gortfolio/ent"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	client := ent.Migrate()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	r.GET("/qiita", func(c *gin.Context) {
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
