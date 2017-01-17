package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

type message2 struct {
	Text string
}

/*
по этому можно сделать так, что рутина 1 - инициатор для рутины 2,
 в то время как рутина 2 инициатор для 3 (она пошлет "ответ" не в канал для 1
 а в канал для 3), а 3 потом в конце ответит в канал 1,  чтобы она тоже могла завершиться (
*/
func main() {
	start := time.Now()

	channel_1 := make(chan message2)
	channel_2 := make(chan message2)
	channel_3 := make(chan message2)
	testChannel := make(chan struct{})

	go rooteen1(channel_1, channel_3, testChannel)
	go rooteen2(channel_1, channel_2, testChannel)
	go rooteen3(channel_2, channel_3, testChannel)

	for i := 0; i < 3; i++ {
		<-testChannel
	}
	//time.Sleep(time.Minute*1)//time.Second*10)
	FullStackPrintf("check")
	dt := time.Since(start)
	log.Printf("THE TIME IS: '%v'\n", dt)
	return
}

func rooteen1(channel_1, channel_3 chan message2, testChannel chan struct{}) {
	channel_1 <- message2{"Дупло, Дупло, Я гнездо, как слышишь?"}
	mes := <-channel_3
	fmt.Printf("Message from channel3: '%v'\n", mes)
	testChannel <- struct{}{}
}

func rooteen2(channel_1, channel_2 chan message2, testChannel chan struct{}) {
	mes := <-channel_1
	fmt.Printf("Message from roteen_1: '%v'\n", mes)
	channel_2 <- message2{"Гнездо, Гнездо, Я Дупло, слышу отлично!"}
	testChannel <- struct{}{}
}

func rooteen3(channel_2, channel_3 chan message2, testChannel chan struct{}) {
	mes := <-channel_2
	fmt.Printf("Message from roteen_2: '%v'\n", mes)
	channel_3 <- mes
	testChannel <- struct{}{}
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
