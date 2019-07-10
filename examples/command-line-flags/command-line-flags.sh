# コマンドラインフラグをもつプログラムを試すためには、
# 先にコンパイルして、出力されたバイナリを直接実行するのがよいです。
$ go build command-line-flags.go

# まずはビルドされたプログラムにすべてのフラグの値を指定して
# みましょう。
$ ./command-line-flags -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []

# フラグを省略すると、
# 自動的にフラグごとのデフォルト値となる点に注意してください。
$ ./command-line-flags -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []

# 位置引数は、任意のフラグの後に続けて指定することができます。
$ ./command-line-flags -word=opt a1 a2 a3
word: opt
...
tail: [a1 a2 a3]

# `flag` パッケージは、すべてのフラグが位置引数の前に
# 出現することを要求する点に注意してください。
# (さもないと、フラグも位置引数として解釈されてしまいます)
$ ./command-line-flags -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3 -numb=7]

# コマンドラインプログラムに対して自動生成された
# ヘルプテキストを取得するには、`-h` または `--help`
# フラグを使います。
$ ./command-line-flags -h
Usage of ./command-line-flags:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -word="foo": a string

# `flag` パッケージに指定されていないフラグを与えた場合は、
# プログラムはエラーメッセージを表示して、ヘルプテキストを
# 再度表示します。
$ ./command-line-flags -wat
flag provided but not defined: -wat
Usage of ./command-line-flags:
...
