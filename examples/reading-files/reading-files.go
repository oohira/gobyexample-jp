// ファイルの読み書きは、多くの Go プログラムで必要になる、
// 基本的なタスクです。
// 最初に、ファイル読み込みの例をいくつか見ていきましょう。

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// ファイル読み込みは、ほとんどの関数呼び出しでエラーチェックが必要です。
// このヘルパーは、以降のエラーチェック処理を簡易化します。
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// おそらく最も基本的なファイル読み込みのタスクは、
	// ファイルの中身をすべてメモリへ読み込むことです。
	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	// ファイルのどの部分をどのように読み込むかを
	// もっとコントロールしたいと思うこともよくあるでしょう。
	// その場合は、`os.File` 値を取得するためにファイルを
	// `Open` することから始めます。
	f, err := os.Open("/tmp/dat")
	check(err)

	// ファイルの先頭から何バイトか読み込みます。
	// 5 バイトまでの読み込みを許しますが、
	// 実際に何バイト読まれたかも記録します。
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// ファイルの特定の位置を `Seek` して、
	// そこから `Read` することもできます。
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	// `io` パッケージは、ファイル読み込みに便利な関数を提供しています。
	// 例えば、前述の例のような読み込みは、`ReadAtLeast`
	// を使ってより頑強に実装することができます。
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// 組み込みのファイル巻き戻し機能はありませんが、
	// `Seek(0, 0)` で実現できます。
	_, err = f.Seek(0, 0)
	check(err)

	// `bufio` パッケージは、バッファ読み込みを実装しており、
	// 少しずつ何回も読み込む場合の効率と、
	// 追加でサポートしている読み込みメソッドの両方で有用です。
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// 処理が完了したらファイルをクローズします
	// (普通は、`defer` を使って `Open` の直後に予約します)。
	f.Close()
}
