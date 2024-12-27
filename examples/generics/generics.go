// Go 1.18 以降で、_ジェネリクス (generic)_ がサポートされました。
// ジェネリクスは、_型パラメーター (type parameter)_ としても知られます。

package main

import "fmt"

// ジェネリック関数の例を見ていきます。`SlicesIndex` は、
// 任意の `comparable` 型のスライス s と、その型の要素 v
// を受け取り、s 内で最初に v が出現する位置を返します。
// 存在しない場合は、-1 を返します。`comparable` 制約は、
// この型の値を `==` および `!=` 演算子で比較できることを意味します。
// この型シグネチャの詳細な説明は、
// [このブログ記事](https://go.dev/blog/deconstructing-type-parameters)
// を参照してください。なお、この関数は、標準ライブラリに
// [slices.Index](https://pkg.go.dev/slices#Index) として存在します。
func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

// ジェネリック型の例として、`List` は任意の型の値をもつ単方向リストです。
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// 通常の型と同様に、ジェネリック型にもメソッドを定義できますが、
// 型パラメーターを保持する必要があります。
// 型は `List[T]` であり、`List` ではありません。
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// AllElements は、リスト内のすべての要素をスライスで返します。
// カスタム型の要素を反復処理するより慣用的な方法は、次回紹介します。
func (lst *List[T]) AllElements() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	var s = []string{"foo", "bar", "zoo"}

	// ジェネリック関数を呼び出すとき、_型推論_ に依存できます。
	// `SlicesIndex` を呼び出すのに、`S` と `E` の型を
	// 指定する必要はありません。コンパイラが自動的に推論します。
	fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

	// 型を明示的に指定もできます。
	_ = SlicesIndex[[]string, string](s, "zoo")

	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.AllElements())
}
