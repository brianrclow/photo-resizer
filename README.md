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

MIT License

Copyright (c) 2020 Brian R. Clow

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.