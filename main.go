package main

import (
	"crypto/elliptic"
	"fmt"
	"math/big"
)

type ECPoint struct {
	Curve elliptic.Curve
	X     *big.Int
	Y     *big.Int
}

func BasePointGGet() (point ECPoint) {
	return ECPointGen(point.X, point.Y)
} // G-generator receiving

func ECPointGen(x, y *big.Int) (point ECPoint) {
	return ECPoint{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}
} // ECPoint creation with pre-defined parameters

func IsOnCurveCheck(a ECPoint) (c bool) {
	return a.Curve.IsOnCurve(a.X, a.Y)
} // P ∈ CURVE?

func AddECPoints(a, b ECPoint) (c ECPoint) {
	p := ECPoint{Curve: a.Curve}
	p.X, p.Y = a.Curve.Add(a.X, a.Y, b.X, b.Y)
	return p
} // P + Q

func DoubleECPoints(a ECPoint) (c ECPoint) {
	p := ECPoint{Curve: a.Curve}
	p.X, p.Y = a.Curve.Double(a.X, a.Y)
	return p
} // 2P

func ScalarMult(a ECPoint, k *big.Int) (c ECPoint) {
	n := new(big.Int).Mod(k, a.Curve.Params().N)
	p := new(ECPoint)
	p.Curve = a.Curve
	p.X, p.Y = a.Curve.ScalarMult(a.X, a.Y, n.Bytes())
	return
} // k * P

func PrintECPoint(point ECPoint) {
	fmt.Printf("X: %v", point.X)
	fmt.Printf("Y: %v", point.Y)
} //Print point

func ECPointToString(point ECPoint) (s string) {
	x := fmt.Sprintf("%v", point.X)
	y := fmt.Sprintf("%v", point.Y)
	return x + "/" + y

} //Convert point to string
func main() {

}
