package workerpool

import (
	"log"
	"sync"
)

type task struct {
	args []interface{}
	fn   interface{}
}

type routinePool struct {
	max       int //go程上限
	count     int //当前go程数量
	freeCount int //空闲数量
	taskCount int //待处理任务数量
	taskQue   chan *task
	mtx       sync.Mutex
	closed    bool
	closeWait sync.WaitGroup
}

const BlockClose bool = true

/*
*   创建一个go程池，池子内go程上限为max
*   (注意，将创建一个大小为max的chan作为任务队列，调用AddTask时chan满的话将导致AddTask被阻塞,
*   此时task数量至少有max:正在处理的任务 + max:chan中待处理的任务 2*max个，即当系统中待处理及尚未处理完成的
*   任务数量超过2*max，AddTask将阻塞)
*
 */

func NewRoutinePool(max int) *routinePool {
	if max <= 0 {
		max = 1024
	}

	r := &routinePool{
		max:     max,
		count:   0,
		taskQue: make(chan *task, max),
	}

	return r
}

func pcall(t *task) {

	if _, err := ProtectCall(t.fn, t.args); err != nil {

		log.Println(err)
	}
}

func (this *routinePool) newRoutine() {
	//创建一个新的go程
	go func() {
		this.mtx.Lock()
		defer this.mtx.Unlock()
		this.freeCount++

		for t := range this.taskQue {
			//准备执行任务，空闲数量减1
			this.taskCount--
			this.freeCount--

			pcall(t)

			//执行完毕，空闲数量加1
			this.freeCount++

			//一次只取一个任务
			break
		}

		this.count--
		this.freeCount--

		this.closeWait.Done()
	}()
}

func (this *routinePool) AddTask(fn interface{}, args ...interface{}) bool {

	t := &task{fn: fn, args: args}
	this.mtx.Lock()
	defer this.mtx.Unlock()

	if this.closed {
		return false
	} else {
		this.taskCount++
		if this.freeCount < this.taskCount && this.count < this.max {
			this.count++
			this.closeWait.Add(1)
			//没有空闲go程，且go程数量尚未达到上限
			this.newRoutine()
		}

		this.taskQue <- t
		return true
	}
}

func (this *routinePool) Close(block ...bool) {
	this.mtx.Lock()
	defer func() {
		this.mtx.Unlock()
		if len(block) > 0 && block[0] {
			this.closeWait.Wait()
		}

	}()
	if !this.closed {
		this.closed = true
		close(this.taskQue)
	}
}
