package utils

import (
	"image"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func ProcessImage(url string) int {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Failed to download image from URL %s: %v", url, err)
		return -1
	}
	defer resp.Body.Close()

	img, _, err := image.DecodeConfig(resp.Body)
	if err != nil {
		log.Printf("Failed to decode image from URL %s: %v", url, err)
		return -1
	}

	height := img.Height
	width := img.Width
	perimeter := 2 * (height + width)

	time.Sleep(time.Duration(rand.Intn(300)+100) * time.Millisecond)
	return perimeter
}
