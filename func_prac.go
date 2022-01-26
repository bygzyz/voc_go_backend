/*
go 中函数是一等公民，可以作为函数参数、返回值、赋值给某个变量、
*/

package main

import "fmt"

func score(a int, f func(a int) bool) bool {
	if f(a) {
		return true
	} else {
		return false
	}
}

func filter(a []int, f func(a int) bool) []int {
	// make 初始化一个切片必须要指定长度
	goodScore := make([]int, 0)
	for _, v := range a {
		if f(v) {
			goodScore = append(goodScore, v)
		}
	}
	return goodScore
}

func main() {
	//res := score(60, func(a int) bool {
	//	if a >= 61 {
	//		return true
	//	} else {
	//		return false
	//	}
	//})
	//fmt.Println(res)

	res := filter([]int{60, 70, 80, 90}, func(a int) bool {
		if a >= 80 {
			return true
		} else {
			return false
		}
	})
	fmt.Println(res)

}
