package belajar_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func Run() {
	fmt.Println("hello world")
}

func TestCreateGoroutine(t *testing.T) {
	go Run()
	fmt.Println("ups")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("display : ", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
