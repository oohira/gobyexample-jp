// Go は、バージョン 1.23 から
// [イテレーター (iterator)](https://go.dev/blog/range-functions)
// をサポートし、ほぼすべてのものが `range` で扱えるようになりました！

package main

import (
	"fmt"
	"iter"
	"slices"
)

// [前回の例](generics) で見た `List` 型を再び取り上げましょう。
// その例では、リスト内のすべての要素をスライスとして返す `AllElements`
// メソッドがありました。Go のイテレーターを使えば、もっとうまく実現できます。
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

// All は、_イテレーター (iterator)_ を返します。Go では、
// イテレーターは [特別なシグネチャ](https://pkg.go.dev/iter#Seq)
// をもつ関数です。
func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		// イテレーター関数は、引数として別の関数を受け取ります。
		// 通常、この関数は `yield` と呼ばれますが、名前は任意です。
		// イテレーターは、繰り返したい各要素に対して `yield` を呼び出し、
		// `yield` の戻り値で早期終了を判定します。
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

// イテレーターは、必ずしも対象のデータ構造を必要とせず、
// 有限である必要もありません！以下は、フィボナッチ数列を生成する
// イテレーターを返す関数です。これは、`yield` が `true`
// を返す限り終わることはありません。
func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1

		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func main() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)

	// `List.All` がイテレーターを返すため、通常の `range`
	// ループで使用できます。
	for e := range lst.All() {
		fmt.Println(e)
	}

	// [slices](https://pkg.go.dev/slices) パッケージなどには、
	// イテレーターを操作するための便利な関数が多数含まれています。
	// 例えば、`Collect` は任意のイテレーターを受け取り、
	// その値をすべてスライスに収集します。
	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	for n := range genFib() {

		// ループが `break` または早期リターンに達すると、
		// イテレーターに渡された `yield` 関数は `false` を返します。
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}
