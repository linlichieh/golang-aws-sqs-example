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
	QueueUrl    = "https://sqs.ap-northeast-1.amazonaws.com/3**********2/survey-worker"
	Region      = "ap-northeast-1"
	CredPath    = "/Users/jex/.aws/credentials"
	CredProfile = "survey-worker"
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
		MessageBody: aws.String("message body"), // Required
		QueueUrl:    aws.String(QueueUrl),       // Required
	}
	send_resp, err := svc.SendMessage(send_params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Send message :\n\n%v \n\n", send_resp)

	// Receive message
	receive_params := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(QueueUrl),
		MaxNumberOfMessages: aws.Int64(3),
		VisibilityTimeout:   aws.Int64(1),
		WaitTimeSeconds:     aws.Int64(20),
	}

	receive_resp, err := svc.ReceiveMessage(receive_params)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Receive message :\n\n%v \n\n", receive_resp)

	// Delete message
	for _, message := range receive_resp.Messages {
		delete_params := &sqs.DeleteMessageInput{
			QueueUrl:      aws.String(QueueUrl),  // Required
			ReceiptHandle: message.ReceiptHandle, // Required

		}

		_, err := svc.DeleteMessage(delete_params) // 成功刪除不會返回錯誤訊息
		if err != nil {
			log.Println(err)
		}
		fmt.Println("Successfully delete message")
	}
}
