// Go では、エラーを明示的な、別の戻り値として扱うのが特徴です。
// これは、Java や Ruby のような言語で使われる例外や、
// C で時々使われる結果/エラーを多重定義した単一の値とは
// 対照的です。
// Go のアプローチは、どの関数がエラーを返したかを調べやすくし、
// エラー以外のほかのタスクに使うのと同じ言語機能でエラーも
// 扱えるようにします。

package main

import "errors"
import "fmt"

// 慣例的に、エラーは戻り値の最後にし、
// 組み込みインターフェースである `error` 型をもちます。
func f1(arg int) (int, error) {
    if arg == 42 {

        // `errors.New` は、与えられたエラーメッセージをもつ
        // 基本的な `error` 値を作ります。
        return -1, errors.New("can't work with 42")

    }

    // エラーに対する `nil` 値は、エラーがなかったことを示します。
    return arg + 3, nil
}

// `Error()` メソッドを実装することで、独自の型を
// `error` として使うことができます。次の例は、
// 明示的に引数エラーを表現する独自の型を使って、
// 上のサンプルコードを書き換えたものです。
type argError struct {
    arg  int
    prob string
}

func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
    if arg == 42 {

        // この場合、新しい構造体を作るために `&argError`
        // という構文を使って、 `arg` と `prob` の 2
        // つのフィールドに値を設定しています。
        return -1, &argError{arg, "can't work with it"}
    }
    return arg + 3, nil
}

func main() {

    // 次の 2 つのループは、先のエラーを返す各関数を試します。
    // この `if` を使ったインラインのエラーチェックは、Go
    // コードの一般的なイディオムであることに注意してください。
    for _, i := range []int{7, 42} {
        if r, e := f1(i); e != nil {
            fmt.Println("f1 failed:", e)
        } else {
            fmt.Println("f1 worked:", r)
        }
    }
    for _, i := range []int{7, 42} {
        if r, e := f2(i); e != nil {
            fmt.Println("f2 failed:", e)
        } else {
            fmt.Println("f2 worked:", r)
        }
    }

    // もし独自エラーのデータをプログラムで使いたければ、
    // 型アサーション (type assertion) を使って、
    // 独自エラー型のインスタンスとしてエラーを取得する
    // 必要があるでしょう。
    _, e := f2(42)
    if ae, ok := e.(*argError); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
    }
}
