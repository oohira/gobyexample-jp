// Go の `math/rand` パッケージは、
// [擬似乱数](http://en.wikipedia.org/wiki/Pseudorandom_number_generator)
// の生成をサポートします。

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 例えば、`rand.Intn` は、`0 <= n < 100`
	// となる `int` 型の乱数 `n` を返します。
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	// `rand.Float64` は、`0.0 <= f < 1.0`
	// となる `float64` 型の乱数 `f` を返します。
	fmt.Println(rand.Float64())

	// 次のようにすれば、例えば `5.0 <= f < 10.0`
	// のように異なる範囲の乱数を生成することもできます。
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// デフォルトの乱数生成器は決定性があります。すなわち、
	// 生成される一連の数列は毎回同じになります。
	// 異なる数列を生成するには、異なるシードを与えます。
	// なお、これを秘密にしたい乱数のために使うのは安全でない、
	// ということに注意してください。
	// そういった場合は、`crypto/rand` を使います。
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// 返された `rand.Rand` に対して、`rand`
	// パッケージの関数と同じように関数を呼び出します。
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	// 同じ値をシードとして与えると、同じ乱数列を生成します。
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}
