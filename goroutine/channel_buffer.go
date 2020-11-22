package main

import (
	"fmt"
	"sync"
	"time"
)

/*
시간의 순서를 보장해 주지는 않는다.

*/

var wg1 sync.WaitGroup

func foo1(c chan int, someValue int) {
	defer wg1.Done()
	c <- someValue * 5

}

func main() {

	fooVal := make(chan int, 10)
	for i := 0; i < 10; i++ {
		wg1.Add(1)
		go foo1(fooVal, i)
	}

	wg1.Wait()
	close(fooVal)

	for item := range fooVal {
		fmt.Println(item)
	}
	time.Sleep(time.Second * 2)
}
