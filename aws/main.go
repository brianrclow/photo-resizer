package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/s3"
    "github.com/aws/aws-sdk-go/service/s3/s3manager"
)

/*
	Resizes photo to width iPhone 11 Pro 2436 pixels width and scales the rest of the photo proportionally. 
	This has been adapted to test Go and also to use with AWS Lambda.
*/
func main() {
	lambda.Start(HandleRequest)
}

func exitErrorf(msg string, args ...interface{}) {
    fmt.Fprintf(os.Stderr, msg+"\n", args...)
    os.Exit(1)
}

func HandleRequest(ctx context.Context) (error) {
	var bucket = os.Getenv("imageBucket")
	fmt.Println(bucket)
	var item = os.Getenv("imageInputDirectory")


	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)
	
	downloader := s3manager.NewDownloader(sess)

	photos, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(item),
		})
	if err != nil {
		exitErrorf("Unable to download item %q, %v", item, err)
	}


	// reads the directory given below
//		photos, err := ioutil.ReadDir("input/")
//	if err != nil {
//		fmt.Println("error reading directory")
//		log.Fatal(err)
//	}
	// resizes all photos in the path
	for _, f := range photos {

		// open photo
		file, err := os.Open("input/"+f.Name())
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

		out, err := os.Create("output/"+f.Name())
		if err != nil {
			fmt.Println("error creating outputfile")
			log.Fatal(err)
		}
		defer out.Close()

		// write new image to file
		jpeg.Encode(out, m, nil)
		fmt.Println(f.Name() + " Done!")
	}
	return nil
}