package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 評価に用いるページを決定するためのコード
// スプシにコピーしやすいように改行をいれてある
func main() {
	ms := []int{616}
	for j, m := range ms {
    fmt.Printf("%dです！\n", j)
		for i := 0; i < 10; i++ {
			rand.Seed(time.Now().UnixNano())
			fmt.Println(rand.Intn(m-1) + 1)
    fmt.Println("")
    fmt.Println("")
		}
    fmt.Println("")
	}
}

