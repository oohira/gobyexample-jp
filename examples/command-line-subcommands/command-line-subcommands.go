// `go` や `git` のようなコマンドラインツールには多くの
// *サブコマンド* があり、それぞれが固有のフラグをもっています。
// 例えば、`go build` と `go get` は、`go` ツールの
// サブコマンドのうちの 2 つです。
// `flag` パッケージを使うと、固有のフラグをもつシンプルな
// サブコマンドを簡単に定義できます。

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// `NewFlagSet` 関数を使ってサブコマンドを宣言し、
	// 続けてこのサブコマンド固有のフラグを定義します。
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	// 異なるサブコマンドに対して、異なるフラグを定義できます。
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	// サブコマンドは、プログラムの最初の引数となる想定です。
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	// どのサブコマンドが呼ばれたか確認します。
	switch os.Args[1] {

	// 各サブコマンドに対して、そのフラグをパースし、
	// 残りの引数にアクセスします。
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}
