package main

import (
	"fmt"
)

func main() {
	// 現状、はてブの人気エントリーのURLをスクレイピングして標準出力に出力するだけ。
	// スプシにコピペして使う。
	// スプシに自動で保存しようと思ってたけどそれはやめた。
	urls, err := crawl()
	if err != nil {
		panic(err)
	}
	for _, u := range urls {
		fmt.Println(u)
	}
}
