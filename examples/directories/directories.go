// Go には、ファイルシステムの *ディレクトリ* を扱うための
// 便利な関数があります。

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// 現在の作業ディレクトリに、新しい作業ディレクトリを作ります。
	err := os.Mkdir("subdir", 0755)
	check(err)

	// 一時ディレクトリを作る場合は、`defer` でそのディレクトリを
	// 遅延削除するとよいでしょう。`os.RemoveAll` は、
	// すべてのディレクトリツリーを削除します（ `rm -rf` と同じ）。
	defer os.RemoveAll("subdir")

	// これは、空のファイルを新規に作成するヘルパー関数です。
	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	// `MkdirAll` を使うと、親ディレクトリも含めてディレクトリ階層を
	// 作成できます。コマンドラインの `mkdir -p` と同様です。
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// `ReadDir` は、ディレクトリの内容を一覧し、`os.DirEntry`
	// オブジェクトのスライスを返します。
	c, err := os.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// `Chdir` は、現在の作業ディレクトリを変更します。
	// `cd` コマンドのようなものです。
	err = os.Chdir("subdir/parent/child")
	check(err)

	// *カレント* ディレクトリの内容を一覧すると、
	// `subdir/parent/child` の結果が表示されます。
	c, err = os.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// 最初のディレクトリに戻ります。
	err = os.Chdir("../../..")
	check(err)

	// サブディレクトリも含めたすべてのディレクトリを *再帰的に*
	// 走査することもできます。`Walk` に、見つかったファイル
	// またはディレクトリを処理するコールバック関数を指定します。
	fmt.Println("Visiting subdir")
	err = filepath.Walk("subdir", visit)
}

// `visit` は、`filepath.Walk` でファイルやディレクトリが
// 再帰的に見つかるたびに実行されます。
func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}
