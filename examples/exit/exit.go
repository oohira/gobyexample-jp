// 指定したステータスですぐさま終了するには、
// `os.Exit` を使います。

package main

import "fmt"
import "os"

func main() {

    // `os.Exit` を使う場合は、`defer` は実行 _されません_ 。
    // そのため、この `fmt.Println` は決して呼ばれません。
    defer fmt.Println("!")

    // ステータス 3 で終了します。
    os.Exit(3)
}

// 例えば C とは異なり、Go は終了ステータスを示すために、
// `main` からの整数型の戻り値を使いません。
// 非ゼロのステータスで終了したい場合は、`os.Exit`
// を使わなければなりません。
