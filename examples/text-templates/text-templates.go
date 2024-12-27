// Go は、動的なコンテンツを作成したり、カスタマイズされた出力をユーザーに表示したり
// する機能を、`text/template` パッケージで組み込みサポートしています。
// これと対になる `html/template` パッケージでは、同じ API で追加のセキュリティ
// 機能を備えているので、HTML を生成する場合は使用してください。

package main

import (
	"os"
	"text/template"
)

func main() {

	// 新しいテンプレートを作成し、文字列からその内容をパースします。
	// テンプレートは、静的なテキストと `{{...}}` で囲まれた
	// "アクション (actions)" の組み合わせで、アクション部分に
	// 動的なコンテンツが挿入されます。
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	// また、`template.Must` 関数を使えば、`Parse` がエラーを返した
	// 場合に panic を起こせます。テンプレートをグローバルスコープで
	// 初期化する場合に便利です。
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	// テンプレートを "実行 (executing)" することで、アクションに特定の
	// 値を挿入したテキストを生成します。`{{.}}` アクションは、`Execute`
	// に渡された値で置き換えられます。
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	// ヘルパー関数を定義します。
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	// データが構造体の場合、`{{.FieldName}}` アクションでフィールドに
	// アクセスできます。フィールドは、テンプレート実行時にアクセスできるよう
	// エクスポートされている必要があります。
	t2 := Create("t2", "Name: {{.Name}}\n")

	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	// マップでも同様です。マップの場合は、キー名が大文字かどうかに
	// 制限はありません。
	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})

	// if/else はテンプレートで条件付き実行を提供します。値が、
	// 0 や 空文字列、nil ポインタなど型のデフォルト値であれば、
	// `false`と見なされます。
	// また、テンプレートの別の機能として、アクション内で `-`
	// を使用して空白を削除できます。
	t3 := Create("t3",
		"{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	// range ブロックを使用すると、スライス、配列、マップ、そして
	// チャネルを反復処理できます。range ブロック内では、`{{.}}`
	// が現在の要素を指します。
	t4 := Create("t4",
		"Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout,
		[]string{
			"Go",
			"Rust",
			"C++",
			"C#",
		})
}
