package calculator

import "fmt"

var offset float64 = 1 // private変数
var Offset float64 = 1 // public変数

func Sum(a, b float64) float64 {
	fmt.Println("multiply: ",multiply(a, b)) // パッケージ間でのアクセスは可能
	fmt.Println("Multiply: ",multiply(a, b))
	return a + b + offset
}

