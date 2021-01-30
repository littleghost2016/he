package he

import (
	"crypto/rand"
	"math/big"

	"github.com/bwesterb/go-ristretto"
)

// The prime order of the base point is 2^252 + 27742317777372353535851937790883648493.
var n25519, _ = new(big.Int).SetString("7237005577332262213973186563042994240857116359379907606001950938285454250989", 10)

// CommitTo 承诺函数
// C = rG + xH
// Commit to a value x
// H - Random secondary point on the curve
// r - Private key used as blinding factor
// x - The value (number of tokens)
// 其中C为密文，x为给定的值
func CommitTo(H *ristretto.Point, r, x *ristretto.Scalar) *ristretto.Point {

	// startTime := time.Now().UnixNano()

	result := new(ristretto.Point)
	rPoint := new(ristretto.Point)

	transferPoint := new(ristretto.Point)
	// result, rPoint, transferPoint := new(ristretto.Point)
	rPoint.ScalarMultBase(r)
	transferPoint.ScalarMult(H, x)
	result.Add(rPoint, transferPoint)

	// endTime := time.Now().UnixNano()

	// fmt.Println(startTime, endTime)
	// fmt.Println("CommitTo的耗时（纳秒）:", endTime-startTime)

	return result
}

// GenerateH Generate a random point on the curve
func GenerateH() *ristretto.Point {
	random := new(ristretto.Scalar).Rand()
	H := new(ristretto.Point)
	H.ScalarMultBase(random)
	return H
}

// // Subtract two commitments using homomorphic encryption
// func Add(cX, cY *ristretto.Point) ristretto.Point {
// 	var addPoint ristretto.Point
// 	addPoint.Add(cX, cY)
// 	return addPoint
// }

// // Subtract two known values with blinding factors
// //   and compute the committed value
// //   add rX - rY (blinding factor private keys)
// //   add vX - vY (hidden values)
// func AddPrivately(H *ristretto.Point, rX, rY *ristretto.Scalar, vX, vY *big.Int) ristretto.Point {
// 	var rDif ristretto.Scalar
// 	var vDif big.Int
// 	rDif.Add(rX, rY)
// 	vDif.Add(vX, vY)
// 	vDif.Mod(&vDif, n25519)

// 	var vScalar ristretto.Scalar
// 	var rPoint ristretto.Point
// 	vScalar.SetBigInt(&vDif)

// 	rPoint.ScalarMultBase(&rDif)
// 	var vPoint, result ristretto.Point
// 	vPoint.ScalarMult(H, &vScalar)
// 	result.Add(&rPoint, &vPoint)
// 	return result
// }

// AddMoreThanTwo 多于两个相加
func AddMoreThanTwo(cN []*ristretto.Point) *ristretto.Point {

	addPoint := new(ristretto.Point).Add(cN[0], cN[1])
	for i := 2; i < len(cN); i++ {
		addPoint.Add(addPoint, cN[i])
	}
	// fmt.Println(addPoint)
	return addPoint
}

// AddPrivatelyMoreThanTwo 多于两个私钥相加
func AddPrivatelyMoreThanTwo(H *ristretto.Point, rN []*ristretto.Scalar, xN []*big.Int) *ristretto.Point {
	rDif := new(ristretto.Scalar).Add(rN[0], rN[1])
	for i := 2; i < len(rN); i++ {
		rDif.Add(rDif, rN[i])
	}

	xDif := new(big.Int).Add(xN[0], xN[1])
	for i := 2; i < len(xN); i++ {
		xDif.Add(xDif, xN[i])
	}
	xDif.Mod(xDif, n25519)

	xScalar := new(ristretto.Scalar)
	rPoint := new(ristretto.Point)
	xScalar.SetBigInt(xDif)

	rPoint.ScalarMultBase(rDif)
	xPoint := new(ristretto.Point)
	result := new(ristretto.Point)
	xPoint.ScalarMult(H, xScalar)
	result.Add(rPoint, xPoint)

	return result
}

// GetRandomBigInt 生成指定长度的随机切片
func GetRandomBigInt(sliceLength uint) (serialNumber *big.Int) {

	max := new(big.Int).Lsh(big.NewInt(1), sliceLength)
	// 以下一行仅供测试显示max值
	// fmt.Println(max)

	// 生成大整数随机数
	serialNumber, _ = rand.Int(rand.Reader, max)
	// 以下一行仅供测试显示serialNumber值
	// fmt.Println(serialNumber)

	// 以下两行为第一版生成随机数的过程，返回[]byte
	// randomBytesSlice = make([]byte, sliceLength)
	// rand.Read(randomBytesSlice)
	return
}
