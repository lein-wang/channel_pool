package dispatcher

import (
	_ "channel_test/worker"
	worker2 "channel_test/worker"
)
import (
	"channel_test/job"
	"fmt"
	"math/rand"
)

type Dispatcher struct {
	WorkPool   []chan job.Job
	MaxWorkers int
}

//新建分发器：指定最大worker数，即pool的size，pool里放的是worker的jobchannel
//分发器的作用就是把job丢到pool里，也就是worker的jobchannel里
func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make([]chan job.Job,0)
	fmt.Println(fmt.Sprintf("pool obj %v", pool))
	return &Dispatcher{WorkPool: pool, MaxWorkers: maxWorkers}
}

/**
分发器启动的时候就去创建worker，把worker的jobchannel组成一个pool
*/
func (d *Dispatcher) Run() {
	//开启指定个数 worker
	for i := 0; i < d.MaxWorkers; i++ {
		worker := worker2.NewWorker()
		d.WorkPool = append(d.WorkPool,worker.JobChannel)
		worker.Start()
	}
		fmt.Println(fmt.Sprintf("d %v", d))
	//监听jobQueue的状态，做分发
	go d.Dispatch()
}

/**
pool建好了，自然要往里头放东西咯
*/
func (d *Dispatcher) Dispatch() {
	for {
		select {
		case j := <-job.JobQueue:
			//fmt.Println(fmt.Sprintf("job channel has value : %v",j))
			//job的管道里有东西了
			go func(job job.Job) {
				//fmt.Println(fmt.Sprintf("job 11111obj : %v", job))
				//fmt.Println(fmt.Sprintf("d.workpool why empty : %v", d.WorkPool))

				//获取pool
				jobChannel := d.WorkPool[rand.Intn(len(d.WorkPool))]
				//fmt.Println(fmt.Sprintf("get d.WorkPoll --> jobchannel %v", jobChannel))
				//把job丢到pool里
				select {
				case jobChannel <- job:
					//fmt.Println(fmt.Sprintf("job insert into jobchannel, jobchannel is %v", jobChannel))

				}

			}(j)
		}
	}
}
