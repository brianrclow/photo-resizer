# photo-resizer
This go program will resize jpg images proportionally to a width of 2436pixels. This project was created to test out the decode and encode process of Go along with using AWS Lambda to see how this type of project could be scaled. This has been adapted from: https://github.com/nfnt/resize to allow for scalable adornment of photos.

How to run:

1. Drop all photos that need to be resized into the input folder
2. Make sure that the filepath is correct in the program for input and output
3. Run "go run main.go"
4. View the processed photos from the output folder
