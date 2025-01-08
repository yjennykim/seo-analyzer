package pkg

type KeyWords struct {
	Word           string  `json:"word"`
	KeyWordDensity float64 `json:"keyword density"`
	Frequency      int     `json:"word frequency"`
}

type MaxHeap []KeyWords

// MaxHeap methods
func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i].Frequency > h[j].Frequency
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(KeyWords))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
