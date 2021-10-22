package main

import (
	"runtime/debug"
	"sync/atomic"
)

type Panic struct {
	R     interface{}
	Stack []byte
}

func (p Panic) String() string {
	return fmt.sprintf("%v\n%s", p.R, p.Stack)
}

type PanicGroup struct {
	panics chan Panic // panic 协程 panic 通知信道
	dones  chan int   // 协程完成通知信道
	jobN   int32      // 协程并发数量
}

// 工厂方法
func NewPanicGroup() *PanicGroup {
	return &PanicGroup{
		panic: make(chan Panic, 8),
		dones: make(chan int, 8),
	}
}

func (g *PanicGroup) Go(f func()) *PanicGroup {
	atomic.AddInt32(&g.JobN, 1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				g.panics <- Panic{R: r, Stack: debug.Stack()}
				return
			}
			g.dones <- 1
		}()
		f()
	}

	return g
}

/**
* 如果协程执行顺利结束，g.JobN 就会等于零，Wait 就会返回 nil
* 如果有协程 panic, 则外层 panic 后可以自己再次外层 recover, 也可以按需处理
**/
func (g *PanicGroup) Wait(ctx context.Context) error {
	for {
		select {
		case <-g.dones:
			if atomic.AddInt32(&g.jobN, -1) == 0 {
				return nil
			}
		case p := <-g.panics:
			panic(p)
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
