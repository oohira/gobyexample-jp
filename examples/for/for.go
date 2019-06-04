// `for` は、Go の唯一のループ構文です。
// 3 種類の基本的な `for` ループは次の通りです。

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
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// 条件式がない `for` は、 `break` でループから抜けるか、
	// `return` で関数自体から抜けるまでループし続けます。
	for {
		fmt.Println("loop")
		break
	}

	// `continue` でループの次の実行に進むこともできます。
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
