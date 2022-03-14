package main

import (
	"context"
	"fmt"
	"time"
)

func DoSomething(ctx context.Context) {
	select {
	case <-ctx.Done(): //ctx is cancel
		err := ctx.Err()
		if err != nil {
			fmt.Println(err.Error())
		}
	case <-time.After(3 * time.Second): //3秒一到，time.After就会返回一个<-Time
		fmt.Println("finish something")
	}

}

func main() {
	//创建空context的两种方法
	ctx := context.Background() //返回一个空的context，不能被cancel，kv为空，根context

	//todoctx := context.TODO()	 //和Background类似，当你不确定的适合使用

	ctx, cancel := context.WithCancel(ctx) //由根context生成的子context

	go func() {
		time.Sleep(time.Second)
		cancel()
	}()

	DoSomething(ctx)
}
