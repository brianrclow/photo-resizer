# photo-resizer

photo-resizer is an application written in Go for converting large jpg images of any resolution to the size of the iPhone 11 Pro screen resolution. It resizes the images proportionally to a width of 2436 pixels.

This project allows for the application to run locally or on AWS Lambda using Amazon S3 for storage of the images. This has been adapted from: https://github.com/nfnt/resize to allow for scalable adornment of photos.

Test images provided free from unsplash.com

Current
* Local implementation allows for resizing of mulitple images in a folder
* AWS implementation only allows for resizing a single image


## Youtube Video

I talk about AWS Amplify and Cognito in this video as well as demo the app as a completely new user. You can view the video:
https://youtu.be/WlnMekNAZaw

## Usage

### Local
How to run:

1. Navigate to src/local/
2. Create 2 directories named "input" and "output"
3. Add jpg images you want to resize to "input" directory
4. Run "go run main.go"
4. View the resized images in src/local/output/

### AWS Lambda

Prerequisities:
* Install the AWS Go SDK along with AWS Lambda Go.
* Create a Lambda Function with a trigger to a specific bucket/directory.
* Add 3 environment variables to your lambda function
    * imageDestinationPath - the path you want the resized image to be saved to.
    * imageName - full name of the image including its suffix (.jpg)
    * imageSrcPath - the source bucket/directory of the image.
* Set your Lambda Handler to "main"
* Setup a Role that has access to lambda, cloudwatchlogs and S3

How to run:
1. Navigate to the src/aws/
2. Run "GOOS=linux go build main.go" to ensure the go executable is compatible with the go runtime.
3. Run "zip function.zip main" to create the zip file with the go executable
4. Upload function.zip to AWS Lambda


## File Structure
```
images
 └─── test images
src
 ├── aws
 │     └─── main.go (AWS implementation)
 └── local
       └─── main.go (local implementation)
```


## License

[MIT](LICENSE)
