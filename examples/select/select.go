// Go の _select_ を使うと、複数のチャネル操作を待つことができます。
// ゴルーチンとチャネルを `select` で扱えるのが、Go の強力な特長です。

package main

import (
	"fmt"
	"time"
)

func main() {

	// 2 つのチャネルに対して `select` する例を見ていきます。
	c1 := make(chan string)
	c2 := make(chan string)

	// 各チャネルは、一定時間後に値を受信します。これは、
	// 例えば同期的な RPC 操作をゴルーチンで並行実行する場合を
	// シミュレートしています。
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// これらの値を同時に待つために `select` を使い、
	// 受信したものから画面に表示します。
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
