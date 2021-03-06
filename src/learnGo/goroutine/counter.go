package main

import (
	"fmt"
	"runtime"
)

var counter = 1

func count(ch chan int) {
	counter++
	fmt.Println("count ++:", counter)
	ch <- 1
}
func main() {

	//cpu nums
	fmt.Println(runtime.NumCPU())
	//set cup nums
	runtime.GOMAXPROCS(4)

	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go count(chs[i])
	}
	for _, ch := range chs {
		<-ch
	}
}
