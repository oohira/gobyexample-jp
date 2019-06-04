// Go でのファイル書き込みは、
// 前回見たファイル読み込みと同じパターンに従います。

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// まず初めに、文字列 (または単なるバイト列) を書き出す方法です。
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// より細かく書き込むには、ファイルをオープンします。
	f, err := os.Create("/tmp/dat2")
	check(err)

	// ファイルをオープンした直後に、遅延実行を使って
	// `Close` するのが慣例です。
	defer f.Close()

	// バイト型のスライスを期待通り `Write` できます。
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// `WriteString` 関数も使えます。
	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	// 書き込みをストレージにフラッシュするには
	// `Sync` 関数を実行します。
	f.Sync()

	// `bufio` パッケージは、前回学んだバッファ読み込みに加えて、
	// バッファ書き込みをサポートしています。
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	// バッファリングされた操作が下位の `Writer`
	// にすべて適用されることを保証するために `Flush`
	// を実行します。
	w.Flush()

}
