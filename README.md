# bunnysay in Golang
Bunny Sign for terminals with full width support (rewritten in Golang)

```
｜￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣｜
｜　　Ｍｕｌｔｉ－ｌｉｎｅｄ　　　｜
｜　ｓｕｐｐｏｒｔ　ｗｉｔｈ　　　｜
｜　　ｌｏｏｋ　ａｈｅａｄ　　　　｜
｜ｓｐａｃｉｎｇ！　　Ａｌｓｏ　　｜
｜　　　ｗｉｔｈ　ｓｏｍｅ　　　　｜
｜　　　ｋｉｃｋ－ａｓｓ　　　　　｜
｜　　　ｃｅｎｔｅｒｉｎｇ．　　　｜
｜＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿｜
(\__/) ||
(•ㅅ•) ||
/ 　 づ

```

## Summary

Takes input from standard input, you can use

`echo '<string>' | bunnysay`

if you wish to use from terminal.

## Build
Requires go and you can simply run the following...


```
go get github.com/0x4445565a/bunnysay
# If you get complaints about golang.org/x/text/width missing run
# go get golang.org/x/text/width
go install github.com/0x4445565a/bunnysay
echo "Hack the Planet" | bunnysay
```