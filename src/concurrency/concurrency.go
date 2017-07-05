package concurrency

import (
	"fmt"
	"sync"
	"time"
)

//Smiple func. Add twenty.
func Add20(x int, ch chan int) {
	for i := 0; i < 5; i++ {
		x += 4
		fmt.Printf("%v\n", x)
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println("Done")
	ch <- x //This is used to send one piece of data to the caller
}

//Add 20 to x and send to c
func Add20ToChan(x int, c chan int) {
	c <- x + 20
}

//This function simply sends multiple ints into the channel
func Send(ch chan int) {
	ch <- 20
	ch <- 3
	ch <- -43
	ch <- 15
}

func Fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y //This line demonstrates that all values are calculated before being assigned
	}
	close(c) //This is needed to tell range that the channel will not be sending any more data
}

//A select is like a switch, but for channels
//If multiple channels are ready at once, it chooses randomly and only executes one per loop
//Can use an infinite for to process multiple channel sends with one select block
func Select(x <-chan int, y <-chan int, quit chan string) { //Specify that two of these are only to be used to receive
	j, k, count := 0, 0, 0

	for {
		count++
		select {
		case i := <-x:
			// 			fmt.Printf("x selected, %v\n", count)
			j += i
		case i := <-y:
			// 			fmt.Printf("y selected, %v\n", count)
			k += i
		case msg := <-quit:
			fmt.Printf("j %v, k %v\n", j, k)
			fmt.Println(msg)
			return
		//The default case executes when none of the channel cases have any data queued
		//This can be used to kickoff an alternative workflow and stack work in the queue,
		//so that the goroutine works in batches
		default:
			fmt.Printf("j %v, k %v\n", j, k)
			count = 0
			//This will cause sends to buffer for a second
			time.Sleep(1 * time.Second)
		}
	}
}

type SafeCounter struct {
	X   int
	mux sync.Mutex
}

func (sc *SafeCounter) Inc() {
	sc.mux.Lock()
	defer sc.mux.Unlock()
	sc.X++
}

func (sc *SafeCounter) Value() int {
	sc.mux.Lock()
	defer sc.mux.Unlock()
	return sc.X
}
