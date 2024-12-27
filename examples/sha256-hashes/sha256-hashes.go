// [_SHA256 ハッシュ_](https://en.wikipedia.org/wiki/SHA-2) は、
// バイナリやテキストのかたまりに対して短い ID を計算するのによく使われます。
// 例えば、TLS/SSL 証明書は、SHA256 を使用して証明書の署名を計算します。
// 以下は、Go で SHA256 ハッシュを計算する方法です。

package main

// Go は、`crypto/*` パッケージで
// 複数のハッシュ関数を実装しています。
import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "sha256 this string"

	// ハッシュの生成から始めます。
	h := sha256.New()

	// `Write` はバイト列を受け取ります。文字列 `s` がある場合は、
	// `[]byte(s)` を使って強制的にバイト列にします。
	h.Write([]byte(s))

	// バイト型のスライスとして最終的なハッシュ値を取得します。
	// `Sum` への引数は、既存のバイト型スライスへハッシュ値を
	// 追記したい場合に使えますが、通常は必要ありません。
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
