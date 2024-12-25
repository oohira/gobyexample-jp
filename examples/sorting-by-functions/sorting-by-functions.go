// ときには、コレクションを自然順以外でソートしたいこともあります。
// 例えば、文字列をアルファベット順ではなく、
// 長さ順でソートしたいとしましょう。
// Go のカスタムソートの例は次の通りです。

package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	fruits := []string{"peach", "banana", "kiwi"}

	// 文字列の長さに対する比較関数を実装します。
	// `cmp.Compare` を使うのが便利です。
	lenCmp := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	// このカスタムの比較関数を使って `slices.SortFunc`
	// を呼び出せば、`fruits` を名前の長さ順にソートできます。
	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	// 同じ方法を使って、組み込み型ではない値のスライスもソートできます。
	type Person struct {
		name string
		age  int
	}

	people := []Person{
		Person{name: "Jax", age: 37},
		Person{name: "TJ", age: 25},
		Person{name: "Alex", age: 72},
	}

	// `slices.SortFunc` を使って `people` を年齢でソートします。
	//
	// 注意: `Person` 構造体が大きな場合は、スライスには `*Person`
	// を含め、それに合わせたソート関数を使った方がよいかもしれません。
	// 疑わしい場合は、 [ベンチマーク](testing-and-benchmarking)
	// を取ってください！
	slices.SortFunc(people,
		func(a, b Person) int {
			return cmp.Compare(a.age, b.age)
		})
	fmt.Println(people)
}
