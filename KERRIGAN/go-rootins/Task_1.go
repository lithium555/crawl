package main

import (
	"fmt"
	"time"
)

func main() {
	result := make(chan int)
	done := make(chan bool, 20)
	//Генерация значений
	go func() {
		for x := 0; x < 20; x++ {
			result <- x
		}
	}()
	//time.Sleep(time.Second)
	time.Sleep(time.Second)
	for {
		//fmt.Println(<-result)
		a := <-result
		if a == 10 {
			fmt.Println(a)
		}
	}
}

//
//func Add(value int) (int){
//	value++
//	return value
//}
