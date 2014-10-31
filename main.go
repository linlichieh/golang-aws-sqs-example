package main

import (
    "fmt"
    "github.com/crowdmob/goamz/sqs"
    "github.com/crowdmob/goamz/aws"
    "golang-aws-sqs-example/worker"
)

const (
    accessKey = "A******************Q"
    secretKey = "Q**************************************T"
    queueName = "https://sqs.ap-northeast-1.amazonaws.com/5**********5/TestQueue"
)

func Print(msg *sqs.Message) error {
    fmt.Println(msg.Body)
    return nil
}

func main() {
    auth := aws.Auth{AccessKey: accessKey, SecretKey: secretKey}
    mySqs := sqs.New(auth, aws.APNortheast)
    queue := &sqs.Queue{mySqs, queueName};

    worker.Start(queue, worker.HandlerFunc(Print))
}
