package main

import (
	"fmt"
	"time"
)

func main() {
	A := make(chan int)
	B := make(chan int)

	fmt.Println("create chans")

	go green(A)
	go blue(B)

	recv("A", A)
	recv("B", B)
	recv("A", A)

	sleep(time.Second * 3)

	recv("B", B)
}

func sleep(dt time.Duration) { // на входе перемнная, отвечающая за ограницение во времени
	fmt.Println("sleep for", dt)
	time.Sleep(dt) //Функция Sleep останавливает текущую рутину на dt времени.  A negative or zero duration causes Sleep to return immediately.
	fmt.Println("sleep done")
}

func send(name string, ch chan int, v int) { // отправка данных в канал
	fmt.Println(name, "pre", v)
	ch <- v // отправляю перемнную v в канал ch
	fmt.Println(name, "post", v)
}

func recv(name string, ch chan int) {
	fmt.Println("<-", name)
	v := <-ch // извлечение из канала данных в новую переменную v, ЗДЕСЬ И СОЗДАЛ, это н епеременная со строки 33!!
	fmt.Println("<-", name, "=", v)
}

func blue(ch chan int) { //функция отвечающая за синий поток, аргумент значение канала
	send("blue", ch, 200)
	send("blue", ch, 201)
}

func green(ch chan int) {
	send("green", ch, 100)
	send("green", ch, 101)
}
