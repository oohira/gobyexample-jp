// 前回の例では、[アトミックな操作](atomic-counters) を使って、
// シンプルなカウンターの状態を管理する方法を見ました。
// より複雑な状態の場合は、複数のゴルーチンから安全にデータへ
// アクセスするために [_ミューテックス (mutex)_](http://en.wikipedia.org/wiki/Mutual_exclusion) を使えます。

package main

import (
	"fmt"
	"sync"
)

// Container は、カウンタのマップを保持します。
// 複数のゴルーチンから並行して更新したいので、
// アクセスを同期化するために `Mutex` を追加しています。
// ミューテックスはコピーできないので、この `struct`
// を渡すときはポインタで渡す必要があります。
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	// `counters` へアクセスする前にミューテックスをロックし、
	// [defer](defer) 文を使って関数の最後でアンロックします。
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	c := Container{
		// ミューテックスはゼロ値がそのまま使えるので、
		// ここでは初期化が不要である点に注意してください。
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	// この関数はループを使って名前付きのカウンタをインクリメントします。
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	// 複数のゴルーチンを並行実行します。これらはすべて同じ
	// `Container` にアクセスし、そのうちの 2 つは
	// カウンタも同じであることに注意してください。
	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	// ゴルーチンが完了するのを待ちます。
	wg.Wait()
	fmt.Println(c.counters)
}
