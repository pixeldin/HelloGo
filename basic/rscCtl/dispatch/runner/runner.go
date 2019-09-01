package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

//使用通道监视程序, 作为程序调度的监控 [来自Go语言实战]
type Runner struct {
	//来自OS 的signal
	interrupt chan os.Signal

	//
	complete chan error

	//
	timeout <-chan time.Time

	//tasks 函数切片列表
	tasks []func(int)
}

var (
	ErrTimeout   = errors.New("received timeout")
	ErrInterrupt = errors.New("received interrupt")
)

//初始化时候就执行倒计时
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		//计时器, timeout after duration d
		timeout:   time.After(d),
	}
}

func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		//Receive interrupt
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

func (r *Runner) Start() error {
	//Listen interrupt signal for gotInterrupt()
	signal.Notify(r.interrupt, os.Interrupt)
	go func() {
		r.complete <- r.run()
	}()
	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		//some case pass timeout decision
		return ErrTimeout
	}
}
