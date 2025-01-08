package api

import (
	"container/heap"
	"net/http"

	"github.com/yjennykim/seo-analyzer/pkg/models"

	"github.com/gin-gonic/gin"
)

func GetTopKWordDensities(c *gin.Context) {
	doc, err := getDoc(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	topK, err := getK(c)
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
		density := float64(count) / float64(totalWords) * 100
		heap.Push(h, pkg.KeyWords{
			Word:           word,
			KeyWordDensity: density,
			Frequency:      count,
		})
	}

	// get topK words
	var sortedKeywords []pkg.KeyWords
	for h.Len() > 0 && topK > 0 {
		sortedKeywords = append(sortedKeywords, heap.Pop(h).(pkg.KeyWords))
		topK--
	}

	c.IndentedJSON(http.StatusOK, gin.H{"keywordData": sortedKeywords})
}
