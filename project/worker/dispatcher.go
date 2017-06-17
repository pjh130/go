package main

import (
	"log"
	_ "log"
	"sync"
	"time"
)

func StartDispatcher() {
	dispatcher := NewDispatcher(MAX_WORKERS)
	dispatcher.Run()
}

type Dispatcher struct {
	// A pool of workers channels that are registered with the dispatcher
	Lock       sync.Mutex
	WorkerPool map[int]*Worker
	MaxWorkers int
	NowWorkers int
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	return &Dispatcher{WorkerPool: make(map[int]*Worker),
		MaxWorkers: maxWorkers,
		NowWorkers: 0}
}

func (d *Dispatcher) Run() {
	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	//定期扫描闲置的worker
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			d.Lock.Lock()
			//todo...
			for key, w := range d.WorkerPool {
				if w.Working == false {
					log.Println("Delete a idle worker id:", key)
					w.Stop()
					delete(d.WorkerPool, key)
					d.NowWorkers--
				} else {
					//					log.Println("Delete working")
				}
			}
			d.Lock.Unlock()
		case job := <-JobQueue:
			log.Println("Get a job from queue id: ", job.Id)
			// a job request has been received
			//			go func(j Job) {
			//查找工作池
			find := false
			now := time.Now().Unix()
			for find == false {
				d.Lock.Lock()
				for key, w := range d.WorkerPool {
					if w.Working == false {
						log.Println("Find a idle worker id: ", key)

						w.SetWorking()
						d.WorkerPool[key] = w

						w.JobChannel <- job
						find = true
						break
					}
				}
				d.Lock.Unlock()

				//如果从工作池中取不到空闲的数据
				if find == false {
					d.Lock.Lock()
					if d.NowWorkers < d.MaxWorkers {
						//创建工作

						worker := NewWorker()
						worker.Start()

						log.Println("Create a new worker id: ", d.NowWorkers)
						d.NowWorkers++

						worker.SetWorking()
						d.WorkerPool[d.NowWorkers] = &worker

						worker.JobChannel <- job
						find = true
					} else {
						//如果没找到就继续循环等待
						//							log.Println("NowWorkers:", d.NowWorkers, " MaxWorkers: ", d.MaxWorkers)
					}
					d.Lock.Unlock()
				}
			}
			// dispatch the job to the worker job channel
			log.Println("time:", time.Now().Unix()-now)
			//			}(job)
		}
	}
}
