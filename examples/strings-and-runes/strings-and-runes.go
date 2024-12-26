// Go の文字列は、バイトの読み取り専用スライスです。言語と標準ライブラリは、
// 文字列を [UTF-8](https://en.wikipedia.org/wiki/UTF-8)
// でエンコードされたテキストのコンテナとして特別に扱います。
// 他の言語では、文字列は "文字 (characters)" の集合とみなされますが、
// Go には、`rune` という概念があります。`rune` は、Unicode
// コードポイントを表す整数です。
// [この Go ブログ記事](https://go.dev/blog/strings) は、
// このトピックの良い入門資料です。

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// `s` は、タイ語で "こんにちは" を表すリテラル値をもつ
	// `string` 型です。Go の文字列リテラルは、UTF-8
	// でエンコードされています。
	const s = "สวัสดี"

	// 文字列は `[]byte` と等しいため、これは内部に格納されている
	// 生のバイト列の長さを出力します。
	fmt.Println("Len:", len(s))

	// 文字列へのインデックス操作は、各インデックスでの生のバイト値を生成します。
	// このループは、`s` のコードポイントを構成するバイトの16進値を出力します。
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// 文字列内の _rune_ を数えるには、`utf8` パッケージを使用します。
	// `RuneCountInString` は、UTF-8 の各 rune を順にデコード
	// しなければならないため、実行時間が文字列のサイズに依存するので
	// 注意してください。タイ語の文字は、UTF-8 のコードポイントが複数
	// バイトにまたがる場合もあるため、結果が予想外に感じるかもしれません。
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// `range` ループは、文字列を特別に扱い、各 `rune`
	// を文字列内でのオフセットとともにデコードします。
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// `utf8.DecodeRuneInString` 関数を使用すれば、
	// 同じ反復処理を実現できます。
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		// `rune` 値を関数に渡す例です。
		examineRune(runeValue)
	}
}

func examineRune(r rune) {

	// シングルクォートで囲まれた値は、_rune リテラル_ です。
	// `rune` 値を rune リテラルと直接比較できます。
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
