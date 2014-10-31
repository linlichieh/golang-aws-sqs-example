package main

import (
    "fmt"
    "github.com/crowdmob/goamz/sqs"
    "github.com/crowdmob/goamz/aws"
)

const (
    accessKey = "A******************Q"
    secretKey = "Q**************************************T"
    queueName = "https://sqs.ap-northeast-1.amazonaws.com/5**********5/TestQueue"
)

func main() {

    auth := aws.Auth{AccessKey: accessKey, SecretKey: secretKey}
    mySqs := sqs.New(auth, aws.APNortheast) // or aws.APNortheast
    q := &sqs.Queue{mySqs, queueName}

    //send message
    resp, _ := q.SendMessage("This is a test message from golang")
    fmt.Println(resp)
}

