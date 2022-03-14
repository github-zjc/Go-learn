package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

//给定一系列task，在timeout内完成，否则报错
//如果操作系统给了终端信号也报错
var (
	ErrTimeout   = errors.New("cannot finish tasks within the timeout")
	ErrInterrupt = errors.New("received interrupt from OS")
)

type Runner struct {
	interrupt chan os.Signal
	complete  chan error
	timeout   <-chan time.Time //用来计时
	tasks     []func(int)      //任务列表
}

func New(t time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(t), //当t秒过后，time.After这个函数将会把t这个时间写入到timeout这个<-chan
		tasks:     make([]func(int), 0),
	}
}

func (r *Runner) AddTasks(task ...func(int)) {
	r.tasks = append(r.tasks, task...)
}

func (r *Runner) Run() error {
	for id, task := range r.tasks {
		select {
		case <-r.interrupt:
			signal.Stop(r.interrupt)
			return ErrInterrupt
		case <-r.timeout:
			return ErrTimeout
		default:
			task(id)
		}
	}
	return nil
}

func (r *Runner) Start() error {
	//等待接收操作系统的中断信号
	signal.Notify(r.interrupt, os.Interrupt)

	//运行tasks
	go func() {
		r.complete <- r.Run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
		// default:   //不可以加这一条，因为r.Run是以协程启动的，返回的err还没来得及r.Start的select就选到了default退出
		// 	return nil
	}

}
