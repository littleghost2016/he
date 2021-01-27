package he

import (
	"math/big"
	"testing"

	"github.com/bwesterb/go-ristretto"
)

func BenchmarkXN(b *testing.B) {

	for i := 0; i < b.N; i++ {
		ttt()
	}

}

func ttt() {
	xNString := []string{
		// "1",
		"444444444444",
		"2037035976334486086268445688409378161051468393665936250636140449354381299763336706183397376",
		// "2",
		// "34234",
	}
	// 给定一组x的值，每个值为字符串类型

	xN := make([]*big.Int, len(xNString))
	rN := make([]*ristretto.Scalar, len(xNString))
	cN := make([]*ristretto.Point, len(xNString))
	// xN rN cN中每个指针对应的值是一一对应的
	// fmt.Println("xN, rN, cN", xN, rN, cN)

	H := GenerateH()
	// 生成H
	// fmt.Println("H", H)

	for i := 0; i < len(xN); i++ {
		xN[i], _ = new(big.Int).SetString(xNString[i], 10)
		// 将字符串转成大数表示的方法
		// fmt.Println(xN[i])

		rN[i] = new(ristretto.Scalar).Rand()
		// 随机产生r
		// fmt.Println("rN[i]", rN[i])

		cN[i] = CommitTo(H, rN[i], new(ristretto.Scalar).SetBigInt(xN[i]))
		// 通过CommitTo函数进行计算密文C
		// fmt.Println("xN[i]: ", xN[i], "rN[i]: ", rN[i], "cN[i]: ", cN[i])
	}

	// C1 := AddMoreThanTwo(cN)
	_ = AddMoreThanTwo(cN)
	// 密文相加（先加密后相加）

	// startTime := time.Now().UnixNano()
	// C2 := AddPrivatelyMoreThanTwo(H, rN, xN)
	_ = AddPrivatelyMoreThanTwo(H, rN, xN)
	// 参数相加（先相加后加密）

	// fmt.Println("C1: ", C1, "C2: ", C2)
	// if C2.Equals(C1) {
	// 	// 判断计算结果是否相等
	// 	fmt.Println("equal")

	// } else {
	// 	fmt.Println("not equal")
	// }

	// endTime := time.Now().UnixNano()
	// fmt.Println("第二个耗时：（纳秒）", endTime-startTime)
}
