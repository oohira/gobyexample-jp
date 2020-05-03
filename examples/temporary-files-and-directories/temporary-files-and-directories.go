// プログラムの実行中に、終了後は必要がなくなるデータを
// 作りたくなることがよくあります。
// *一時ファイルと一時ディレクトリ* は、ファイルシステムを
// 汚さないので、そのために使えます。

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// 一時ファイルを作成する最も簡単な方法は、`ioutil.TempFile`
	// を実行することです。これは、ファイルを作成し、*かつ*
	// 読み書きのためにオープンします。この例では、`ioutil.TempFile`
	// が OS デフォルトの場所にファイルを作成するよう
	// 第1引数に `""` を指定しています。
	f, err := ioutil.TempFile("", "sample")
	check(err)

	// 一時ファイルの名前を表示します。Unix ベースの OS では、
	// ディレクトリは `/tmp` のようになるでしょう。
	// ファイル名は、`ioutil.TempFile` の第2引数で指定された
	// 接頭辞から始まり、残りの部分は並列呼び出しされても常に
	// 異なる名前になるよう自動的に決定されます。
	fmt.Println("Temp file name:", f.Name())

	// ファイルを後始末します。OS は一時ファイルを適当なタイミングで
	// 削除してくれることが多いですが、明示的に実行する方がよいです。
	defer os.Remove(f.Name())

	// ファイルにデータを書き出すこともできます。
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	// 多くの一時ファイルに書き出すことを想定している場合、
	// 一時 *ディレクトリ* を作った方がよいかもしれません。
	// `ioutil.TempDir` の引数は `TempFile` と同じですが、
	// ファイルを開く代わりにディレクトリの *名前* を返します。
	dname, err := ioutil.TempDir("", "sampledir")
	check(err)
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	// この一時ディレクトリを使って、一時ファイルの名前を
	// 合成することができます。
	fname := filepath.Join(dname, "file1")
	err = ioutil.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}
