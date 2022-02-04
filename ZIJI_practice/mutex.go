package main

import (
	"fmt"
	"sync"
)

/*
互斥锁同步协程
*/

var total int
var wg1 sync.WaitGroup
var lock sync.Mutex

func add() {
	defer wg1.Done()
	for i := 0; i < 10000; i++ {
		lock.Lock()
		//加1
		total = total + 1
		lock.Unlock()
	}

}

func sub() {
	defer wg1.Done()
	for i := 0; i < 10000; i++ {
		lock.Lock()
		//减1
		total = total - 1
		lock.Unlock()
	}

}

func main() {

	wg1.Add(2)
	go add()
	go sub()
	wg1.Wait()
	fmt.Println("最终的total是", total)

}
