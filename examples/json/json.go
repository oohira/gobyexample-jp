// Go は、JSON のエンコードおよびデコードを組み込みで
// サポートしています。組み込み型やカスタムデータ型と
// JSON との相互変換も含みます。

package main

import "encoding/json"
import "fmt"
import "os"

// これら 2 つの構造体は、カスタムデータ型のエンコードと
// デコードをデモするために使います。
type response1 struct {
    Page   int
    Fruits []string
}
type response2 struct {
    Page   int      `json:"page"`
    Fruits []string `json:"fruits"`
}

func main() {

    // まず最初に、基本的なデータ型から JSON 文字列への
    // エンコードを確認します。アトミックな値の例は次の通りです。
    bolB, _ := json.Marshal(true)
    fmt.Println(string(bolB))

    intB, _ := json.Marshal(1)
    fmt.Println(string(intB))

    fltB, _ := json.Marshal(2.34)
    fmt.Println(string(fltB))

    strB, _ := json.Marshal("gopher")
    fmt.Println(string(strB))

    // スライスとマップの場合は次のようになります。期待通り、
    // JSON の配列とオブジェクトにエンコードされます。
    slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
    fmt.Println(string(slcB))

    mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB))

    // JSON パッケージは、カスタムデータ型も自動的に
    // エンコードできます。エンコード結果には、
    // エクスポートされたフィールドだけを含み、
    // デフォルトではそのフィールド名が JSON キーになります。
    res1D := &response1{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B))

    // エンコードされる JSON キー名をカスタマイズするには、
    // 構造体のフィールド宣言にタグを指定します。
    // `response2` の定義を確認してみてください。
    res2D := &response2{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res2B, _ := json.Marshal(res2D)
    fmt.Println(string(res2B))

    // 次に、JSON から Go の値へのデコードを見ていきます。
    // 一般的なデータ構造に対する例は、次の通りです。
    byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

    // JSON パッケージがデコードしたデータを格納できる
    // 変数を用意する必要があります。`map[string]interface{}`
    // は、文字列から任意のデータ型へのマップを保持します。
    var dat map[string]interface{}

    // 実際のデコード処理と、関連するエラーのチェックは、
    // 次のようになります。
    if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    fmt.Println(dat)

    // デコードされたマップの値を使うためには、
    // 適切な型へ変換する必要があります。例えば、
    // 次の例では `num` の値を `float64` に変換します。
    num := dat["num"].(float64)
    fmt.Println(num)

    // ネストしたデータにアクセスするためには、
    // 一連の型変換が必要になります。
    strs := dat["strs"].([]interface{})
    str1 := strs[0].(string)
    fmt.Println(str1)

    // JSON をカスタムデータ型にデコードすることもできます。
    // この方法は、プログラムを型安全にし、デコードされた
    // データにアクセスするときの型チェックを不要にできる
    // という利点があります。
    str := `{"page": 1, "fruits": ["apple", "peach"]}`
    res := response2{}
    json.Unmarshal([]byte(str), &res)
    fmt.Println(res)
    fmt.Println(res.Fruits[0])

    // これまでの例では、データを JSON 形式で標準出力へ
    // 出力するために、中間形式としてバイト型と文字列型を
    // 常に使っていました。
    // `os.Stdout` や HTTP レスポンスボディのような
    // `os.Writer`s に、エンコードした JSON
    // を直接ストリーム出力することもできます。
    enc := json.NewEncoder(os.Stdout)
    d := map[string]int{"apple": 5, "lettuce": 7}
    enc.Encode(d)
}
