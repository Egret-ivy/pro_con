package main

import (
	"fmt"
	"time"
)

var cnt = 0

func odd() {
	s := make([]int, 5)
	for i := 1; i <= 5; i++ {
		s[i-1] = cnt + (2*i - 1)
	}
	fmt.Println(s)
}

func even() {
	s := make([]int, 5)
	for i := 0; i <= 4; i++ {
		s[i] = cnt + 2*i
	}
	fmt.Println(s)
}

func main() {
	for cnt < 100 {
		go even()
		go odd()
		time.Sleep(time.Second)
		cnt += 10
	}

	/*不可以主函数里面只有go的命令 因为这样主函数一下子就结束了，
	而主函数的结束其他的goroutine强制结束*/
}
