// `//go:embed` は、
// [コンパイラディレクティブ](https://pkg.go.dev/cmd/compile#hdr-Compiler_Directives)
// で、ビルド時に任意のファイルやフォルダを Go バイナリに埋め込めます。
// `embed` ディレクティブの詳細は、[こちら](https://pkg.go.dev/embed)
// を参照してください。
package main

// `embed` パッケージをインポートします。もしこのパッケージから
// エクスポートされた識別子を使わない場合は、`_ "embed"`
// を使ったブランクインポートが可能です。
import (
	"embed"
)

// `embed` ディレクティブは、Go ソースファイルを含むディレクトリからの
// 相対パスを受け入れます。このディレクティブは、直後に続く `string`
// 型の変数に、ファイルの内容を埋め込みます。
//
//go:embed folder/single_file.txt
var fileString string

// ファイルの内容を `[]byte` 型に埋め込むこともできます。
//
//go:embed folder/single_file.txt
var fileByte []byte

// ワイルドカードを使って複数のファイルやフォルダを埋め込むこともできます。
// この場合、[embed.FS 型](https://pkg.go.dev/embed#FS) 変数を使います。
// `embed.FS` は、シンプルな仮想ファイルシステムを実装しています。
//
//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {

	// `single_file.txt` の内容を出力します。
	print(fileString)
	print(string(fileByte))

	// 埋め込まれたフォルダからいくつかのファイルを取得します。
	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}
