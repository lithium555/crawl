package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"strings"
	"time"
)

/*
Делают иначе - создают фиксированное количество рутин, скажем 10
и заставляют их в цикле брать с общего канала задачи (в данном случае номера),
выполнять работу, отдавать результаты а другой общий канал и опять делать все сначала

Мейн же запустив рутины пишет задачи в канал и собирает ответы

В конце командует рутинам выйти

Идея такая: ты делаешь новый канал с задачами, и сразу запускаешь рутины
которые в вечном цикле берут из этого канала
И только потом после их запуска начинаешь из мейна кидать в тот канал задачи


Иногда ты с помощью селекта можешь проигнорировать то что послала рутина и выйти,
например. Но в этом случае ты оставишь повисшую рутину, а это утечка памяти. Для того
чтобы это не происходило либо ставят еще канал чтобы сказать той рутине остановиться,
либо делают канал с буфером, чтобы рутина при любых обстоятельствах могла положить
сообщение и выйти (а ненужный канал с сообщением в буфере заберет GC)

Не совсем) Ты можешь не давать им четкого правила, а просто кидать в канал задачи,
а рутины будут их забирать по мере возможности. В остальном все так же - рутины отдают
все в Мейн как и раньше, а он делает слайс.

Идея такая: ты делаешь новый канал с задачами, и сразу запускаешь рутины
 которые в вечном цикле берут из этого канала

И только потом после их запуска начинаешь из мейна кидать в тот канал задачи
*/

type data struct {
	info int
	err  error
}

const taskNumber = 100
const routinesNumber = 10

func main() {
	start := time.Now()

	chanel := make(chan data, taskNumber-11)
	//chanelERR := make(chan error, taskNumber-90)
	taskCHANEEL := make(chan int, taskNumber-9)
	for k := 0; k < routinesNumber; k++ {
		go master(taskCHANEEL, chanel)
	}
	for i := 0; i < taskNumber; i++ {
		taskCHANEEL <- i
	}
	var slice []int
	var sliErr []error
	for j := 0; j < taskNumber; j++ {
		//select {
		//case here := <-chanel:
		//	slice = append(slice, here)
		//	fmt.Printf("HERE IS: '%v'\n", here)
		//case receive_err := <-chanelERR:
		//	sliErr = append(sliErr, receive_err)
		//}

		z := <-chanel
		slice = append(slice, z.info)
		fmt.Printf("HERE IS: '%v'\n", z)
		sliErr = append(sliErr, z.err)

	}
	log.Printf("SLICE: '%v'\n", slice)
	log.Printf("SliERR: '%v'\n", sliErr)

	dt := time.Since(start)
	log.Printf("THE TIME IS: '%v'\n", dt)

	FullStackPrintf("check")
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
	time.Sleep(time.Second/4 + time.Duration(rand.Intn(int(time.Second*1))))
	rez := rand.Intn(10)
	if rez >= 7 {
		return 0, fmt.Errorf("error, number is too large: %v, routine %d", rez, i)
	}
	return rez, nil
}

func master(taskCHANEEL chan int, chanel chan data) {
	for {
		i := <-taskCHANEEL
		rezult, err := calculate(i)
		if err != nil {
			log.Printf("CONTINUE '%v'", i)
			chanel <- data{0, err}
			continue
		}
		chanel <- data{rezult, nil}
		fmt.Printf("Number of rooten: '%v'\n", i)
	}
}

func stackPrintf(full bool, format string, args ...interface{}) {
	buf := bytes.NewBuffer(nil)
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Fprintf(buf, format, args...)
	buf.Write(Stack(full))
	os.Stdout.Write(buf.Bytes())
}
func StackPrintf(format string, args ...interface{}) {
	stackPrintf(false, format, args...)
}

func FullStackPrintf(format string, args ...interface{}) {
	stackPrintf(true, format, args...)
}

func Stack(full bool) []byte {
	buf := make([]byte, 1024)
	for {
		n := runtime.Stack(buf, full)
		if n < len(buf) {
			return buf[:n]
		}
		buf = make([]byte, 2*len(buf))
	}
}
