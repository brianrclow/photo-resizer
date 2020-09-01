package main

import (
	"fmt"
	"github.com/nfnt/resize"
	"image/jpeg"
	"log"
	"os"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
    "github.com/aws/aws-sdk-go/service/s3/s3manager"
)

/*
	Resizes photo to width iPhone 11 Pro 2436 pixels width and scales the rest of the photo proportionally. 
	This is version that can be run on AWS Lambda.
*/
func main() {
	lambda.Start(HandleRequest)
}

// for error logging
func exitErrorf(msg string, args ...interface{}) {
    fmt.Fprintf(os.Stderr, msg+"\n", args...)
    os.Exit(1)
}

func HandleRequest() {

// download from S3 input directory
	var imageSrcPath = os.Getenv("imageSrcPath")
	// fmt.Println("imageSrcPath: " + imageSrcPath)
	var originalImage = os.Getenv("imageName")
	// fmt.Println("imageName: " + originalImage)
	var imageDestinationPath = os.Getenv("imageDestinationPath")
	// fmt.Println("imageDestinationPath: " + imageDestinationPath)

	// create the image to be downloaded
	unprocessedImage, err := os.Create("/tmp/"+originalImage)
    if err != nil {
        exitErrorf("Unable to open file %q, %v", originalImage, err)
	}
	
	// initializes session in us-west-2 with credentials
	sessDn, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)
	// creates new downloader object
	downloader := s3manager.NewDownloader(sessDn)
	

	// downloads the unprocessed image from the imageSrcPath
	numBytes, err := downloader.Download(unprocessedImage, &s3.GetObjectInput{
			Bucket: aws.String(imageSrcPath),
			Key:    aws.String(originalImage),
		})
	if err != nil {
		exitErrorf("Unable to download originalImage %q, %v", originalImage, err)
	}
	fmt.Println("Downloaded", unprocessedImage.Name(), numBytes, "bytes")
	fmt.Println("Download Success!")


// resize in lambda

	fmt.Println("Starting resize...")


		// open image
		openedImage, err := os.Open(unprocessedImage.Name())
		if err != nil {
			fmt.Println("error opening photo and saving it as a file")
			log.Fatal(err)
		}

		// decode jpeg into image.Image
		decodedImage, err := jpeg.Decode(openedImage)
		if err != nil {
			fmt.Println("error decoding jpeg")
			log.Fatal(err)
		}
		openedImage.Close()

		// resize to width 2436 using Lanczos resampling and preserve aspect ratio
		resizedImage := resize.Resize(2436, 0, decodedImage, resize.Lanczos3)

		// create output file for image using original name
		out, err := os.Create(unprocessedImage.Name())
		if err != nil {
			fmt.Println("error creating outputfile")
			log.Fatal(err)
		}
		defer out.Close()

		// write new image to file
		jpeg.Encode(out, resizedImage, nil)
		fmt.Println(unprocessedImage.Name() + " resized")
		fmt.Println("Resize Success!")


// upload to S3 output directory
	fmt.Println("Starting upload to S3...")


	processedImage, err := os.Open(unprocessedImage.Name())
    if err != nil {
        exitErrorf("Unable to open file %q, %v", err)
    }

    defer processedImage.Close()

    sessUp, err := session.NewSession(&aws.Config{
        Region: aws.String("us-west-2")},
    )

    uploader := s3manager.NewUploader(sessUp)

    // Upload the file's body to S3 bucket as an object with the key being the
    // same as the filename.
    _, err = uploader.Upload(&s3manager.UploadInput{
        Bucket: aws.String(imageDestinationPath),
        Key: aws.String(originalImage),
        Body: processedImage,
    })
    if err != nil {
        // Print the error and exit.
        exitErrorf("Unable to upload %q to %q, %v", originalImage, imageDestinationPath, err)
    }

    fmt.Printf("Successfully uploaded %q to %q\n", originalImage, imageDestinationPath)


// after upload clean out tmp folder
	// delete anything in tmp folder
	

}