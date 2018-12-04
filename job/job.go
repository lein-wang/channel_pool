package job

import (
	"fmt"
)

type Job struct {
	Body string `json:"body"`
}

var JobQueue chan Job

func (j Job) Do() {
	//time.Sleep(time.Second * 1)
	fmt.Println(fmt.Sprintf("im doing job's body %s", j.Body))
}
