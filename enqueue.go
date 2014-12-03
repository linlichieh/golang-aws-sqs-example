package main

import (
	"fmt"
	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/sqs"
	"time"
)

const (
	accessKey = "A******************A"
	secretKey = "/dF/R**********************************W"
	queueName = "https://sqs.us-west-2.amazonaws.com/1**********7/job_queue"
)

func main() {

	auth := aws.Auth{AccessKey: accessKey, SecretKey: secretKey}
	mySqs := sqs.New(auth, aws.USWest2)
	q := &sqs.Queue{mySqs, queueName}

	//send message
	resp, _ := q.SendMessage(fmt.Sprintf("This is a message from testing. (%v)", time.Now().Format(time.RFC850)))
	fmt.Println(resp)
}
