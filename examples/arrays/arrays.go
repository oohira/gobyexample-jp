// Go では、 _配列 (array)_ は特定の長さをもつ番号付けられた要素の列です。
// 典型的な Go のコードでは [スライス](slices) の方がより一般的ですが、
// 配列はいくつかの特殊なシナリオで役に立ちます。

package main

import "fmt"

func main() {

	// ここでは、ちょうど 5 つの `int` をもつ配列
	// `a` を作成しています。要素の型と長さはいずれも
	// 配列の型の一部です。デフォルトでは、配列は
	// ゼロ値で初期化されるので、 `int` の場合は
	// `0` になります。
	var a [5]int
	fmt.Println("emp:", a)

	// `array[index] = value` で指定位置に値を設定し、
	// `array[index]` によって値を取得できます。
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// `len` ビルトイン関数は配列の長さを返します。
	fmt.Println("len:", len(a))

	// 配列を 1 行で宣言かつ初期化するには、
	// 次の記法を使います。
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// `...` を使ってコンパイラに要素数を数えさせることもできます。
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// `:` を使ってインデックスを指定すると、間にある要素は
	// 0 で初期化されます。
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	// 配列型は 1 次元ですが、多次元のデータ構造を
	// 表す型も構成できます。
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// 多次元配列を一度に初期化もできます。
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}
