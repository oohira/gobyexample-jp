// Go で状態を管理するための最も重要な仕組みは、
// チャネルを使った通信です。これについては、
// [ワーカープール](worker-pools) の例で見ました。
// しかし、状態を管理する方法は他にもいくつかあります。
// ここでは、複数のゴルーチンからアクセスされる
// _アトミックカウンター (atomic counters)_ のための、
// `sync/atomic` パッケージを使う方法を見ていきます。

package main

import "fmt"
import "time"
import "sync/atomic"

func main() {

    // カウンターのために符号なし整数を使います。
    var ops uint64 = 0

    // 同時更新をシミュレートするために、1 ミリ秒に
    // 1 回カウンターをインクリメントするゴルーチンを
    // 50 個開始します。
    for i := 0; i < 50; i++ {
        go func() {
            for {
                // カウンターをアトミックにインクリメントするため、
                // `&` 構文で `ops` カウンターのメモリアドレスを
                // `AddUint64` に与えます。
                atomic.AddUint64(&ops, 1)

                // 次のインクリメントまで少しだけ待ちます。
                time.Sleep(time.Millisecond)
            }
        }()
    }

    // 加算が進むよう 1 秒間待ちます。
    time.Sleep(time.Second)

    // まだ他のゴルーチンで更新され続けているカウンターを
    // 安全に使うため、`LoadUint64` で現在の値のコピーを
    // `opsFinal` に取り出します。先の例と同様に、
    // この関数にも `&ops` で値の取得元のメモリアドレスを
    // 与える必要があります。
    opsFinal := atomic.LoadUint64(&ops)
    fmt.Println("ops:", opsFinal)
}
