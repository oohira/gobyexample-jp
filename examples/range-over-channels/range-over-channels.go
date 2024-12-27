// [以前](range-over-built-in-types) の例で、`for` と `range` が基本的なデータ構造に
// 対して、どのように反復処理を提供するかを見ました。
// この構文は、チャネルから受信した値を反復処理する場合にも
// 使うことができます。

package main

import "fmt"

func main() {

	// `queue` チャネルの 2 つの値を反復処理するとします。
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// この `range` は、`queue` から受信した要素を反復処理します。
	// 上でチャネルを `close` したので、反復処理は 2 つの要素を
	// 受信したときに終了します。
	for elem := range queue {
		fmt.Println(elem)
	}
}
