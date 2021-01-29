package he

import (
	"math/big"
	"testing"

	"github.com/bwesterb/go-ristretto"
)

func BenchmarkCommitTo512Bit(b *testing.B) {

	// 获取随机生成的指定长度的x
	x := GetRandomBigInt(512)

	H := GenerateH()

	// 随机产生r
	r := new(ristretto.Scalar).Rand()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = CommitTo(H, r, new(ristretto.Scalar).SetBigInt(x))
	}

}

func BenchmarkAddMoreThanTwo512Bit(b *testing.B) {

	number := 2

	xN := make([]*big.Int, number)
	rN := make([]*ristretto.Scalar, number)
	cN := make([]*ristretto.Point, number)

	H := GenerateH()

	for i := 0; i < number; i++ {
		xN[i] = GetRandomBigInt(512)

		rN[i] = new(ristretto.Scalar).Rand()

		cN[i] = CommitTo(H, rN[i], new(ristretto.Scalar).SetBigInt(xN[i]))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = AddMoreThanTwo(cN)
	}

}

func BenchmarkAddPrivatelyMoreThanTwo512Bit(b *testing.B) {
	number := 2

	xN := make([]*big.Int, number)
	rN := make([]*ristretto.Scalar, number)
	// cN := make([]*ristretto.Point, number)

	H := GenerateH()

	for i := 0; i < number; i++ {
		xN[i] = GetRandomBigInt(512)

		rN[i] = new(ristretto.Scalar).Rand()

		// cN[i] = CommitTo(H, rN[i], new(ristretto.Scalar).SetBigInt(xN[i]))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = AddPrivatelyMoreThanTwo(H, rN, xN)
	}
}

func BenchmarkCheckIfEqual512Bit(b *testing.B) {

	// 表示需要做几次加法
	number := 2

	xN := make([]*big.Int, number)
	rN := make([]*ristretto.Scalar, number)
	cN := make([]*ristretto.Point, number)

	H := GenerateH()

	for i := 0; i < number; i++ {
		// 将字符串转成大数表示的方法
		xN[i] = GetRandomBigInt(512)
		// 以下一行仅供测试显示xN[i]的值
		// fmt.Println(xN[i])

		// 随机产生r
		rN[i] = new(ristretto.Scalar).Rand()
		// 以下一行仅供测试显示rN[i]的值
		// fmt.Println("rN[i]", rN[i])

		// 通过CommitTo函数进行计算密文C
		cN[i] = CommitTo(H, rN[i], new(ristretto.Scalar).SetBigInt(xN[i]))
		// 以下一行仅供测试显示xN[i], rN[i], cN[i]的值
		// b.Log("xN[i]: ", xN[i], "rN[i]: ", rN[i], "cN[i]: ", cN[i])
	}

	C1 := AddMoreThanTwo(cN)
	// 密文相加（先加密后相加）

	C2 := AddPrivatelyMoreThanTwo(H, rN, xN)
	// 参数相加（先相加后加密）

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if !C2.Equals(C1) {
			b.Error("Not Equal!")
		}
	}
}

func BenchmarkCommitTo1024Bit(b *testing.B) {

	// 获取随机生成的指定长度的x
	x := GetRandomBigInt(1024)

	H := GenerateH()

	// 随机产生r
	r := new(ristretto.Scalar).Rand()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = CommitTo(H, r, new(ristretto.Scalar).SetBigInt(x))
	}

}

func BenchmarkAddMoreThanTwo1024Bit(b *testing.B) {

	number := 2

	xN := make([]*big.Int, number)
	rN := make([]*ristretto.Scalar, number)
	cN := make([]*ristretto.Point, number)

	H := GenerateH()

	for i := 0; i < number; i++ {
		xN[i] = GetRandomBigInt(1024)

		rN[i] = new(ristretto.Scalar).Rand()

		cN[i] = CommitTo(H, rN[i], new(ristretto.Scalar).SetBigInt(xN[i]))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = AddMoreThanTwo(cN)
	}

}

