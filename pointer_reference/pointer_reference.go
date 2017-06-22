package main

import "fmt"

func main() {
	a, b := 1, 1
	// 引数を使い関数に変数を渡す
	// aは値をそのまま渡す(値渡し)
	// bはアドレス演算子を使ってポインタとして渡す(ポインタ渡し/参照渡し)
	double(a, &b)

	fmt.Println("値渡し:", a)
	fmt.Println("参照渡し:", b)

	// new関数で新しいメモリを割り当てる
	var i *int = new(int)
	fmt.Printf("%s", "ポインタiはのアドレス:", i) // *int=0xc42000e270
}

// パラメータxは値のコピーを受け取る
// パラメータyはポインタのコピーを受け取る
func double(x int, y *int) {
	// 値のコピーを変更する
	x = x * 2
	// 間接参照演算子を利用してポインタyが指し示す変数の値を変更する
	*y = *y * 2
}
