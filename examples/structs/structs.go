// Go の _構造体 (structs)_ は、フィールドの集まりで
// 定義される型です。構造体は、データをグループ化して
// レコードを作るのに役立ちます。

package main

import "fmt"

// この `person` 構造体型は、 `name` フィールドと
// `age` フィールドをもちます。
type person struct {
	name string
	age  int
}

// `newPerson` は、指定した名前で新しい person 構造体を作ります。
func newPerson(name string) *person {
	// Go はガベージコレクション機能をもつ言語です。
	// ローカル変数へのポインタでも関数から安全に返すことができます。
	// 変数へのアクティブな参照がなくなったときのみ、
	// ガベージコレクタによって解放されます。
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	// この構文は新しい構造体を作ります。
	fmt.Println(person{"Bob", 20})

	// 構造体を初期化するときに、
	// フィールド名を指定することもできます。
	fmt.Println(person{name: "Alice", age: 30})

	// 省略されたフィールドはゼロ値になります。
	fmt.Println(person{name: "Fred"})

	// `&` を頭につけると構造体へのポインタになります。
	fmt.Println(&person{name: "Ann", age: 40})

	// 構造体の生成をコンストラクタ関数でカプセル化する慣用記法です。
	fmt.Println(newPerson("Jon"))

	// ドットを使ってフィールドにアクセスします。
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// 構造体のポインタにもドットが使えます。
	// この場合、ポインタは自動的にデリファレンスされます。
	sp := &s
	fmt.Println(sp.age)

	// 構造体は変更可能 (mutable) です。
	sp.age = 51
	fmt.Println(sp.age)

	// 構造体型が単一の値だけで使用される場合、名前を付ける必要はありません。
	// その場合、値は無名の構造体型をもつことになります。この手法は、
	// [テーブル駆動テスト](testing-and-benchmarking) でよく使用されます。
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
