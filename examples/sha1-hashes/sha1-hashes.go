// [_SHA1 ハッシュ_](http://en.wikipedia.org/wiki/SHA-1)
// は、バイナリやテキストのかたまりに対して短い ID
// を計算するのによく使われます。例えば、
// [Git バージョン管理システム](http://git-scm.com/) は
// バージョン管理されるファイルやディレクトリを識別するために、
// SHA1 を広範囲にわたって使用しています。
// 以下は、Go で SHA1 ハッシュを計算する方法です。

package main

// Go は様々な `crypto/*` パッケージで、
// 複数のハッシュ関数を実装しています。
import (
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "sha1 this string"

	// ハッシュを生成するパターンは、`sha1.New()` の次に
	// `sha1.Write(bytes)`、そして `sha1.Sum([]byte{})`
	// です。以下では、ハッシュの生成から始めます。
	h := sha1.New()

	// `Write` はバイト列を受け取ります。文字列 `s` がある場合は、
	// `[]byte(s)` を使って強制的にバイト列にします。
	h.Write([]byte(s))

	// バイト型のスライスとして最終的なハッシュ値を取得します。
	// `Sum` への引数は、既存のバイト型スライスへハッシュ値を
	// 追記したい場合に使えますが、通常は必要ありません。
	bs := h.Sum(nil)

	// 例えば Git のコミットなど、SHA1 値はよく 16
	// 進数で表示されます。ハッシュ値を 16 進文字列に変換するには、
	// `%x` フォーマット指定子を使ってください。
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
