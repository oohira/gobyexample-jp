// Go は時刻と期間に対する広範なサポートを提供します。
// 以下は、いくつかの例です。

package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// 現在時刻を取得するところから始めます。
	now := time.Now()
	p(now)

	// 年、月、日などを与えて、`time` 構造体を作ることができます。
	// 時刻は常に `Location` すなわちタイムゾーンとひも付きます。
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// 時刻の様々な要素を期待通り取得することができます。
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// 月曜から日曜までの曜日を取得する `Weekday` も使えます。
	p(then.Weekday())

	// これらのメソッドは 2 つの時刻を比較し、1 つ目の時刻が
	// 2 つ目の時刻の前か、後か、あるいは同じかを判定します。
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// `Sub` メソッドは、2 つの時刻の間隔を表す
	// `Duration` を返します。
	diff := now.Sub(then)
	p(diff)

	// 期間の長さは、様々な単位に換算できます。
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// `Add` を使って、指定した期間だけ時刻を進めたり、
	// あるいは `-` を使って戻したりすることができます。
	p(then.Add(diff))
	p(then.Add(-diff))
}
