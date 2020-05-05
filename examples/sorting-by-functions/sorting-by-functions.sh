# プログラムを実行すると、期待通り文字列の長さ順にソートされた
# リストが表示されます。
$ go run sorting-by-functions.go
[kiwi peach banana]

# ユーザー定義型を作成して 3 つの `Interface` メソッドを実装し、
# `sort.Sort` 関数を呼び出す、という同じパターンに従うことで、
# Go のスライスを任意の関数でソートができます。
