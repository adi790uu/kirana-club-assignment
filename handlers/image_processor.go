package handlers

import (
	"math/rand"
	"time"
)

func ProcessImage(url string) int {
	time.Sleep(time.Duration(rand.Intn(300)+100) * time.Millisecond)
	return 2 * (100 + 50)
}
