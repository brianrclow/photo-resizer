# photo-resizer

photo-resizer is an application written in Go for converting large jpg images of any resolution to the size of the iPhone 11 Pro screen resolution. It resizes the images proportionally to a width of 2436 pixels.

This project allows for the application to run locally or on AWS Lambda using Amazon S3 for storage of the images. This has been adapted from: https://github.com/nfnt/resize to allow for scalable adornment of photos.

## Usage

### Local
How to run:

1. Navigate to src/local/
2. Create 2 directories named "input" and "output"
3. Add jpg images you want to resize to "input" directory
4. Run "go run main.go"
4. View the resized images in src/local/output/

### AWS Lambda
How to run:

This is being worked on now...


## License

[MIT](LICENSE)
