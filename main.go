package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	// // deklarasi channel
	// var messages = make(chan string)

	// var sayHelloTo = func(who string) {
	// 	time.Sleep(5 * time.Second)
	// 	var data = fmt.Sprintf("hello %s", who)
	// 	// send data ke channel 'message'
	// 	messages <- data
	// }

	// // jalankan fungsi 'sayHelloTo' dalam goroutine (concurrent)
	// go sayHelloTo("Imam Permana")
	// go sayHelloTo("Budi")
	// go sayHelloTo("Anya Geraldine")

	// // receive data dari channel 'messages'
	// fmt.Println(<-messages)
	// fmt.Println(<-messages)
	// fmt.Println(<-messages)

	// time.Sleep(2 * time.Second)

	// runtime.GOMAXPROCS(1)
	// var start = time.Now()

	// fmt.Println("main started at time ", time.Since(start))
	// c := make(chan string)
	// go func() {
	// 	time.Sleep(10 * time.Second)
	// 	fmt.Printf("hello from goroutine, at time %v\n", time.Since(start))
	// 	c <- "goroutine say hi"
	// }()

	// fmt.Printf("goroutine sent this: '%v'. At time %v\n", <-c, time.Since(start))
	// fmt.Printf("main stopped at time %v\n", time.Since(start))

	var message = make(chan string)

	nama := []string{"Imam", "Permana", "Budi", "Anya"}

	for _, n := range nama {
		go hello(message, n)
	}

	for i := 0; i < len(nama); i++ {
		fmt.Println(<-message)
	}
	fmt.Println(message)

}

func hello(msg chan string, nama string) {
	msg <- nama
}
