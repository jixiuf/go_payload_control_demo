package main

import (
	"fmt"

	"time"

	"github.com/jixiuf/go_payload_control_demo/payloadcontrol"
)

type Handler struct {
	Name string
	Args string
}

func (this Handler) Handle() error {
	fmt.Printf("handle %s with args %s\n", this.Name, this.Args)
	time.Sleep(time.Second)
	return nil
}

func main() {
	// 同时有2个worker,
	dispatcher := payloadcontrol.NewDispatcher(2, 0)
	dispatcher.Run()

	for i := 0; i < 10; i++ {
		go func(idx int) {
			for {
				startTime := time.Now()
				// wonot block here
				dispatcher.Push(Handler{Name: fmt.Sprintf("hello %d", idx), Args: "http request args"})
				fmt.Println("push", time.Since(startTime)/time.Millisecond)
				// time.Sleep(time.Millisecond * 100)
			}
		}(i)
	}
	<-make(chan bool)

}
