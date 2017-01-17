package main

import "fmt"

func main() {
	var jobList []Job
	jobs := make(chan Job)
	done := make(chan bool, len(jobList))

	go func() {
		for _, job := range jobList {
			jobs <- job // Заблокируется, пока принимающая строна не прочитает задание
		}
		close(jobs)
	}()

	go func() {
		for job := range jobs { //Заблокируется в ожидании передачи
			fmt.Println(job) // Выполнение задачи
			done <- true
		}
	}()

	for i := 0; i < len(jobList); i++ {
		<-done // заблокируется в ожидании передачи
	}
}
