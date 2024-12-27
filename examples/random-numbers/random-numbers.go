// Go の `math/rand/v2` パッケージは、
// [擬似乱数](https://en.wikipedia.org/wiki/Pseudorandom_number_generator)
// の生成をサポートします。

package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	// 例えば、`rand.IntN` は、`0 <= n < 100`
	// となる `int` 型の乱数 `n` を返します。
	fmt.Print(rand.IntN(100), ",")
	fmt.Print(rand.IntN(100))
	fmt.Println()

	// `rand.Float64` は、`0.0 <= f < 1.0`
	// となる `float64` 型の乱数 `f` を返します。
	fmt.Println(rand.Float64())

	// 次のようにすれば、例えば `5.0 <= f < 10.0`
	// のように異なる範囲の乱数も生成できます。
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// 既知のシードを使いたい場合は、 `rand.Source`
	// を新規に作成し、 `New` コンストラクタに渡します。
	// `NewPCG` は、2 つの `uint64` 型の数値をシードとして
	// [PCG](https://en.wikipedia.org/wiki/Permuted_congruential_generator)
	// ソースを作成できます。
	s2 := rand.NewPCG(42, 1024)
	r2 := rand.New(s2)
	fmt.Print(r2.IntN(100), ",")
	fmt.Print(r2.IntN(100))
	fmt.Println()

	s3 := rand.NewPCG(42, 1024)
	r3 := rand.New(s3)
	fmt.Print(r3.IntN(100), ",")
	fmt.Print(r3.IntN(100))
	fmt.Println()
}
