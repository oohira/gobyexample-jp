# 標準のエンコーダーと URL 互換エンコーダーで
# 微妙に異なる値になります (末尾が `+` か `-`) が、
# いずれも期待通り元の文字列にデコードされます。
$ go run base64-encoding.go
YWJjMTIzIT8kKiYoKSctPUB+
abc123!?$*&()'-=@~

YWJjMTIzIT8kKiYoKSctPUB-
abc123!?$*&()'-=@~
