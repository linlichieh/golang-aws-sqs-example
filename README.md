# SQS

這裡的 Credential 用 Shared Credentials Provider 的方式

(你也可以直接用最簡單的 Access Key 及 Secret Key 來建立 AWS Session 連接 SQS)

> 要記得把 SQS 的權限給你的 User

# Install

使用 glide 套件管理安裝 :

    glide install

# Config (in main.go)

    QueueURL    = "https://sqs.ap-northeast-1.amazonaws.com/3**********2/my-queue"
    Region      = "ap-northeast-1"
    CredPath    = "/Users/home/.aws/credentials"
    CredProfile = "aws-cred-profile"

# Run

    // Show how to Add, Get and Del a message.
    go run simple_example.go

Output :

    Send message :

    {
      MD5OfMessageBody: "d29343907090dff4cec4a9a0efb80d20",
      MessageId: "e4f8a1c5-3401-4388-82c7-e130dc01266d"
    }

    Receive message :

    {
      Messages: [{
          Body: "message body",
          MD5OfBody: "d29343907090dff4cec4a9a0efb80d20",
          MessageId: "e4f8a1c5-3401-4388-82c7-e130dc01266d",
          ReceiptHandle: "AQEBXFv4TCjC6RUaYmVWono55zy+46VMqUH4Jp1k1jY5f8i1smT02A437JeyHU7XTZWFoRjIFlDukVpb4Dzxdwn8dkHqmn+vTCfq8YLB43g5AWVFdFgCprXS2yxM11wm4NrYZvvUhqgIq3wH6CPUKzAzQDFGjYmYho2hmYBohmjT4HsgvOGQbMPC5js0XaQKM71dK31A3uF/6UFnyDPgwr74VRIUHuCuKcD1PwdvcDtG/HaCVAYjDbkxXRgnnU7fhaHMDP+hTd1y0+VI5Fwyn9bxGmCSyVoxwceBXzuIItjZAPFQjIRKoRPxI28NXvBOKS9hUSIEToDq6feE3wsYfDvztuQUEnsyG8jpes2i+rrzZ18MRYJRJbjaFZrsicS3skIoHDuZ1XyshIt8IULOiZwLmg=="
        }]
    }

    Successfully delete message


