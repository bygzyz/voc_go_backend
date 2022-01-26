package main

import "fmt"

func main() {
	//a := 1
	//b := &a //  &代表取出a变量的地址，a本身可以是一个普通变量也可以是一个指针变量
	//c := *b // * 取地址对应的值
	//println(a)
	//fmt.Printf("%T\n",b)
	//fmt.Printf("%T\n",c)
	//println(c)
	//res, _ := http.Get("http://1.117.229.170:9200/voc_yuwell")
	//fmt.Println(res.StatusCode)

	var a map[string]interface{} = map[string]interface{}{
		"1": "1",
	}
	fmt.Println(a)
}
