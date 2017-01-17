package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"sort"
	"strings"
	"sync"
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

	chanel := make(chan data)
	taskCHANEEL := make(chan int)

	var wg sync.WaitGroup
	wg.Add(routinesNumber)
	for k := 0; k < routinesNumber; k++ {
		go func() {
			defer wg.Done()
			master(taskCHANEEL, chanel)
		}()
	}
	/*
	   суть такая:
	   1) создается вейт-группа `var wg sync.WaitGroup`
	   2) *перед* запуском рутины которую нужно подождать делается `wg.Add(1)`
	   3) рутина делает сразу `defer wg.Done()` чтобы сообщить через вейт-группу что она закончила
	   4) мейн чтобы подождать рутины выполняет `wg.Wait()`. в этот момент он повиснет пока все рутины
	   не сделают `wg.Done`

	   1)Аргумент в `wg.Add` это количество работы (и вызовов `wg.Done`), а не номер рутины
	   если ты скажешь wg.Add(3), то потом нужно вызвать трижды Done чтобы Wait освободить
	   2) Разные переменные `sync.WaitGroup` считают рутины *отдельно* и никак не связаны, плюс,
	   эти объекты нельзя копировать, иначе опять же, они перестанут быть связаны
	   3) И да, лучше чтобы та функция что объявляет `sync.WaitGroup` ставила все рутины в анонимные
	   функции в которых сразу видно что они в конце вызывают `Done`. Это уменьшит количество возможных
	   ошибок, плюс сразу видно кто вызывает эти `Done`. (edited)


	   	   так что есть два решения:
	   	2) заставить мейн одновременно писать задачи и собирать ответы в одном цикле,
	   	и убрать буфер опять же
	*/
	var slice []int
	var sliErr []error
	var writeSlice []int
	var (
		sent int
		got  int
		//sendCh = taskCHANEEL
	)
	/*
	   	суть такая: мейн отдает сколько-то задач, получает сколько-то ответов и
	   	стоит ждет на wg.Wait(), но получается что он не дождался последних ответов
	   	(или рутины отвечают лишний раз) и т.к. рутины повисли на отправке они не
	   	выходят и мейн не выходят ожидая их

	   	там явно что-то не так с количеством задач/ответов

	     	что-то не совпадает
	*/
	for sent < taskNumber || got < taskNumber {
		// кидаем задачу либо получаем ответ
		i := sent // (количество отправленных задач) одновременно является индексом следующей *не* отправленной
		//если 3 задачи отправлено (0,1,2), то номер следующей (3) равен количеству уже отправленных
		select {
		case taskCHANEEL <- i:
			sent++
			if sent == taskNumber {
				close(taskCHANEEL)
				taskCHANEEL = nil // выключаем канал, если все отправлено
				/*
					использует тот факт что по нулевым каналам никогда нельзя отправить
					или получить данные (edited)
					их так можно "отключать"

					вначале присвоить настоящий канал (tasks), а потом когда что-то случиться
					(все задачи отправлены),
					ставим его на `nil` и первая ветка `select` перестает вообще срабатывать
				*/
			}
		case response := <-chanel:
			got++
			if response.err != nil {
				sliErr = append(sliErr, response.err)
			} else {
				slice = append(slice, response.info)
				//writeSlice = append(writeSlice, i)
			}
		}
	}

	/*
		нужно короче делать там "вечный" цикл
	3*taskNumber
		и следить отдельно за тем сколько он отправил задач и сколько получил ответов
	*/
	sort.Ints(slice)
	log.Printf("SLICE: '%v'\n", slice)
	log.Printf("SliERR: '%v'\n", sliErr)

	dt := time.Since(start)
	log.Printf("THE TIME IS: '%v'\n", dt)

	log.Printf("Th elength od slice = '%v'\n", len(slice))
	log.Printf("writeSlice= '%v'\n", writeSlice)
	log.Printf("sent: '%v'\n", sent)
	log.Printf("got: '%v'\n", got)
	FullStackPrintf("check")
	wg.Wait()
	fmt.Println("Done")
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
	return i, nil
}

func master(taskCHANEEL chan int, chanel chan data) {
	for {
		i, ok := <-taskCHANEEL
		if !ok {
			return
		}
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

/*
  for k := 0; k < routinesNumber; k++ {
        go master(taskCHANEEL, chanel)
    }
    go getTask(taskCHANEEL)
да, тут такие вещи происходят:
getTask распределяет задачи по рутинам, но блокируется пока задачи не соберут
каждая рутина master забирает задачу тем самым разблокирует getTask и транслирует каждую такую задачу
в результат, но блокируется при отправке по второму каналу
мэйн собирает результаты со второго канала, освобождая рутины

[8:14]
и у тебя может быть еще и десятки таких "слоев" и ничего не будет ломаться пока ты четко себе
представляешь как взаимодействуют соседние рутины

суть в том, что все переменные которые видит рутина не будут собраны GC пока рутина не выйдет
значит если рутина висит она все еще держит память
а тут таких 10

т.е. если канал из интов, то он блокируется при попытке получения, как обычно,
но если его закрыть, то он всегда моментально возвращает 0

[7:10]
толк в том, что он в таком режиме не "заканчивается" (можно вечно получать) и не блокирует рутины

в нашем примере мы передаем рутинам задачи через канал:
```for {
   i := <- ch
}

чтобы определить закрыт ли он, можно добавить "ок" как в случае с мапами:
```for {
  i, ok := <-ch
  if !ok {
    return
  }

  6. лучше для каждой передачи - записи и считывани яоткрывать новую пару каналов
Просто стоит следить ели ты закрываешь канал, то делаешь это только с рутины которая пишет.
И только однажды. Но даже если прога упадет она все равно скажет где именно она пробовала его
закрыть и откуда пришла, так что можно это исправить.



Вот последняя недоделка это выключение рутин

[10:16]
вторая штука - нужно чтобы мейн дождался пока рутины выйдут

		for j := 0; j < taskNumber; j++ {
		//	log.Printf("Value j AT THE BEGINING of cycle = '%v'\n", j)
			select {
			case taskCHANEEL <- j:
				sent++
				writeSlice = append(writeSlice, j)
				if sent == taskNumber{
					taskCHANEEL = nil// выключаем канал, если все отправлено
				}
			//	fmt.Println("You are in case taskCHANEEL")
			//	fmt.Printf("j = '%v'\n", j)
			case	z, ok := <-chanel:
				got++
				if !ok{
					return
				}
				slice = append(slice, z.info)
				fmt.Printf("HERE IS: '%v'\n", z)
				sliErr = append(sliErr, z.err)
				taskCHANEEL <- j
				writeSlice = append(writeSlice, j)
				//fmt.Println("Youa are in case chanel")
				//fmt.Printf("j = '%v'\n", j)
			}
//			log.Printf("Value j in THE END of cycle = '%v'\n", j)
		}
*/
