package main

import (
	"fmt"
	"sync"
	"time"
)

/*
互斥锁能够将敏感的数字加减原子化，从而保证数据的准确性，
但是加锁对性能的影响较大；
以下介绍一种新锁，为读写锁；读锁之前无影响，一旦写锁启动便无法再读
*/

var wg sync.WaitGroup
var rwLock sync.RWMutex

func read() {
	defer wg.Done()
	rwLock.RLock()

	fmt.Println("开始读数据")
	time.Sleep(time.Second)

	rwLock.RUnlock()

}

func write() {
	defer wg.Done()
	fmt.Println("开始写数据")
	time.Sleep(time.Second * 5)
}

func main() {
	wg.Add(6)
	for i := 0; i < 5; i++ {
		go read()
	}
	for i := 0; i < 1; i++ {
		go write()
	}

	wg.Wait()
}
