// Go は、`encoding.xml` パッケージで XML や XML
// ライクなフォーマットを組み込みでサポートします。

package main

import (
	"encoding/xml"
	"fmt"
)

// Plant は、XML にマッピングされます。JSON の例と同様に、
// フィールドタグがエンコーダーやデコーダーへの指示を含みます。
// ここでは、XML パッケージ特有の機能もいくつか使っています。
// `XMLName` フィールドは、この構造体を表す XML 要素の名前を
// 指示します。`id,attr` は `Id` フィールドがネストした
// 要素ではなく、XML _属性_ であることを意味します。
type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	// 植物を表す XML を出力します。`MarshalIndent`
	// を使うことで、より人間が読みやすい出力にできます。
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))

	// XML ヘッダを出力するには、明示的に加えます。
	fmt.Println(xml.Header + string(out))

	// XML のバイト列からデータ構造にパースするには `Unmarhshal`
	// を使用します。もし XML が不正だったり、Plant にマッピング
	// できなかったりする場合は、説明的なエラーが返されます。
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	// `parent>child>plant` フィールドタグは、すべての `plant`
	// を `<parent><child>...` の下へネストするようエンコーダーに
	// 指示します。
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}
