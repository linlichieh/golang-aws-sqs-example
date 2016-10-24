package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

const (
	QueueUrl    = "https://sqs.ap-northeast-1.amazonaws.com/3**********2/my-queue"
	Region      = "ap-northeast-1"
	CredPath    = "/Users/home/.aws/credentials"
	CredProfile = "aws-cred-profile"
)

func main() {

	sess := session.New(&aws.Config{
		Region:      aws.String(Region),
		Credentials: credentials.NewSharedCredentials(CredPath, CredProfile),
		MaxRetries:  aws.Int(5),
	})

	svc := sqs.New(sess)

	// Send message
	send_params := &sqs.SendMessageInput{
		MessageBody:  aws.String("message body"), // Required
		QueueUrl:     aws.String(QueueUrl),       // Required
		DelaySeconds: aws.Int64(3),               // (optional) 傳進去的 message 延遲 n 秒才會被取出, 0 ~ 900s (15 minutes)
	}
	send_resp, err := svc.SendMessage(send_params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("[Send message] \n%v \n\n", send_resp)

	// Receive message
	receive_params := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(QueueUrl),
		MaxNumberOfMessages: aws.Int64(3),  // 一次最多取幾個 message
		VisibilityTimeout:   aws.Int64(30), // 如果這個 message 沒刪除，下次再被取出來的時間
		WaitTimeSeconds:     aws.Int64(20), // long polling 方式取，會建立一條長連線並且等在那邊，直到 SQS 收到新 message 回傳給這條連線才中斷
	}
	receive_resp, err := svc.ReceiveMessage(receive_params)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("[Receive message] \n%v \n\n", receive_resp)

	// Delete message
	for _, message := range receive_resp.Messages {
		delete_params := &sqs.DeleteMessageInput{
			QueueUrl:      aws.String(QueueUrl),  // Required
			ReceiptHandle: message.ReceiptHandle, // Required

		}
		_, err := svc.DeleteMessage(delete_params) // No response returned when successed.
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("[Delete message] \nMessage ID: %s has beed deleted.\n\n", *message.MessageId)
	}
}
