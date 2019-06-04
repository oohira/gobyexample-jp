// _フィルタ (line filter)_ は、標準入力から入力を読み取り、
// それを処理し、そして結果を標準出力へ出力する、一般的なプログラムです。
// `grep` や `sed` は、よく知られたフィルタプログラムです。

// これは、すべての入力テキストを大文字にして書き出す、
// Go でのフィルタプログラムの例です。自分で Go
// のフィルタを書くときには、このパターンを利用できるでしょう。
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// バッファリングされない `os.Stdin`
	// をスキャナでラップすることで、便利な `Scan`
	// メソッドを使えるようになります。
	// それはスキャナを次のトークン
	// (標準のスキャナーでは次の行) まで進めます。
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// `Text` は、入力から現在のトークン、
		// ここでは次の行、を返します。
		ucl := strings.ToUpper(scanner.Text())

		// 大文字に変換した行を書き出します。
		fmt.Println(ucl)
	}

	// `Scan` 中にエラーがなかったかを確認します。
	// EOF (ファイルの末尾) が期待され、その場合は
	// `Scan` にエラーとして報告されません。
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
