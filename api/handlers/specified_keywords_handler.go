package api

import (
	"container/heap"
	"net/http"

	"github.com/yjennykim/seo-analyzer/pkg/models"

	"github.com/gin-gonic/gin"
)

func GetSpecifiedWordDensities(c *gin.Context) {
	doc, err := getDoc(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	searchWords, err := getSearchTerms(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// count word frequencies
	keyWordsCounter := make(map[string]int)
	totalWords := searchAllKeywords(doc, &keyWordsCounter)

	// put words into heap
	h := &pkg.MaxHeap{}
	heap.Init(h)

	for word, count := range keyWordsCounter {
		// add to heap if it's a keyword
		_, ok := searchWords[word]
		if ok {
			density := float64(count) / float64(totalWords) * 100
			heap.Push(h, pkg.KeyWords{
				Word:           word,
				KeyWordDensity: density,
				Frequency:      count,
			})
			delete(searchWords, word)
		}
	}

	// add search words that were missing from search
	for word := range searchWords {
		heap.Push(h, pkg.KeyWords{
			Word:           word,
			KeyWordDensity: 0,
			Frequency:      0,
		})
	}

	// add words to result
	var sortedKeywords []pkg.KeyWords
	for h.Len() > 0 {
		word := heap.Pop(h).(pkg.KeyWords)
		sortedKeywords = append(sortedKeywords, word)
	}

	c.IndentedJSON(http.StatusOK, gin.H{"keywordData": sortedKeywords})
}
