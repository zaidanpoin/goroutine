package belajar_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChanel(t *testing.T) {
	chanel := make(chan string)
	defer close(chanel)

	go func() {
		time.Sleep(2 * time.Second)

		fmt.Println("selesai mengirim data")
	}()

	p := <-chanel

	fmt.Println(p)
	time.Sleep(5 * time.Second)
}
