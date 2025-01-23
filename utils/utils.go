package utils

import (
	"bytes"
	"image"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"

	"golang.org/x/image/webp"
)

func ProcessImage(url string) int {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Failed to download image from URL %s: %v", url, err)
		return -1
	}
	defer resp.Body.Close()

	imgData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read image data from URL %s: %v", url, err)
		return -1
	}

	var img image.Image
	img, _, err = image.Decode(bytes.NewReader(imgData))
	if err != nil {
		img, err = webp.Decode(bytes.NewReader(imgData))
		if err != nil {
			log.Printf("Failed to decode WebP image from URL %s: %v", url, err)
			return -1
		}
	}
	height := img.Bounds().Dy() 
	width := img.Bounds().Dx()  
	perimeter := 2 * (height + width)

	time.Sleep(time.Duration(rand.Intn(300)+100) * time.Millisecond)

	log.Printf("Processed image from URL %s: Height=%d, Width=%d, Perimeter=%d", url, height, width, perimeter)

	return perimeter
}