package main

import (
	"fmt"
	"time"
)

/*
Каналы обеспечивают возможность общения нескольких горутин друг с другом,
чтобы синхронизировать их выполнение.
*/
func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}

/*
 Программа будет постоянно выводить «ping» (нажмите enter, чтобы её остановить).
 Тип канала представлен ключевым словом chan, за которым следует тип, который будет
 передаваться по каналу (в данном случае мы передаем строки). Оператор <- (стрелка влево)
 используется для отправки и получения сообщений по каналу. Конструкция c <- "ping"
 означает отправку "ping", а msg := <- c — его получение и сохранение в переменную msg.
 Строка с fmt может быть записана другим способом: fmt.Println(<-c), тогда можно было
 бы удалить предыдущую строку.
*/
