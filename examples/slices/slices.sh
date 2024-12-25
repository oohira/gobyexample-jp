# スライスは配列とは異なる型ですが、 `fmt.Println`
# では同じように表現されることに注意してください。
$ go run slices.go
uninit: [] true true
emp: [  ] len: 3 cap: 3
set: [a b c]
get: c
len: 3
apd: [a b c d e f]
cpy: [a b c d e f]
sl1: [c d e]
sl2: [a b c d e]
sl3: [c d e f]
dcl: [g h i]
t == t2
2d:  [[0] [1 2] [2 3 4]]

# Go におけるスライスの設計や実装に関する詳細は、
# Go チームによるこの [素晴らしいブログ記事](https://go.dev/blog/slices-intro)
# を参考にしてください。

# これで配列とスライスを学んだので、次は Go のもう
# 1 つの主要な組み込み型であるマップを見ていきましょう。
