// `for` は、Go の唯一のループ構文です。
// いくつかの基本的な `for` ループは次の通りです。

package main

import "fmt"

func main() {

	// 条件式しかない最も基本的なタイプです。
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// 初期化・条件式・後処理をもつ典型的な `for` ループです。
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// "N 回繰り返す" を実現するための別の方法は、整数に対して
	// `range` を使うことです。
	for i := range 3 {
		fmt.Println("range", i)
	}

	// 条件式がない `for` は、 `break` でループから抜けるか、
	// `return` によって関数自体から抜けるまでループし続けます。
	for {
		fmt.Println("loop")
		break
	}

	// `continue` でループの次の実行に進むこともできます。
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
