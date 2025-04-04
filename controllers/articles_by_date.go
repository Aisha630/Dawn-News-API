package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/abdulalikhan/Dawn-News-API/models"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
)

// Get : Return news articles by Date for Pakistan
// @Summary Get news by date for Pakistan
// @Description Fetches news articles published on a specific date (format: YYYY-MM-DD) in Pakistan
// @Tags News
// @Produce json
// @Param date path string true "Date in YYYY-MM-DD format"
// @Success 200 {object} []models.Article
// @Router /date/{date} [get]
func FetchNewsOnDate(c *gin.Context) {
	articles := make([]models.ArticleDetails, 0)

	date := c.Param("date")
	if date == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Date parameter is required"})
		return
	}

	cly := colly.NewCollector(
		colly.AllowedDomains("www.dawn.com", "dawn.com"),
	)

	cly.OnHTML("article", func(e *colly.HTMLElement) {
		headline := e.ChildText(".story__link  ")
		link := e.ChildAttr(".story__link  ", "href")
		publishTime := e.ChildAttr(".timestamp--time", "title")
		excerpt := e.ChildText(".story__excerpt")
		layout := "2006-01-02T15:04:05+05:00"
		publishTimeParsed, _ := time.Parse(layout, publishTime)
		if len(publishTime) > 0 {
			article := models.ArticleDetails{
				Headline:    headline,
				URL:         link,
				PublishTime: publishTimeParsed.Format("2006-01-02 03:04 PM"),
				Excerpt:     excerpt,
			}
			articles = append(articles, article)
		}
	})
	cly.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	
	if err:= cly.Visit("https://www.dawn.com/pakistan/" + date); err != nil {
		fmt.Printf("Error visiting URL: %v\n", err)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")

	// Dump json to the standard output
	//enc.Encode(articles)
	c.JSON(http.StatusOK, articles)
}
