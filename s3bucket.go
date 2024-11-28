// s3-bucket.go
package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	// Create a new session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		fmt.Println("Failed to create session:", err)
		return
	}

	// Create an S3 service client
	s3Client := s3.New(sess)

	// Specify the bucket name
	bucketName := "mygobucket07bd" // Replace with your bucket name

	// Create the S3 bucket
	_, err = s3Client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		fmt.Println("Failed to create S3 bucket:", err)
		return
	}

	fmt.Println("S3 Bucket created successfully:", bucketName)
}

// need to configure below things

//go mod init s3bucket

//go get github.com/aws/aws-sdk-go/aws
//go get github.com/aws/aws-sdk-go/aws/session
//go get github.com/aws/aws-sdk-go/service/s3

//aws configure
