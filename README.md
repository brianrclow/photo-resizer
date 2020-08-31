# photo-resizer

photo-resizer is an application written in Go for converting large jpg images of any resolution to the size of the iPhone 11 Pro screen resolution. It resizes the images proportionally to a width of 2436 pixels.

This project allows for the application to run locally or on AWS Lambda using Amazon S3 for storage of the images. This has been adapted from: https://github.com/nfnt/resize to allow for scalable adornment of photos.


## Installation

## Usage

### Local
How to run:

1. Drop all photos that need to be resized into the input folder
2. Make sure that the filepath is correct in the program for input and output
3. Run "go run main.go"
4. View the processed photos from the output folder


### AWS Lambda



## License

MIT