package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handler start")
	ctx := r.Context()

	complete := make(chan struct{})
	go func() {
		time.Sleep(5 * time.Second)
		complete <- struct{}{}
	}()

	select {
	case <-ctx.Done(): //ctx is cancel
		err := ctx.Err()
		if err != nil {
			fmt.Println(err.Error())
		}
		// case <-time.After(3 * time.Second): //3秒一到，time.After就会返回一个<-Time
		// 	fmt.Println("finish something")
	case <-complete:
		fmt.Println("finish something")
	}
	fmt.Println("hander ends")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
