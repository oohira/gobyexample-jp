// Go では、 _変数 (variables)_ は明示的に宣言され、コンパイラによって
// 関数呼び出しの型チェックなどに使われます。

package main

import "fmt"

func main() {

	// `var` は 1 つ以上の変数を宣言します。
	var a = "initial"
	fmt.Println(a)

	// 複数の変数を一度に宣言することもできます。
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go は初期化された変数の型を推論できます。
	var d = true
	fmt.Println(d)

	// 初期化せずに宣言された変数は、 _ゼロ値 (zero-valued)_
	// となります。例えば、 `int` のゼロ値は `0` です。
	var e int
	fmt.Println(e)

	// `:=` は、変数を宣言かつ初期化する簡略記法です。
	// この例では、 `var f string = "apple"` と同じです。
	// この文法は、関数の中でのみ利用できます。
	f := "apple"
	fmt.Println(f)
}
