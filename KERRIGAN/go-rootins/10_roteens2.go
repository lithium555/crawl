package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()

	chanel := make(chan int)
	chanelERR := make(chan error)
	for i := 0; i < 10; i++ {
		go master(i, chanelERR, chanel)
	}
	var slice []int
	var sliErr []error
	for j := 0; j < 10; j++ {
		select {
		case here := <-chanel:
			slice = append(slice, here)
			fmt.Printf("HERE IS: '%v'\n", here)
		case receive_err := <-chanelERR:
			sliErr = append(sliErr, receive_err)
		}
	}
	log.Printf("SLICE: '%v'\n", slice)
	log.Printf("SliERR: '%v'\n", sliErr)

	dt := time.Since(start)
	log.Printf("THE TIME IS: '%v'\n", dt)
}

/*
 Теперь попробуем такое. Есть функция, которая делает вид что выполняет
 тяжелую работу и возвращает результат либо ошибку:

 Нужно написать программу которая запустит 10 рутин, каждая из которых вызовет эту
 функцию передав свой номер. Выйдет что каждая рутина получит свой уникальный результат,
 либо оишбку. Мейн должен распечатать результаты от всех рутин. Порядок не важен.
*/
func calculate(i int) (int, error) {
	// wait a random amount of time to emulate some work
	time.Sleep(time.Second/4 + time.Duration(rand.Intn(int(time.Second*5))))
	rez := rand.Intn(10)
	if rez >= 7 {
		return 0, fmt.Errorf("error, number is too large: %v, routine %d", rez, i)
	}
	return rez, nil
}

func master(i int, chanelERR chan error, chanel chan int) {
	rezult, err := calculate(i)
	if err != nil {
		log.Printf("CONTINUE '%v'", i)
		chanelERR <- err
		return
	}
	chanel <- rezult
	fmt.Printf("Number of rooten: '%v'\n", i)
}
