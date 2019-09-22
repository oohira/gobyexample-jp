// 前回の例では、[外部プロセスを生成する](spawning-processes)
// 例を見ました。これは、実行中の Go プロセスから外部のプロセスを利用する
// 必要がある場合に使います。しかし、ときには現在の Go プロセスを完全に別の
// (おそらく Go 以外の) プロセスに置き換えたいだけの場合もあります。
// そのためには、古典的な
// <a href="http://en.wikipedia.org/wiki/Exec_(operating_system)"><code>exec</code></a>
// 関数の Go 実装を使います。

package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {

	// 例として `ls` を実行します。
	// Go は、実行したいバイナリの絶対パスを要求するので、
	// `exec.LookPath` を使って探します (おそらく `/bin/ls`)。
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// `Exec` は、引数を (1 つの文字列ではなく) スライスで要求します。
	// いくつか一般的な引数を `ls` に渡してみることにしましょう。
	// 最初の引数は、プログラムの名前でなければならない点に注意してください。
	args := []string{"ls", "-a", "-l", "-h"}

	// また、`Exec` には使用する [環境変数](environment-variables)
	// も必要です。ここでは、現在の環境変数をそのまま渡すことにします。
	env := os.Environ()

	// これが、実際の `syscall.Exec` 呼び出しです。
	// この呼び出しが成功すると、私たちのプロセスの実行は終了し、
	// `/bin/ls -a -l -h` プロセスで置き換わります。
	// もしエラーが発生すると、戻り値を受け取ることになります。
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
