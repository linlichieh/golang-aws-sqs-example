package main

import (
    "fmt"
    "github.com/crowdmob/goamz/aws"
    "github.com/crowdmob/goamz/sqs"
    "time"
)

const (
    accessKey = "A******************Q"
    secretKey = "Q**************************************T"
    queueName = "https://sqs.ap-northeast-1.amazonaws.com/5**********5/TestQueue"
)

func main() {

    auth := aws.Auth{AccessKey: accessKey, SecretKey: secretKey}
    mySqs := sqs.New(auth, aws.APNortheast)
    q := &sqs.Queue{mySqs, queueName}

    //send message
    resp, _ := q.SendMessage(fmt.Sprintf("This is a test message from linode. (%v)", time.Now().Format(time.RFC850)))
    fmt.Println(resp)
}
