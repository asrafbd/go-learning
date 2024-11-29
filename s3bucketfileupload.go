package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadFile(uploader *s3manager.Uploader, filePath string, bucketName string, fileName string) error {

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   file,
	})

	return err
}

func main() {

	// Create a new session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		fmt.Println("Failed to create session:", err)
		return
	}

	// Specify the bucket name
	bucketName := "mygobucket07bd" // Replace with your bucket name
	uploader := s3manager.NewUploader(sess)
	filepath := "./test.txt"
	filename := "test.txt"

	err = UploadFile(uploader, filepath, bucketName, filename)

	if err != nil {

		fmt.Printf("Could not upload file: %v", err)
		return
	}

	fmt.Println("successfully uploaded:", filename)
}

// need to configure below things

//go mod init s3bucket

//go get github.com/aws/aws-sdk-go/aws
//go get github.com/aws/aws-sdk-go/aws/session
//go get github.com/aws/aws-sdk-go/service/s3

//aws configure
