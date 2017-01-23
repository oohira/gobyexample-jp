// 前回の例では、[アトミックな操作](atomic-counters) を使って、
// シンプルなカウンターの状態を管理する方法を見ました。
// より複雑な状態の場合は、複数のゴルーチンから安全にデータへ
// アクセスするために _[ミューテックス (mutex)](http://en.wikipedia.org/wiki/Mutual_exclusion)_ を使えます。

package main

import (
    "fmt"
    "math/rand"
    "sync"
    "sync/atomic"
    "time"
)

func main() {

    // 例として、マップ `state` の状態を管理します。
    var state = make(map[int]int)

    // この `mutex` は、`state` へのアクセスを同期します。
    var mutex = &sync.Mutex{}

    // 読み書き操作をした回数も記録しましょう。
    var readOps uint64 = 0
    var writeOps uint64 = 0

    // ここで、1 ミリ秒に 1 回 `state` から読み込み処理を実行する
    // ゴルーチンを 100 個開始します。
    for r := 0; r < 100; r++ {
        go func() {
            total := 0
            for {

                // 各読み込み毎に、アクセスするキーを選択し、
                // `state` へ排他的にアクセスするため `mutex` を
                // `Lock()` し、選択したキーの値を読み込み、
                // `mutex` を `Unlock()` して、最後に
                // `readOps` の数をインクリメントします。
                key := rand.Intn(5)
                mutex.Lock()
                total += state[key]
                mutex.Unlock()
                atomic.AddUint64(&readOps, 1)

                // 次の読み込みまでちょっとだけ待ちます。
                time.Sleep(time.Millisecond)
            }
        }()
    }

    // 同様に、書き込みをシミュレートするゴルーチンを
    // 10 個開始します。
    for w := 0; w < 10; w++ {
        go func() {
            for {
                key := rand.Intn(5)
                val := rand.Intn(100)
                mutex.Lock()
                state[key] = val
                mutex.Unlock()
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    // 10 個のゴルーチンを 1 秒間だけ `state` と `mutex`
    // に対して動かします。
    time.Sleep(time.Second)

    // 最終的な操作の回数を取得してレポートします。
    readOpsFinal := atomic.LoadUint64(&readOps)
    fmt.Println("readOps:", readOpsFinal)
    writeOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("writeOps:", writeOpsFinal)

    // 最後に `state` をロックして、どうなったかを確認します。
    mutex.Lock()
    fmt.Println("state:", state)
    mutex.Unlock()
}
