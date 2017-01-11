# `zeroval` 関数は `main` の中の `i` の値を変更しませんが、
# `zeroptr` 関数は変更します。というのも、
# この変数のアドレスに対する参照をもっているからです。
$ go run pointers.go
initial: 1
zeroval: 1
zeroptr: 0
pointer: 0x42131100
