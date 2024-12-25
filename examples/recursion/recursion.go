// Go は、
// <a href="https://en.wikipedia.org/wiki/Recursion_(computer_science)"><em>再帰関数 (recursive functions)</em></a>.
// をサポートします。これは典型的な例です。

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

	// 無名関数も再帰的にできます。ただし、関数を保存するための変数を
	// `var` であらかじめ宣言する必要があります。
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		// `fib` は `main` ですでに宣言されているので、Go
		// はここでどの関数を呼び出すべきか知っています。
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
