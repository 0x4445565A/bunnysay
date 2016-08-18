# bunnysay in Golang
Bunny Sign for terminals with wchar support (rewritten in Golang)

```
 _____________________
｜                    ｜
｜Multi-lined support ｜
｜  with look ahead   ｜
｜spacing!  Also with ｜
｜   some kick ass    ｜
｜     centering.     ｜
｜____________________｜
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
git clone git@github.com:0x4445565A/bunnysay.git
cd bunnysay
go build
# And if you so desire
go install
echo "Hack the Planet" | bunnysay
```