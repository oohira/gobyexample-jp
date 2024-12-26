// _Switch ステートメント_ は、多くの分岐をもつ
// 条件文を表現します。

package main

import (
	"fmt"
	"time"
)

func main() {

	// これは基本的な `switch` の例です。
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// 同じ `case` の中で複数の式をカンマで区切って指定できます。
	// また、`default` ケースも使えます。
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// 式のない `switch` は、if/else の代替手段になります。
	// この例は、 `case` が定数式に限らないことも示します。
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// 型 `switch` は、値の代わりに型を比較します。これは、
	// インターフェースの値に対する型を特定するのに使えます。
	// この例では、変数 `t` は case に対応する型をもちます。
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
