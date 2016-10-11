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

# 備註

1. 每一個 request 預設上限大小是 64K，最多可以到 256K
2. 一次收到的 message 預設是 1 筆，但你可以自已設定，最多 10 筆
3. 在 polling message 時使用 WaitTimeSeconds (可以自定義等待時間，最多 20 秒，)，為了避免忙錄的取空的 message，
4. 當你取出來 message 後，這個 message 會暫時消失一段時間 (visibility timeout)，這個時間預設是 30 秒左右，你也可以自已設定，建議設定的時間要比你的 worker 處理這個 message 的時間長，避免在還沒處理完又再次取到。需要考量 network delay / slow machines / too much IO / 刪除 message 的時間 / etc.
5. 當取出一個 message 且處理完後要記得 delete，如果沒有 delete，過了 visibility timeout 後，它還是會再被取出來一次
