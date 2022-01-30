//package main
//
//import (
//	"fmt"
//)
//
//func ChangeMap(m map[int]int) {
//	m[1] = 2
//}
//
//func main() {
//	var a = make(map[int]int)
//	a[1] = 1
//	fmt.Println(a)
//	ChangeMap(a)
//	fmt.Println(a)
//}

//package main
//
//import (
//	"fmt"
//)
//
//func ChangeMap(m map[int]int) {
//	m = nil
//}
//
//func main() {
//	var a map[int]int = make(map[int]int)
//	a[1] = 1
//	fmt.Println(a)
//	ChangeMap(a)
//	fmt.Println(a)
//}

//package main
//
//import (
//"fmt"
//)
//
//func ChangeMap(m map[int]int) {
//	m = nil
//}
//
//func main() {
//	var a map[int]int
//
//	fmt.Println(a==nil)
//}

// panic recover go中把错误和异常区分开来，我们在写程序的时候可预知到的是错误，显示将它捕获并进行处理；不可预知的叫做异常。

//package main
//
//import (
//	"fmt"
//)
//
//func ChangeMap(m map[int]int) {
//	m = nil
//}
//
//func main() {
//	var a map[int]int
//
//	for _,value := range a {
//		fmt.Println(value)
//	}
//}

// 子协程中panic无法被外层recover捕获
package main

import "fmt"

func compute_a() interface{} {
	return 1
}

func main() {
	if a := compute_a(); a == nil {
		fmt.Println("a等于nil")
	} else {
		fmt.Println("a不等于nil")
	}
	d := []int{1, 2}
	b := make([]int, 0, 2)
	for i := range d {
		b = append(b, i)
	}
	fmt.Println(b)

}
