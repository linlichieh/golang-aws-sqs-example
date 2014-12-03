package worker

import (
	"fmt"
	"github.com/crowdmob/goamz/sqs"
	"github.com/garyburd/redigo/redis"
	"log"
	"sync"
	"time"
)

type Handler interface {
	HandleMessage(msg *sqs.Message) error
}

type HandlerFunc func(msg *sqs.Message) error

func (f HandlerFunc) HandleMessage(msg *sqs.Message) error {
	return f(msg)
}

func Start(q *sqs.Queue, h Handler, t time.Duration, receiveMessageNum int) {
	fmt.Println(fmt.Sprintf("worker: Start polling  [%s]", time.Now().Local()))
	// init redis
	pool := redis.Pool{
		MaxIdle:     3,
		MaxActive:   0, // When zero, there is no limit on the number of connections in the pool.
		IdleTimeout: 30 * time.Second,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				log.Fatal(err.Error())
			}
			return conn, err
		},
	}
	for {
		resp, err := q.ReceiveMessage(receiveMessageNum)
		if err != nil {
			log.Println(err)
			continue
		}
		if len(resp.Messages) > 0 {
			fmt.Printf("\r")
			run(q, h, resp, pool)
		} else {
			fmt.Printf(".")
		}
		time.Sleep(t)
	}
}

// poll launches goroutine per received message and wait for all message to be processed
func run(q *sqs.Queue, h Handler, resp *sqs.ReceiveMessageResponse, redisPool redis.Pool) {
	var wg sync.WaitGroup
	for i := range resp.Messages {
		// Get a redis connection from pool
		redisConn := redisPool.Get()
		defer redisConn.Close()

		// Check job flag
		jobFlag, _ := redis.String(redisConn.Do("GET", "aws-sqs-flag-"+resp.Messages[i].MessageId))
		if jobFlag != "" && jobFlag == "1" {
			fmt.Println(fmt.Sprintf(" !!! Message ID : %s  Get the same job! Ignore... [%s]", resp.Messages[i].MessageId, time.Now().Local()))
			continue
		} else {
			wg.Add(1)
			redisConn.Do("SET", "aws-sqs-flag-"+resp.Messages[i].MessageId, "1")
		}

		go func(m *sqs.Message) {
			if err := handleMessage(q, m, h); err != nil {
				log.Fatalf(" *** Message ID : %s  handleMessage error : %s", m.MessageId, err.Error())
			}
			if _, err := redisConn.Do("DEL", "aws-sqs-flag-"+m.MessageId); err != nil {
				log.Fatalf(" *** Message ID : %s  Redis del error : %s", m.MessageId, err.Error())
			}

			wg.Done()
		}(&resp.Messages[i])
	}
	wg.Wait()
}

func handleMessage(q *sqs.Queue, m *sqs.Message, h Handler) error {
	var err error
	err = h.HandleMessage(m)
	if err != nil {
		return err
	}
	// delete
	_, err = q.DeleteMessage(m)
	return err
}
