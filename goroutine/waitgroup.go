package main

import (
	"fmt"
	"sync"
	"time"
)

/*
- WaitGroup의 용도

1. 여러 일을 동시에 실행할 때 좋다.
2. 하나의 DB에서 데이터를 가지고 오고, 또 다른 종류의 DB에서 데이터를 가지고 온 후
그 두 개의 데이터를 결합하여 다른 서버에 전달하는 경우에 필요할 것 같음
3. 고루틴은 결국 시간이 오래 걸리는 작업에 필요함

- 주의점
만약에 에러가 나서 wg.Done이 안먹히는 경우
무기한의 pending이 뜰 수 있다.
--> Defer을 사용하여 해결함

*/

var wg sync.WaitGroup

func cleanup() {

	defer wg.Done()
	if r := recover(); r != nil {
		fmt.Println("Recovered in cleanup", r)
	}
}

// 3번 반복해서 0.1초 간격으로 출력한다.
func say(s string) {
	defer cleanup()

	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
		if i == 2 {
			panic("Oh dear, a 2")
		}
	}
}

func main() {

	wg.Add(1)
	go say("Hey")
	wg.Add(1)
	go say("There")
	wg.Wait()

	wg.Add(1)
	go say("Hi")
	wg.Wait()

	time.Sleep(time.Second)
}
