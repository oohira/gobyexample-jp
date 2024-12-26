// _マップ (Maps)_ は、Go における [連想配列](https://en.wikipedia.org/wiki/Associative_array)
// のための組み込み型です。別の言語では、_ハッシュ (hashes)_
// や _辞書 (dicts)_ と呼ばれることもあります。

package main

import (
	"fmt"
	"maps"
)

func main() {

	// 空のマップを作成するには、 `make` ビルトイン関数を使って
	// `make(map[キーの型]値の型)` とします。
	m := make(map[string]int)

	// 典型的な `name[key] = val` という記法を使って、
	// キーと値のペアを設定します。
	m["k1"] = 7
	m["k2"] = 13

	// 例えば、 `fmt.Println` などでマップを表示すると、
	// すべてのキーと値のペアが出力されるでしょう。
	fmt.Println("map:", m)

	// `name[key]` でキーに対する値を取得できます。
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// キーが存在しない場合、その値に対する
	// [ゼロ値](https://go.dev/ref/spec#The_zero_value)
	// が返されます。
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	// `len` ビルトイン関数は、マップに対しては
	// キーと値のペアの数を返します。
	fmt.Println("len:", len(m))

	// `delete` ビルトイン関数は、マップから
	// キーと値のペアを削除します。
	delete(m, "k2")
	fmt.Println("map:", m)

	// マップから *すべての* キーと値のペアを削除するには、
	// `clear` ビルトイン関数を使います。
	clear(m)
	fmt.Println("map:", m)

	// マップから値を取得するときの戻り値にはオプションで
	// 2 番目があり、指定したキーがマップに存在したか
	// どうかを表します。これは、キーが存在しなかったのか、
	// `0` や `""` などのゼロ値で存在したのかを区別するのに
	// 使えます。この例では、値自体は必要なかったので、
	// _ブランク識別子 (blank identifier)_ と呼ばれる `_`
	// で無視しています。
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// 新しいマップの宣言と初期化を次のように
	// 同じ行で書くこともできます。
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// `maps` パッケージにはマップ操作に便利な関数があります。
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
