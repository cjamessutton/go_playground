package concurrency_test

import (
	"concurrency"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestAdd20(t *testing.T) {
	done := make(chan int)
	go concurrency.Add20(10, done)
	fmt.Println("Add 20!")
	<-done //This line keeps the caller from closing before the goroutines finish executing
}

func TestChannel(t *testing.T) {
	c := make(chan int)
	go concurrency.Add20ToChan(20, c)
	go concurrency.Add20ToChan(-15, c)
	x, y := <-c, <-c
	fmt.Printf("%v + %v = %v\n", x, y, x+y)
}

func TestBuffer(t *testing.T) {
	ch := make(chan int, 2)
	go concurrency.Send(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func TestFibonacci(t *testing.T) {
	c := make(chan int, 10)
	go concurrency.Fibonacci(10, c)
	for i := range c {
		fmt.Println(i)
	}
}

func TestSelect(t *testing.T) {
	x, y, quit := make(chan int, 20), make(chan int, 20), make(chan string)
	go concurrency.Select(x, y, quit)
	for i := 0; i < 10; i++ {
		r := rand.Intn(3)
		send := rand.Intn(100)
		switch r {
		case 0:
			x <- send
		case 1:
			y <- send
		default:
			x <- send
			y <- send
		}
	}
	quit <- "Closing"
	time.Sleep(2 * time.Millisecond)
}

func TestSafeCounter(t *testing.T) {
	sc := concurrency.SafeCounter{X: 0}
	for i := 0; i < 100; i++ { //Will call Inc 100x
		go sc.Inc()
	}

	time.Sleep(2 * time.Second)
	fmt.Printf("%v\n", sc.Value())
}
