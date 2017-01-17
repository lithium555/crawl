package main

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

type Message struct {
	Text string
	V    bool
}

func main() {
	he := make(chan Message)
	she := make(chan Message)

	go man(he, she)
	go woman(he, she)

	//chat_sleep(time.Second*5)
	//FullStackPrintf("bla")
	time.Sleep(time.Minute * 1) //time.Second*10)
	FullStackPrintf("check")
}

func man(he, she chan Message) {
	garbage := <-she
	fmt.Printf("GARBAGE in channel she: '%v'\n", garbage)
	he <- Message{"Hello, Ivanna!!!", true}
	if garbage.V == true {
		fmt.Printf("Message from Ivanna: '%v'\n", garbage.Text)
	}
}

func woman(he, she chan Message) {
	she <- Message{"Hello, Slava !!!", true}
	garbage := <-he
	fmt.Printf("GARBAGE in channel he: '%v'\n", garbage)
	if garbage.V == true {
		fmt.Printf("Message from Slava: '%v'\n", garbage.Text)
	}
}

func chat_sleep(dt time.Duration) {
	fmt.Println("chat is sleeping for", dt)
	time.Sleep(dt)
	fmt.Println("chat got up")
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
