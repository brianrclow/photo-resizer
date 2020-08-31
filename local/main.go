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
	Resizes photo to width iPhone 11 Pro 2436 pixels width and scales the rest of the photo proportionally. 
	This is version that can run locally.
*/
func main() {
	// reads the directory given below
	photos, err := ioutil.ReadDir("./input")
	if err != nil {
		fmt.Println("error reading directory")
		log.Fatal(err)
	}
	// resizes all photos in the path
	for _, f := range photos {

		// open photo
		file, err := os.Open("./input/"+f.Name())
		if err != nil {
			fmt.Println("error opening photo and saving it as a file")
			log.Fatal(err)
		}

		// decode jpeg into image.Image
		img, err := jpeg.Decode(file)
		if err != nil {
			fmt.Println("error decoding jpeg")
			log.Fatal(err)
		}
		file.Close()

		// resize to width 2436 using Lanczos resampling and preserve aspect ratio
		m := resize.Resize(2436, 0, img, resize.Lanczos3)

		out, err := os.Create("./output/"+f.Name())
		if err != nil {
			fmt.Println("error creating outputfile")
			log.Fatal(err)
		}
		defer out.Close()

		// write new image to file
		jpeg.Encode(out, m, nil)
		fmt.Println(f.Name() + " Done!")
	}
}