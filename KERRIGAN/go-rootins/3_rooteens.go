package main

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strings"
)

type message struct {
	Text string
	V    bool
}

/*
по этому можно сделать так, что рутина 1 - инициатор для рутины 2,
 в то время как рутина 2 инициатор для 3 (она пошлет "ответ" не в канал для 1
 а в канал для 3), а 3 потом в конце ответит в канал 1,  чтобы она тоже могла завершиться (
*/
func main() {
	channel_1 := make(chan message)
	channel_2 := make(chan message)
	channel_3 := make(chan message)
	test1 := make(chan struct{})
	test2 := make(chan struct{})
	test3 := make(chan struct{})

	go rooteen_1(channel_1, channel_3, test1)
	go rooteen_2(channel_1, channel_2, test2)
	go rooteen_3(channel_2, channel_3, test3)

	<-test1
	<-test2
	<-test3
	//time.Sleep(time.Minute*1)//time.Second*10)
	FullStackPrintf("check")
	return
}

func rooteen_1(channel_1, channel_3 chan message, test1 chan struct{}) {
	channel_1 <- message{"Дупло, Дупло, Я гнездо, как слышишь?", true}
	mes := <-channel_3
	fmt.Printf("Message from channel3: '%v'\n", mes)
	test1 <- struct{}{}
}

func rooteen_2(channel_1, channel_2 chan message, test2 chan struct{}) {
	mes := <-channel_1
	fmt.Printf("Message from roteen_1: '%v'\n", mes)
	channel_2 <- message{"Гнездо, Гнездо, Я Дупло, слышу отлично!", true}
	test2 <- struct{}{}
}

func rooteen_3(channel_2, channel_3 chan message, test3 chan struct{}) {
	mes := <-channel_2
	fmt.Printf("Message from roteen_2: '%v'\n", mes)
	channel_3 <- mes
	test3 <- struct{}{}
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
