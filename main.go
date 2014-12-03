package main

import (
	"fmt"
	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/sqs"
	"golang-aws-sqs-example/worker"
	"runtime"
	"time"
)

const (
	accessKey = "A******************A"
	secretKey = "/dF/R**********************************W"
	queueName = "https://sqs.us-west-2.amazonaws.com/1**********7/job_queue"
)

func Print(msg *sqs.Message) error {
	//log.Println(msg.Body)
	fmt.Println(fmt.Sprintf("Message ID : %v done!  Current goroutine num : %d  [%s]", msg.MessageId, runtime.NumGoroutine(), time.Now().Local()))
	return nil
}

func main() {
	sleepTime := time.Millisecond * 200 // 0.2 second
	receiveMessageNum := 10
	fmt.Println("===========================================")
	fmt.Println(fmt.Sprintf(" Use CPU(s) num      : %d", runtime.NumCPU()))
	fmt.Println(fmt.Sprintf(" Receive message num : %d", receiveMessageNum))
	fmt.Println(fmt.Sprintf(" Sleep time          : 0.2 second(s)"))
	fmt.Println("===========================================")

	runtime.GOMAXPROCS(runtime.NumCPU())
	auth := aws.Auth{AccessKey: accessKey, SecretKey: secretKey}
	mySqs := sqs.New(auth, aws.USWest2)
	queue := &sqs.Queue{mySqs, queueName}
	worker.Start(queue, worker.HandlerFunc(Print), sleepTime, receiveMessageNum)
}
