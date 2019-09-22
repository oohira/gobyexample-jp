// Go の `sort` パッケージは、組み込み型とユーザー定義型のための
// ソートを実装しています。
// 組み込み型のためのソートを先に見ていきましょう。

package main

import (
	"fmt"
	"sort"
)

func main() {

	// ソートメソッドは、組み込み型ごとに特有です。
	// 例えば、文字列に対する例は次の通りです。
	// ソートはインプレースで行われる、すなわち、
	// 与えられたスライスを直接変更し、新しくは返さない点に
	// 注意してください。
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// `int` をソートする例は次の通りです。
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	// スライスがすでにソートされているかを確認するために
	// `sort` パッケージを使うこともできます。
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
