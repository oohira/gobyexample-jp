// _ゴルーチン (goroutine)_ は、軽量スレッドです。

package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// 関数呼び出し `f(s)` があるとしましょう。
	// 通常の方法で呼び出すと、同期的に実行されます。
	f("direct")

	// この関数をゴルーチンとして呼び出すには、
	// `go f(s)` とします。この新しいゴルーチンは、
	// 実行元とは並行して実行されます。
	go f("goroutine")

	// 無名関数に対してもゴルーチンを開始できます。
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 上の 2 つの関数呼び出しは別々のゴルーチンで非同期に
	// 実行されるので、プログラムの実行はすぐにここへきます。
	// それらが完了するまで待ちます（より堅牢な方法としては、
	// [WaitGroup](waitgroups) を使ってください）。
	time.Sleep(time.Second)
	fmt.Println("done")
}
