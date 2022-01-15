// ときには、私たちの Go プログラムは他の Go
// でないプロセスを生成する必要があります。

package main

import (
	"fmt"
	"io"
	"os/exec"
)

func main() {

	// まずは引数なし、つまり入力を受け取らずに、
	// 標準出力へ何か表示するだけの単純なコマンドから始めましょう。
	// `exec.Command` ヘルパーは、
	// この外部プロセスを表すオブジェクトを作成します。
	dateCmd := exec.Command("date")

	// `Output` は、もう 1 つのヘルパーです。
	// コマンドを実行して終了を待ち、出力を取得する
	// という一般的なケースを扱えます。エラーがなければ、
	// `dateOut` は日付情報のバイト列を保持します。
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// 次に、もう少し複雑なケースとして、
	// パイプを使って外部プロセスの `stdin` にデータを与え、
	// 結果を `stdout` から取得する例を見てみましょう。
	grepCmd := exec.Command("grep", "hello")

	// 次の例では、標準入力/出力のパイプを明示的につかみ、
	// プロセスを開始します。入力データを書き込み、出力を読み込み、
	// そして最後にプロセスが終了するのを待ちます。
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()

	// 上の例ではエラーチェックを省略しましたが、いつもの
	// `if err != nil` パターンを使うことができます。
	// また、上の例では `StdoutPipe` の結果だけを取得しましたが、
	// 全く同じ方法で `StderrPipe` から取得することもできます。
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// コマンドを起動するときは、1 つのコマンドライン文字列を
	// 渡せるのではなく、コマンドと引数の配列を明示的に
	// 渡す必要がある点に注意してください。
	// 文字列でコマンドを起動したいなら、
	// `bash` の `-c` オプションを使うことができます。
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
