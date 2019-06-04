// Go は、
// <a href="http://en.wikipedia.org/wiki/Recursion_(computer_science)"><em>再帰関数 (recursive functions)</em></a>
// をサポートします。これは典型的な階乗関数の例です。

package main

import "fmt"

// この `fact` 関数は、終了条件となる `fact(0)`
// に達するまで、自分自身を呼び出します。
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
}
