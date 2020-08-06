package gorotine

import (
	"fmt"
	"sync"
	"time"
)

var chanInt chan int = make(chan int, 10)
var timeout chan bool = make(chan bool)

var WG sync.WaitGroup

func Loop() {
	for i := 1; i < 11; i++ {
		time.Sleep(time.Microsecond * 10)
		fmt.Println(i)
	}
}

func Send() {
	time.Sleep(time.Second * 1)
	chanInt <- 1
	time.Sleep(time.Second * 1)
	chanInt <- 2
	time.Sleep(time.Second * 1)
	chanInt <- 3
	time.Sleep(time.Second * 2)
	timeout <- true
}

func Receive() {
	//num1 := <-chanInt
	//fmt.Println("num:", num1)
	//num2 := <-chanInt
	//fmt.Println("num:", num2)
	//num3 := <-chanInt
	//fmt.Println("num:", num3)

	for {
		select {
		case num := <-chanInt:
			fmt.Println("num:", num)
		case <-timeout:
			fmt.Println("timeout:", timeout)
		}
	}
}

func Read() {
	for i := 0; i < 3; i++ {
		WG.Add(1)
	}
}

func Write() {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second * 2)
		fmt.Println("Done ->", i)
		WG.Done()
	}
}
