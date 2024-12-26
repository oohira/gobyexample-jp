// _範囲 (range)_ は、様々な組み込みのデータ構造を反復処理します。
// これまでに学んだデータ構造で `range` をどのように使うかを
// 見ていきましょう。

package main

import "fmt"

func main() {

	// ここでは、スライス内の数字を合計するために
	// `range` を使っています。配列でも同様に使えます。
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// 配列やスライスに対する `range` は、各要素に対する
	// インデックスと値の両方を提供します。先の例では、
	// インデックスが必要なかったのでブランク識別子 `_`
	// で無視しました。
	// しかし、実際にインデックスを使いたいこともあります。
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// マップに対する `range` は、キーと値のペアを反復処理します。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// `range` はマップのキーだけの反復処理もできます。
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// 文字列に対する `range` は、Unicode コードポイントを反復処理します。
	// 1 つ目の値は文字列の先頭からこの `文字 (rune)` までのバイト数を表し、
	// 2 つ目は `rune` そのものを表します。
	// 詳細は、 [Strings and Runes](strings-and-runes) を参照してください。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
