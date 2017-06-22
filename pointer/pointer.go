package main

import "fmt"

func main() {
	// intのポインタ変数を宣言
	var ptr *int
	// intの変数を宣言
	var i int = 12345
	// アドレス演算子を使ってiのアドレスをptrに代入
	ptr = &i

	fmt.Println("iの値:", i)     // 12345
	fmt.Println("iのアドレス:", &i) // 0xxxxxxx
	fmt.Println("ptrの値:", ptr) // &iと同じ(iのアドレス)
	fmt.Println("ポインタ経由のi値:", *ptr)

	// ポインタ経由でiの値を変更する
	*ptr = 999
	fmt.Println("ポインタ経由で変更したiの値:", i)
}
