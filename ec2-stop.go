package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
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

	// Create an EC2 service client
	ec2Client := ec2.New(sess)

	// Specify the instance ID you want to stop
	instanceID := "i-02532410da953a567" // Replace with your instance ID

	// Stop the EC2 instance
	_, err = ec2Client.StopInstances(&ec2.StopInstancesInput{
		InstanceIds: []*string{aws.String(instanceID)},
	})
	if err != nil {
		fmt.Println("Failed to stop EC2 instance:", err)
		return
	}

	fmt.Println("EC2 Instance stopped:", instanceID)
}

//github.com/aws/aws-sdk-go/service/ec2
