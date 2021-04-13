package main

import "fmt"

func main() {
	var v int
	fmt.Println("请输入一个整型：")
	fmt.Scanf("%d", &v)
	//fmt.Scan(&v)
	fmt.Println("v = ", v)
}

func mainOutput() {

	/*直接输出*/
	printDemo := "Hello World"
	fmt.Print("直接输出：" + printDemo)
	// 输出换行
	fmt.Print("\n")

	/*格式化输出*/
	intDemo := 1
	floatDemo := 1.1
	stringDemo := "HelloWorld"
	charDemo := 't'
	boolDemo := false
	byteDemo := 1

	fmt.Printf("%d,%f,%s,%c,%t,%b\n", intDemo, floatDemo, stringDemo, charDemo, boolDemo, byteDemo)

	/*换行输出*/
	printlnDemo := "你好"
	fmt.Println(printlnDemo)
}
