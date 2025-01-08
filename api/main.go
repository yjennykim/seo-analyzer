package api

import (
	"net/http"

	"github.com/yjennykim/seo-analyzer/api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to SEO Analyzer",
		})
	})

	// required: url -- page to scrape
	// optional: topK -- get topK keywords (default 10)
	router.GET("/topKeywords", api.GetTopKWordDensities)

	// required: url -- page to scrape
	// optional: keywords -- get frequency of the comma-separated keywords
	router.GET("/keywordsFrequencies", api.GetSpecifiedWordDensities)

	router.Run(":8080")
}
