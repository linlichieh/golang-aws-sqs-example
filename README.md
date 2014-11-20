# SQS 需要 IAM 管理的 Access Key 及 Secret Key

SQS 這兩把 key 是使用 IAM 管理, 所以在 IAM 那點擊 User 會看到 Access Key,

如果當初管理者沒有把 Secret 記下來, 那麼就要在該頁方點擊 `Manage Access Keys` 再點 `Create Access Key` 重新建立,

才會顯示 Secret Key, 注意! 記得在關掉 popup 前把 Secret Key 抄下來

**要記得把 SQS 的權限給你的 User, 點擊 `Attach User Policy`**

# Install

    go get github.com/crowdmob/goamz/sqs
    go get github.com/crowdmob/goamz/aws

# Type your access key, secret key, SQS queue name and region (enqueue.go and main.go)

    accessKey = "*************"
    secretKey = "*************"
    queueName = https://sqs.ap-northeast-1.amazonaws.com/5**********5/TestQueue

    mySqs := sqs.New(auth, aws.APNortheast)  <= Change to your current region in enqueue.go and main.go

> [region list](https://github.com/crowdmob/goamz/blob/master/aws/regions.go)

# Run enqueue

    mv enqueue.go /tmp/
    go run /tmp/enqueue.go

# Run worker

    go build
    ./golang-aws-sqs-example


# You will see (in worker) :

    2014/10/29 12:22:22 worker: Start polling
    2014/10/29 12:22:22 worker: Received 1 messages
    2014/10/29 12:22:22 worker: Spawned worker goroutine
    This is a test message from golang
    2014/10/29 12:22:22 worker: Start polling
    2014/10/29 12:22:22 worker: Start polling

# Others

* 用程式撈出來一次大約只能撈到 1~5 筆 (即使程式指定一次要撈 100 筆)
* 被撈過的 job 會被 AWS SQS 內部 flag 起來, 下一次撈就不會再撈到它, 除非 enqueue 一筆新的 job, 才可以再重新撈取 "已被撈取過的 job"
* 建議將撈出來的 job 直接做完及 delete job, 不要撈出來不處理而寄放在 AWS SQS 那邊, 否則如果沒有新 job enqueue 進來, 那筆怎麼撈都不會出現
* 如果使用兩個不同 process 的 worker, 要考慮可能會取到同一個 queue 的問題, 建議搭配 redis 做 flag

# 參考及修改來源 `https://github.com/nabeken/golang-sqs-worker-example`
