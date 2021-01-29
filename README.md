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
BenchmarkCommitTo512Bit-8                          12831             93365 ns/op             472 B/op          7 allocs/op
BenchmarkAddMoreThanTwo512Bit-8                  4311496               277 ns/op             160 B/op          1 allocs/op
BenchmarkCheckIfEqual512Bit-8                    7316947               162 ns/op               0 B/op          0 allocs/op
BenchmarkAddMoreThanTwo1024Bit-8                 4336843               276 ns/op             160 B/op          1 allocs/op
BenchmarkCheckIfEqual1024Bit-8                   7405692               162 ns/op               0 B/op          0 allocs/op
BenchmarkCommitTo2048Bit-8                         12657             94524 ns/op             888 B/op          7 allocs/op
BenchmarkAddMoreThanTwo2048Bit-8                 4357240               277 ns/op             160 B/op          1 allocs/op
BenchmarkAddPrivatelyMoreThanTwo2048Bit-8          12577             95212 ns/op             952 B/op          8 allocs/op
BenchmarkCheckIfEqual2048Bit-8                   7405906               162 ns/op               0 B/op          0 allocs/op
BenchmarkCommitTo4096Bit-8                         12420             95913 ns/op            1432 B/op          7 allocs/op
BenchmarkAddMoreThanTwo4096Bit-8                 4353709               277 ns/op             160 B/op          1 allocs/op
BenchmarkAddPrivatelyMoreThanTwo4096Bit-8          12352             96680 ns/op            1496 B/op          8 allocs/op
BenchmarkCheckIfEqual4096Bit-8                   7409154               162 ns/op               0 B/op          0 allocs/op
PASS
ok      he      28.740s
```

