// Go の `if` と `else` による条件分岐は直感的です。

package main

import "fmt"

func main() {

	// これは基本的な例です。
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// `else` のない `if` ステートメントも可能です。
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// `&&` や `||` のような論理演算子は条件式で役立ちます。
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	// 条件式の前にステートメントを書くこともできます。
	// ここで宣言された変数は、すべての分岐の中で
	// 使うことができます。
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

// Go では、条件式の前後に括弧 (parentheses)
// は必要ありませんが、中括弧 (braces)
// は必要である、という点に注意してください。
