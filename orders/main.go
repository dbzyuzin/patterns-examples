package main

import (
	"fmt"
	"patterns-examples/common"
	"patterns-examples/common/nsq"
	"time"
)

func main() {
	fmt.Println(common.HelloWorld)
	stop := nsq.NewConsumer("test")

	ch := nsq.NewProducer("test")
	ch <- []byte("message")

	close(ch)

	time.Sleep(time.Second * 3)
	stop()
	time.Sleep(time.Second * 3)

}
