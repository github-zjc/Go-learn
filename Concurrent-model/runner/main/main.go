package main

import (
	"fmt"
	"time"
	"runner"
)

func createTask() func(int) {
	return func(id int) {
		time.Sleep(time.Second)
		fmt.Printf("任务完成 #%d\n", id)
	}
}

func main() {
	r := runner.New(3 * time.Second)

	r.AddTasks(createTask(), createTask(), createTask())

	err := r.Start()
	switch err {
	case runner.ErrInterrupt:
		fmt.Println("tasks interrupt")
	case runner.ErrTimeout:
		fmt.Println("tasks timeout")
	default:
		fmt.Println("all tasks finished")
	}
}
