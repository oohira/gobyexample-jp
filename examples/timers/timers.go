// 将来のある時点や一定間隔で繰り返し Go コードを実行したい、
// と思うことがよくあります。Go の組み込み機能である
// _タイマー (timer)_ と _ティッカー (ticker)_ は、
// これら両方のタスクを容易にします。
// まず最初にタイマーを見て、次に [ティッカー](tickers)
// を見ていきましょう。

package main

import "time"
import "fmt"

func main() {

	// タイマーは、将来の 1 回限りのイベントを表します。
	// タイマーに待ち時間を指定すると、指定した時間が
	// 経過した後で通知してくれるチャネルが提供されます。
	// このタイマーは、2 秒間待ちます。
	timer1 := time.NewTimer(2 * time.Second)

	// `<-timer1.C` は、タイマーのチャネル `C`
	// が時間になったことを知らせる値を送信するまでブロックします。
	<-timer1.C
	fmt.Println("Timer 1 expired")

	// ただ待ちたいだけであれば、`time.Sleep` が使えます。
	// タイマーが役に立つ 1 つの理由は、指定時間が経過する前に
	// キャンセルすることができる点です。これがその例です。
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
