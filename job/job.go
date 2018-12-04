package job

import (
	"fmt"
)

type Job struct {
	Body string `json:"body"`
}

var JobQueue chan Job

func (j Job) Do() {
	//如果不sleep，就是模拟CPU密集型，CPU利用率会飚满。这种场景对pool没要求，一个channel就能跑。
	//如果sleep，就是模拟IO密集型，那么，pool的size越大越好，因为pool越大，同时并行执行的数量越大，但是相应的内存占用就越大。
	//time.Sleep(time.Second * 1)
	fmt.Println(fmt.Sprintf("im doing job's body %s", j.Body))
}
