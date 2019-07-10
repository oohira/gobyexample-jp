$ go build command-line-subcommands.go 

# まず、foo サブコマンドを実行します。
$ ./command-line-subcommands foo -enable -name=joe a1 a2
subcommand 'foo'
  enable: true
  name: joe
  tail: [a1 a2]

# 次に、bar を試します。
$ ./command-line-subcommands bar -level 8 a1
subcommand 'bar'
  level: 8
  tail: [a1]

# ただし、bar は foo のフラグは受け付けません。
$ ./command-line-subcommands bar -enable a1
flag provided but not defined: -enable
Usage of bar:
  -level int
    	level

# 次は、プログラムをパラメーター化する別の一般的な方法である、
# 環境変数を見ていきます。
