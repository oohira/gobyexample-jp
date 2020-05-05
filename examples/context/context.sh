# サーバーをバックグラウンドで実行します。
$ go run context-in-http-servers.go &

# クライアントから `/hello` へのリクエストをシミュレートし、
# 少し後で Ctrl+C を押してキャンセルのシグナルを送ります。
$ curl localhost:8090/hello
server: hello handler started
^C
server: context canceled
server: hello handler ended
