// <em>[レートリミット (Rate limiting)](http://en.wikipedia.org/wiki/Rate_limiting)</em>
// は、リソースの使用量を制御し、サービスの品質を維持するための、
// とても重要な仕組みです。
// Go は、ゴルーチンとチャネル、[ティッカー](tickers) により、
// レートリミットをエレガントにサポートします。

package main

import (
	"fmt"
	"time"
)

func main() {

	// まず最初に、基本的なレートリミットを見ていきます。
	// リクエストに対する処理を制限したいとしましょう。
	// これらのリクエストに同じ名前の `requests`
	// チャネルで対応します。
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// `limiter` チャネルは、200 ミリ秒ごとに値を受信します。
	// これは、レートリミットの仕組みの中でレギュレーターの
	// 役割を果たします。
	limiter := time.Tick(200 * time.Millisecond)

	// 各リクエストを処理する前に `limiter` チャネルからの
	// 受信をブロックさせることで、200 ミリ秒に
	// 1 リクエストしか処理できないよう制限しています。
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// あるいは、全体としてはレートリミットを担保しつつ、
	// 一時的なバーストリクエストは許容したいと思うかもしれません。
	// その場合、`limiter` チャネルをバッファリングすれば
	// 実現できます。ここで、`burstyLimiter` チャネルは、
	// 3 リクエストまでならバーストを許します。
	burstyLimiter := make(chan time.Time, 3)

	// 許容されているバースト量を表すため、
	// チャネルに値を満たしておきます。
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// 200 ミリ秒ごとに、`burstyLimiter` の上限である
	// 3 つまで、新しい値を追加します。
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// それでは、5 リクエストきた場合をシミュレートします。
	// これらのうち最初の 3 リクエストは、`burstyLimiter`
	// のバースト機能の恩恵を受けるはずです。
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
