// Go は <em><a href="http://en.wikipedia.org/wiki/Pointer_(computer_programming)">ポインタ (pointers)</a></em>
// をサポートするので、プログラムの中で
// 値やレコードへの参照を渡すことができます。

package main

import "fmt"

// 値と対比しながら、ポインタがどのように動作するかを
// 2 つの関数 `zeroval` と `zeroptr` を使って示します。
// `zeroval` は `int` 型のパラメーターをもつので、
// 引数は値で渡されます。 `zeroval` は、呼び出し元の
// 関数とは別の `ival` コピーを受け取ります。
func zeroval(ival int) {
    ival = 0
}

// 一方、 `zeroptr` は `*int` 型のパラメーターをもつので、
// `int` のポインタを受け取ります。関数本体の `*iptr`
// は、ポインタを _デリファレンス (dereferences)_ し、
// ポインタの指すメモリアドレスから現在の値を取得します。
// ポインタのデリファレンスに値を代入すると、
// 参照されているアドレスにある値が変わります。
func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {
    i := 1
    fmt.Println("initial:", i)

    zeroval(i)
    fmt.Println("zeroval:", i)

    // `&i` という構文で、 `i` のメモリアドレス、
    // すなわち `i` へのポインタを取得できます。
    zeroptr(&i)
    fmt.Println("zeroptr:", i)

    // ポインタは表示することもできます。
    fmt.Println("pointer:", &i)
}
