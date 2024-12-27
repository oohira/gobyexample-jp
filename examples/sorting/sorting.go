// Go の `slices` パッケージは、組み込み型とユーザー定義型のための
// ソートを実装しています。
// 組み込み型のためのソートを先に見ていきましょう。

package main

import (
	"fmt"
	"slices"
)

func main() {

	// ソート関数は汎用的であり、任意の _順序付けられた_ 組み込み型に
	// 対して動作します。順序付けられた型の一覧は、
	// [cmp.Ordered](https://pkg.go.dev/cmp#Ordered) を参照してください。
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs)

	// `int` をソートする例は次の通りです。
	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints:   ", ints)

	// スライスがすでにソートされているかを確認するために
	// `slices` パッケージを使うこともできます。
	s := slices.IsSorted(ints)
	fmt.Println("Sorted: ", s)
}
