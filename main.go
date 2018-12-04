package main

import (
	"channel_test/dispatcher"
	"channel_test/job"
	"fmt"
	"time"
)

//type Job struct{
//	Title string
//}

func main() {
	job.JobQueue = make(chan job.Job,1000)

	dispatcher := dispatcher.NewDispatcher(1000)
	//fmt.Println(fmt.Sprintf("dispatcher object %v",dispatcher))
	dispatcher.Run()

	start := time.Now()

	//for {
		for i := 0; i < 500000; i++ {
			//time.Sleep(time.Nanosecond * 1)
			j := job.Job{Body: fmt.Sprintf("{'index':%d}", i)}
			fmt.Println(j)
			job.JobQueue <- j
		}
	//	time.Sleep(time.Second * 30)
	//}

	used := time.Since(start)
	used /= 1e9
	fmt.Println(fmt.Sprintf("used time : %d", used))

	for {
		time.Sleep(time.Second * 2)
		fmt.Println("waiting.............")
	}

	//channels := make([]chan bool, 6)
	//fmt.Println(cap(channels))
	//for i:= range channels  {
	//	channels[i] = make(chan bool)
	//}
	//go func() {
	//	for {
	//		channels[rand.Intn(6)] <- true
	//	}
	//}()
	//
	//for j:=0; j<36 ; j++ {
	//	var x int
	//	select {
	//	case <- channels[0]:
	//		x=1
	//	case <- channels[1]:
	//		x=2
	//	case <- channels[2]:
	//		x=3
	//	case <- channels[3]:
	//		x=4
	//	case <- channels[4]:
	//		x=5
	//	case <- channels[5]:
	//		x=6
	//	}
	//	//fmt.Println(x)
	//	fmt.Printf(" channel index %d get value \r\n",x)
	//	time.Sleep(time.Millisecond * 500 )
	//}
	//fmt.Println("end")
	//return
	//
	////lang.Shadow()
	////return
	//
	//jobList := make([]Job,100)
	//for i:= 0; i < 100; i++{
	//	var title = fmt.Sprintf("job index: %d",i)
	//	jobList[i] = Job{Title:title}
	//}
	//fmt.Println(jobList)
	//
	//jobs := make(chan Job)
	//done := make(chan bool, len(jobList))
	//
	//go func(){
	//	for _,job := range jobList{
	//		fmt.Printf("insert into job channel ...%s \r\n",job.Title )
	//		jobs <- job //阻塞，等待接受
	//	}
	//}()
	//
	//go func(){
	//	for job := range jobs{
	//		fmt.Printf("consumering from jobs channel .... %s \r\n",job.Title)
	//		done <- true
	//	}
	//}()
	//
	//for i:= 0 ; i< len(jobList); i++{
	//	fmt.Printf("main goroutine is doing...%d \r\n",i)
	//	<- done //阻塞，等待接受
	//}
}
