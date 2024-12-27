# サンプルを実行するには、これらのコマンドを使ってください。
# Go Playground の制限で、このサンプルはローカルマシンでしか
# 実行できないことに注意してください。
$ mkdir -p folder
$ echo "hello go" > folder/single_file.txt
$ echo "123" > folder/file1.hash
$ echo "456" > folder/file2.hash

$ go run embed-directive.go
hello go
hello go
123
456

