package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go f2(i)
	}
	var input string
	fmt.Scanln(&input)
}

func f2(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		//добавим небольшую задержку функции с помощью функции time.Sleep и rand.Intn:
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
		//f выводит числа от 0 до 10, ожидая от 0 до 250 мс после каждой операции вывода.
	}
}
