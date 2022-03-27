package main

import (
	"github.com/kbinani/screenshot"
	"image/png"
	"os"
)

func main() {
	n := screenshot.NumActiveDisplays()

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		imagen, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		FileName := "c:\\users\\public\\fotos.png"
		file, _ := os.Create(FileName)
		defer file.Close()
		png.Encode(file, imagen)

	}
}
