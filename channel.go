package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(ch chan int) {
	rand.Seed(time.Now().Unix())
	x := rand.Intn(1000)
	ch <- x
}

func consumer(ch chan int) {
	//var con int
	//con <- ch
	con := <-ch
	fmt.Println(con)
}

func main() {
	data := make(chan int)
	go producer(data)
	go consumer(data)
	time.Sleep(time.Second)
}
