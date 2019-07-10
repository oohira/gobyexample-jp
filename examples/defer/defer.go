// _Defer_ は、通常クリーンアップなどのために、
// 関数呼び出しが後で実行されることを保証するために使います。
// `defer` は、他のプログラミング言語で `ensure` や `finally`
// が使われる場面でよく使われます。

package main

import "fmt"
import "os"

// ファイルを作成して、何か書き込み、最後にクローズしたいとします。
// `defer` を使って実現する方法は次の通りです。
func main() {

	// `createFile` 関数でファイルオブジェクトを取得した後ですぐに、
	// `closeFile` 関数によるファイルクローズを遅延実行します。
	// これは、`writeFile` 関数が完了した後で、
	// 呼び出し元の関数 (`main`) の最後で実行されます。
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")

}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	// ファイルをクローズする場合には、遅延実行する関数の中であっても
	// エラーを確認することが重要です。
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
