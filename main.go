package main

import (
	"fmt"
	"sync"
	"time"
)

func printRoutine(data chan string) {
	res := <-data
	fmt.Println(res)
}

func newRoutine(data chan string) {
	data <- "from other routines"
}

func main() {
	fmt.Println("goroutine pertama")
	data1 := make(chan string)
	go printRoutine(data1)
	go newRoutine(data1)

	time.Sleep(time.Second * 1)

	fmt.Println("goroutine kedua")
	data2 := make(chan string, 2)
	data2 <- "a"
	data2 <- "b"
	fmt.Println(<-data2)
	fmt.Println(<-data2)

	fmt.Println("goroutine ketiga")
	data3 := make(chan string, 2)
	go func(n int) {
		defer close(data3)
		for i := 0; i < n; i++ {
			data3 <- "a"
		}
	}(cap(data3))

	for val := range data3 {
		fmt.Println(val)
	}

	fmt.Println("goroutine keempat")
	err := make(chan string)
	data4 := make(chan int)

	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func(w *sync.WaitGroup, d chan int, e chan string) {
		defer w.Done()
		for i := 0; i < 10; i++ {
			d <- i
		}
		close(d)
		e <- "error"
	}(wg, data4, err)

	go func(w *sync.WaitGroup, d chan int, e chan string) {
		defer w.Done()
		for {
			select {
			case res := <-e:
				fmt.Println(res)
				return
			case datas := <-d:
				fmt.Println(datas)
			}
		}
	}(wg, data4, err)

	wg.Wait() // Menunggu kedua goroutine selesai
}
