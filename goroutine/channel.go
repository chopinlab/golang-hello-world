package main

import "fmt"

/*
sentdex
https://www.youtube.com/watch?v=S11VFAMEs6E&list=PLQVvvaa0QuDeF3hP0wQoSxpkqgRcgxMqX&index=22
*/

// channel이 약간 포인터 느낌이네.. 파라미터로 받았는데, 그게 적용되는 것 보니
func foo(c chan int, someValue int) {
	c <- someValue * 5
}

func main() {

	fooVal := make(chan int, 2)

	go foo(fooVal, 5)
	go foo(fooVal, 3)

	// channel은 queue의 느낌이다.
	// block이 된다.
	v1, v2 := <-fooVal, <-fooVal

	fmt.Print(v1, v2)

}
