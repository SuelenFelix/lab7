package main

import (
	"time"
	"fmt"
	"math/rand"
)

func exec(max_sleep_ms int) int{
	sleep := rand.Intn(max_sleep_ms)
	time.Sleep(time.Duration(sleep) * time.Millisecond)
	return sleep
}

func async(ch chan int, max_sleep_ms int){
	for i := 0;i < 1000; i++{
		ch <- exec(max_sleep_ms)
	}
}

func aux(max_sleep_ms int) chan int{
	ch := make(chan int)
	go async(ch, max_sleep_ms)
	return ch
}

func main(){ 
	ch1 := aux(20)
	ch2 := aux(200) 
	sum := 0	

	for i := 0;i < 500; i++{
		select {
		case <- ch1:
			sum = sum + <- ch1
		case <- ch2:
			sum = sum + <- ch2
		default:
			sum = sum +  <- ch1 + <- ch2
		}
	}

	fmt.Println(sum)


}