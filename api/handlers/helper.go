package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode"

	"github.com/yjennykim/seo-analyzer/pkg/constants"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

func getDoc(c *gin.Context) (*goquery.Document, error) {
	url := c.DefaultQuery("url", "")

	if url == "" {
		return nil, fmt.Errorf("URL is required")
	}

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode > 400 {
		fmt.Println("Status code:", response.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func getK(c *gin.Context) (int, error) {
	input := c.DefaultQuery("topK", "10")
	k, err := strconv.Atoi(input)

	if err != nil {
		return -1, err
	}

	return k, nil
}

func searchAllKeywords(doc *goquery.Document, keyWordsCounter *map[string]int) int {
	totalWords := 0
	doc.Find("h1, h2, title, p").Each(func(i int, s *goquery.Selection) {
		text := strings.ToLower(s.Text())
		words := strings.FieldsFunc(text, func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsNumber(r)
		})

		for _, word := range words {
			totalWords += 1
			if !pkg.StopWords[word] && len(word) > 1 {
				(*keyWordsCounter)[word]++
			}
		}
	})

	return totalWords
}

func getSearchTerms(c *gin.Context) (map[string]bool, error) {
	input := c.DefaultQuery("keywords", "")

	if input == "" {
		return nil, fmt.Errorf("please specify search words")
	}

	keywords := strings.Split(input, ",")

	keywordsSet := map[string]bool{}
	for _, word := range keywords {
		key := strings.ToLower(strings.TrimSpace(word))
		keywordsSet[key] = true
	}

	return keywordsSet, nil
}
