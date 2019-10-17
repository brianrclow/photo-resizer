package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
)

/*
	Resizes photo to width 1500 pixels and scales the rest of the photo proportionally.
*/
func main() {
	// reads the directory given below
	photos, err := ioutil.ReadDir("./weekend")
	if err != nil {
		log.Fatal(err)
	}
	// resizes all photos in the path
	for _, f := range photos {

		// open photo
		file, err := os.Open("./input/"+f.Name())
		if err != nil {
			log.Fatal(err)
		}

		// decode jpeg into image.Image
		img, err := jpeg.Decode(file)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()

		// resize to width 1500 using Lanczos resampling and preserve aspect ratio
		m := resize.Resize(1500, 0, img, resize.Lanczos3)

		out, err := os.Create("./output/"+f.Name())
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		// write new image to file
		jpeg.Encode(out, m, nil)
		fmt.Println(f.Name())
	}
}
