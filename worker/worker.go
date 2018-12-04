package worker

import (
	"channel_test/job"
	"fmt"
)

type Worker struct {
	JobChannel chan job.Job
	quit       chan bool
}

func NewWorker() *Worker {
	return &Worker{JobChannel: make(chan job.Job), quit: make(chan bool)}
}

func (w *Worker) Start() {
	go func() {
		//处理jobchannel里的job
		for  {
			select {
			case job,ok := <- w.JobChannel:
				if ok {
					//fmt.Println(fmt.Sprintf("worker jobchannel有值了 %v",job))
					job.Do()
				}else{
					fmt.Println("worker's jobchannel has been closed!!!")
				}


			case <-w.quit:
				return
			}
		}


	}()
}

func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
