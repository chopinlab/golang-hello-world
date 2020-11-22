package main

import (
	"fmt"
	"sync"
	"time"
)

/*
https://lynlab.co.kr/blog/82
*/

var globalValue int

//var wg2 sync.WaitGroup

func main() {

	//ch := make(chan int)
	var wg2 sync.WaitGroup
	var mutex sync.Mutex

	startTime := time.Now()

	for i := 0; i < 100; i++ {

		wg2.Add(1)
		go action(i, &mutex, &wg2)
	}

	wg2.Wait()
	delta := time.Now().Sub(startTime)
	fmt.Printf("Result is %d, done in %.3fs.\n", globalValue, delta.Seconds())
}

func action(i int, mutex *sync.Mutex, wg2 *sync.WaitGroup) {
	defer wg2.Done()

	mutex.Lock()
	globalValue += i
	mutex.Unlock()

	time.Sleep(1 * time.Second)
}
