# HE

同态加密模块。

> 本项目参考自[threehook](https://github.com/threehook)/**[go-pedersen-commitment](https://github.com/threehook/go-pedersen-commitment)**

# 使用方法

1. 克隆本仓库
2. 配置go环境并在仓库目录运行以下命令（我使用的是powershell）

```powershell
go test -bench="." -benchmem
```

运行结果

```
goos: windows
goarch: amd64
pkg: he
BenchmarkCommitTo512Bit-8                          12793             93135 ns/op             472 B/op          7 allocs/op
BenchmarkAddMoreThanTwo512Bit-8                  4366563               277 ns/op             160 B/op          1 allocs/op
BenchmarkAddPrivatelyMoreThanTwo512Bit-8           12754             93723 ns/op             536 B/op          8 allocs/op
BenchmarkCheckIfEqual512Bit-8                    7405642               162 ns/op               0 B/op          0 allocs/op
BenchmarkCommitTo1024Bit-8                         12830             93568 ns/op             600 B/op          7 allocs/op
BenchmarkAddMoreThanTwo1024Bit-8                 4331306               277 ns/op             160 B/op          1 allocs/op
BenchmarkAddPrivatelyMoreThanTwo1024Bit-8          12748             94012 ns/op             664 B/op          8 allocs/op
BenchmarkCheckIfEqual1024Bit-8                   7389553               162 ns/op               0 B/op          0 allocs/op
BenchmarkCommitTo2048Bit-8                         12670             94243 ns/op             888 B/op          7 allocs/op
BenchmarkAddMoreThanTwo2048Bit-8                 4343328               277 ns/op             160 B/op          1 allocs/op
BenchmarkAddPrivatelyMoreThanTwo2048Bit-8          12631             94859 ns/op             952 B/op          8 allocs/op
BenchmarkCheckIfEqual2048Bit-8                   7336680               162 ns/op               0 B/op          0 allocs/op
BenchmarkCommitTo4096Bit-8                         12513             95945 ns/op            1432 B/op          7 allocs/op
BenchmarkAddMoreThanTwo4096Bit-8                 4307062               278 ns/op             160 B/op          1 allocs/op
BenchmarkAddPrivatelyMoreThanTwo4096Bit-8          12418             97049 ns/op            1496 B/op          8 allocs/op
BenchmarkCheckIfEqual4096Bit-8                   7416769               162 ns/op               0 B/op          0 allocs/op
PASS
ok      he      28.706s
```