func BenchmarkAddPrivatelyMoreThanTwo1024Bit(b *testing.B) {
	number := 2

	xN := make([]*big.Int, number)
	rN := make([]*ristretto.Scalar, number)
	// cN := make([]*ristretto.Point, number)

	H := GenerateH()

	for i := 0; i < number; i++ {
		xN[i] = GetRandomBigInt(1024)

		rN[i] = new(ristretto.Scalar).Rand()

		// cN[i] = CommitTo(H, rN[i], new(ristretto.Scalar).SetBigInt(xN[i]))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = AddPrivatelyMoreThanTwo(H, rN, xN)
	}
}

func BenchmarkCheckIfEqual1024Bit(b *testing.B) {

	// 表示需要做几次加法
	number := 2

	xN := make([]*big.Int, number)
	rN := make([]*ristretto.Scalar, number)
	cN := make([]*ristretto.Point, number)

	H := GenerateH()

	for i := 0; i < number; i++ {
		// 将字符串转成大数表示的方法
		xN[i] = GetRandomBigInt(1024)
		// 以下一行仅供测试显示xN[i]的值
		// fmt.Println(xN[i])

		// 随机产生r
		rN[i] = new(ristretto.Scalar).Rand()
		// 以下一行仅供测试显示rN[i]的值
		// fmt.Println("rN[i]", rN[i])

		// 通过CommitTo函数进行计算密文C
		cN[i] = CommitTo(H, rN[i], new(ristretto.Scalar).SetBigInt(xN[i]))
		// 以下一行仅供测试显示xN[i], rN[i], cN[i]的值
		// b.Log("xN[i]: ", xN[i], "rN[i]: ", rN[i], "cN[i]: ", cN[i])
	}

	C1 := AddMoreThanTwo(cN)
	// 密文相加（先加密后相加）

	C2 := AddPrivatelyMoreThanTwo(H, rN, xN)
	// 参数相加（先相加后加密）

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if !C2.Equals(C1) {
			b.Error("Not Equal!")
		}
	}
}

func BenchmarkCommitTo2048Bit(b *testing.B) {

	// 获取随机生成的指定长度的x
	x := GetRandomBigInt(2048)

	H := GenerateH()

	// 随机产生r
	r := new(ristretto.Scalar).Rand()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = CommitTo(H, r, new(ristretto.Scalar).SetBigInt(x))
	}

}

func BenchmarkAddMoreThanTwo2048Bit(b *testing.B) {

	number := 2

	xN := make([]*big.Int, number)
	rN := make([]*ristretto.Scalar, number)
	cN := make([]*ristretto.Point, number)

	H := GenerateH()

	for i := 0; i < number; i++ {
		xN[i] = GetRandomBigInt(2048)

		rN[i] = new(ristretto.Scalar).Rand()

		cN[i] = CommitTo(H, rN[i], new(ristretto.Scalar).SetBigInt(xN[i]))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = AddMoreThanTwo(cN)
	}

}

func BenchmarkAddPrivatelyMoreThanTwo2048Bit(b *testing.B) {
	number := 2

	xN := make([]*big.Int, number)
	rN := make([]*ristretto.Scalar, number)
	// cN := make([]*ristretto.Point, number)

	H := GenerateH()

	for i := 0; i < number; i++ {
		xN[i] = GetRandomBigInt(2048)

		rN[i] = new(ristretto.Scalar).Rand()

		// cN[i] = CommitTo(H, rN[i], new(ristretto.Scalar).SetBigInt(xN[i]))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = AddPrivatelyMoreThanTwo(H, rN, xN)
	}
}

