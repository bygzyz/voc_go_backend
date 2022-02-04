package main

import (
	"fmt"
	"sync"
)

/*
channel介绍：
用 make初始化 channel，make一般用来初始化 slice、map和channel
channel 可以关闭，向关闭的管道中继续push值会报错
channel 没有缓冲区时必须立即消费否则会产生死锁，channel为了保证多个goroutine之前的数据安全会获取锁。
*/
var wg2 sync.WaitGroup
var msg chan int

func consumer(queue chan int) {
	defer wg2.Done()
	fmt.Println(<-queue)

}

func main() {
	fmt.Println("进入了主函数")

	msg = make(chan int, 1)
	msg <- 1
	wg2.Add(1)
	go consumer(msg)
	msg <- 2
	fmt.Println(<-msg)
	close(msg)
	wg2.Wait()

}
