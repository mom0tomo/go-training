package main

import "fmt"

func main() {

LBL:
	for v := 0; v < 3; v++ {
		fmt.Println(v)

		for j := 0; j < 3; j++ {
			fmt.Println("	", j)

			if v == 1 && j == 1 {
				// 外側のforへcontinue
				continue LBL

			}
		}
	}
}