func BenchmarkCheckIfEqual2048Bit(b *testing.B) {

	// 表示需要做几次加法
	number := 2

	xN := make([]*big.Int, number)
	rN := make([]*ristretto.Scalar, number)
	cN := make([]*ristretto.Point, number)

	H := GenerateH()

	for i := 0; i < number; i++ {
		// 将字符串转成大数表示的方法
		xN[i] = GetRandomBigInt(2048)
		// 以下一行仅供测试显示xN[i]的值
		// fmt.Println(xN[i])

		// 随机产生r
		rN[i] = new(ristretto.Scalar).Rand()
		// 以下一行仅供测试显示rN[i]的值
		// fmt.Println("rN[i]", rN[i])

		// 通过CommitTo函数进行计算密文C
		cN[i] = CommitTo(H, rN[i], new(ristretto.Scalar).SetBigInt(xN[i]))
		// 以下一行仅供测试显示xN[i], rN[i], cN[i]的值
		// b.Log("xN[i]: ", xN[i], "rN[i]: ", rN[i], "cN[i]: ", cN[i])
	}

	C1 := AddMoreThanTwo(cN)
	// 密文相加（先加密后相加）

	C2 := AddPrivatelyMoreThanTwo(H, rN, xN)
	// 参数相加（先相加后加密）

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if !C2.Equals(C1) {
			b.Error("Not Equal!")
		}
	}
}

func BenchmarkCommitTo4096Bit(b *testing.B) {

	// 获取随机生成的指定长度的x
	x := GetRandomBigInt(4096)

	H := GenerateH()

	// 随机产生r
	r := new(ristretto.Scalar).Rand()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = CommitTo(H, r, new(ristretto.Scalar).SetBigInt(x))
	}

}

func BenchmarkAddMoreThanTwo4096Bit(b *testing.B) {

	number := 2

	xN := make([]*big.Int, number)
	rN := make([]*ristretto.Scalar, number)
	cN := make([]*ristretto.Point, number)

	H := GenerateH()

	for i := 0; i < number; i++ {
		xN[i] = GetRandomBigInt(4096)

		rN[i] = new(ristretto.Scalar).Rand()

		cN[i] = CommitTo(H, rN[i], new(ristretto.Scalar).SetBigInt(xN[i]))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = AddMoreThanTwo(cN)
	}

}

func BenchmarkAddPrivatelyMoreThanTwo4096Bit(b *testing.B) {
	number := 2

	xN := make([]*big.Int, number)
	rN := make([]*ristretto.Scalar, number)
	// cN := make([]*ristretto.Point, number)

	H := GenerateH()

	for i := 0; i < number; i++ {
		xN[i] = GetRandomBigInt(4096)

		rN[i] = new(ristretto.Scalar).Rand()

		// cN[i] = CommitTo(H, rN[i], new(ristretto.Scalar).SetBigInt(xN[i]))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = AddPrivatelyMoreThanTwo(H, rN, xN)
	}
}

func BenchmarkCheckIfEqual4096Bit(b *testing.B) {

	// 表示需要做几次加法
	number := 2

	xN := make([]*big.Int, number)
	rN := make([]*ristretto.Scalar, number)
	cN := make([]*ristretto.Point, number)

	H := GenerateH()

	for i := 0; i < number; i++ {
		// 将字符串转成大数表示的方法
		xN[i] = GetRandomBigInt(4096)
		// 以下一行仅供测试显示xN[i]的值
		// fmt.Println(xN[i])

		// 随机产生r
		rN[i] = new(ristretto.Scalar).Rand()
		// 以下一行仅供测试显示rN[i]的值
		// fmt.Println("rN[i]", rN[i])

		// 通过CommitTo函数进行计算密文C
		cN[i] = CommitTo(H, rN[i], new(ristretto.Scalar).SetBigInt(xN[i]))
		// 以下一行仅供测试显示xN[i], rN[i], cN[i]的值
		// b.Log("xN[i]: ", xN[i], "rN[i]: ", rN[i], "cN[i]: ", cN[i])
	}

	C1 := AddMoreThanTwo(cN)
	// 密文相加（先加密后相加）

	C2 := AddPrivatelyMoreThanTwo(H, rN, xN)
	// 参数相加（先相加后加密）

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if !C2.Equals(C1) {
			b.Error("Not Equal!")
		}
	}
}
