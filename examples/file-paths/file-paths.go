// `filepath` パッケージは、オペレーティングシステム間で移植性がある方法で
// *ファイルパス* をパースして構築する機能を提供します。例えば、
// Linux では `dir/file`, Windows では `dir\file` といった具合です。
package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	// 移植性のある方法でパスを構築するには、`Join` を使用すべきです。
	// これは任意の数の引数を受け取り、階層をもったパスを作ります。
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	// `/` や `\` を手動で連結する代わりに、常に `Join` を使用すべきです。
	// 移植性に加えて、`Join` は不要なセパレーターやディレクトリ変更を
	// 取り除き、パスを正規化してくれます。
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// `Dir` や `Base` は、パスをディレクトリとファイルへ
	// 分割するのに使えます。`Split` は、1回の呼び出しで
	// その両方を返します。
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	// パスが絶対パスかどうかを判定することもできます。
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	// ファイル名によっては、ドットに続く拡張子をもちます。
	// `Ext` を使うとファイル名から拡張子を取り出せます。
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// 拡張子を取り除いたファイル名を取得するには、`strings.TrimSuffix`
	// を使います。
	fmt.Println(strings.TrimSuffix(filename, ext))

	// `Rel` は、*base* から *target* への相対パスを調べます。
	// もしも base から target に相対パスを作れない場合は、error
	// を返します。
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
