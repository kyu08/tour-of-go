package main

import (
	"fmt"
	"math"
)

// これを定義する必要があるのはなぜか
// float64 でも文法的にはできるけど float64 に Sqrt 用のエラーを定義するのはもちろん不適切なのでこう書くべき、ということ？
type ErrNegativeFloat float64

func (e ErrNegativeFloat) Error() string {
	return fmt.Sprintf("negative float is not valid: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		// なんでこれでエラーを呼べるのか -> String(1) のようにキャストしている？
		return 0, ErrNegativeFloat(x)
	}

	z0, z := 1.0, x
	for math.Abs(z-z0)/z > 1e-15 {
		z0, z = z, z-(z*z-x)/(2*z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
