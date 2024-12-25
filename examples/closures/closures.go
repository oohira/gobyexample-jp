// Go は、 [_無名関数 (anonymous functions)_](https://en.wikipedia.org/wiki/Anonymous_function)
// をサポートするため、<a href="https://en.wikipedia.org/wiki/Closure_(computer_science)"><em>クロージャ (closures)</em></a>
// を作れます。無名関数は、名前をつけずに
// インラインで関数を定義したい場合に便利です。

package main

import "fmt"

// この `intSeq` 関数は、 `intSeq` の中で定義した
// 無名関数を返します。返された関数は、変数 `i`
// を _閉じて (closes over)_ クロージャを作ります。
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	// `intSeq` を呼び出した結果 (関数) を `nextInt` に代入します。
	// この関数は、自分自身の `i` の値をキャプチャし、その値は
	// `nextInt` を呼び出すたびに更新されます。
	nextInt := intSeq()

	// `nextInt` を何回か呼び出して、クロージャの効果を見てください。
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// 関数ごとに個別の状態をもっていることを確認するために、
	// もう 1 つ新たに作成して試してみてください。
	newInts := intSeq()
	fmt.Println(newInts())
}
