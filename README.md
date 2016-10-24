# SQS

這裡的 Credential 用 Shared Credentials Provider 的方式

(你也可以直接用最簡單的 Access Key 及 Secret Key 來建立 AWS Session 連接 SQS)

> 要記得把 SQS 的權限給你的 User


# Install

使用 glide 套件管理安裝 :

    glide install

> Install glide : `go get github.com/Masterminds/glide`


# Modify config (in main.go)

    QueueURL    = "https://sqs.ap-northeast-1.amazonaws.com/3**********2/my-queue"
    Region      = "ap-northeast-1"
    CredPath    = "/Users/home/.aws/credentials"
    CredProfile = "aws-cred-profile"


# Run

Show how to `Add`, `Get` and `Del` a message.

    go run main.go

Output :

    [Send message]
    {
      MD5OfMessageBody: "d29343907090dff4cec4a9a0efb80d20",
      MessageId: "021d389c-156e-405a-b672-79b2c8e6af39"
    }

    [Receive message]
    {
      Messages: [{
          Body: "message body",
          MD5OfBody: "d29343907090dff4cec4a9a0efb80d20",
          MessageId: "021d389c-156e-405a-b672-79b2c8e6af39",
          ReceiptHandle: "AQEBdFvUW17FUHIG+iirIstj8O4F2xQ7GS+YjrDnyKl0sMTCpVRErSn0YwujbKbuqgnF6Q4wj8JXUTZ+p28h52WzVo/McV5DCwu68KaUSphPIB3OX1CHHjJX2aLInpGoxbjL84tNRiwPt/LACRGRhvslO2SZvGWeUkQ9YojXR+pmJGmYUCsrCKTHSBxPemS4ksG6VCb3FWt2hev8OWiYGZ0hTX9e5qSMfq1EjEHLRbKHK8aU+4X+RZm+2pd0cM8igsqVOQ6YEMrVr1LOtvWrAvo4kUat8hktQGvV6ZTtQRsBdIYqZz51fQZC3bisD+qvZlo7/ncwTNAHwYz2oTHSNgM1AZX9NUTCYG2bOS8+LzMIkC63sruS4W4m5N6zIVaq5yrr87JlXXp+dZqw9LkN21RJyA=="
        }]
    }

    [Delete message]
    Message ID: 021d389c-156e-405a-b672-79b2c8e6af39 has beed deleted.


# 備註

### Receive message output

    (*sqs.ReceiveMessageOutput)(0xc42043ac40)({
      Messages: [{
          Body: "{\"order\":1}",
          MD5OfBody: "190ad5fd0596a0629a0aa256937135a3",
          MessageId: "21768609-a37d-452d-92cf-6a867855d3a1",
          ReceiptHandle: "AQEBieaPYXFxxwrNc8WhYXZb0lhbW3To1iINXgoID8x3XvsJVGQNvtcmV/snP4NnLLRkKGQTX8m3kMxxbnSo91QHt8wrauwpKCFgEyJ+P/COQ/qzSqy1ZWAeWhm4Po7xpHA5ssC5GLBIZoWwEGVPZvKkXg7N79zoMi93ceeNuMxHQuhlATmm4PqC60E44/q2o2jTeZyr4cXsBFbwWNlEFXA/KCxHTg3z96W1hg5d3SZOQZqD+MH0jnO6d1ImXhHVyzBCuvlbaBTsUCiKPjNoFjOqVl/LKom4W6linhKw1fFp7hZ1RMG2oOQ+cJxkVuQxepBabJs6EbTHNhmNCEXV1palEMdRTHZPHvyA0ftRgRQTM1kLIOckMDd7kxVY5ZNfGaPs+LL+H9049ZITTyHR5MkSXg=="
        }]
    })

> ReceiptHandle 是用來提供刪除 message 的參數而不是 message id，每一次取出來的值都會變，刪除時要提供最後取出來的 ReceiptHandle 值。


### 其他

1. 每一個 request 預設上限大小是 64K，最多可以到 256K
2. 一次收到的 message 預設是 1 筆，但你可以自已調整，最多 10 筆
3. 在 polling message 時，為了避免不停地取空的 message，建議使用 WaitTimeSeconds 參數(可以自已定義等待時間，最多 20 秒)
4. 當你取出來 message 後，這個 message 會暫時消失一段時間 (visibility timeout)，這段時間內不會再被取到，這個時間預設是 30 秒左右，你也可以自已設定。建議設定的時間要比你的 worker 處理這個 message 的時間長，避免在還沒處理完又再次取到。需要考量 network delay / slow machines / too much IO / 刪除 message 的時間 / etc. 再設定要給多長的時間
5. 當取出一個 message 且處理完後要記得 delete，如果沒有 delete，過了 visibility timeout 後，它還是會再被取出來
