package main

import "fmt"

func main() {
	// 独自型を宣言する
	type myType int
	// 独自型の変数をつくる
	var i myType = 10
	fmt.Printf("%s\n", i) // main.myType=10

	// 型を変換する
	var u uint32 = uint32(i)
	fmt.Printf("%s", u) // uint32=10
}
