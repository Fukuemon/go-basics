package variables

import "fmt"

func Init() {
	// var i int // varで変数を宣言（初期値は0）
	// var i int = 2 // varで変数を宣言（初期値は2）
	// var i = 2 // varで変数を宣言（型推論）
	i := 2 // :=で変数を宣言（型推論）
	ui := uint16(2)
	fmt.Println(i)
	fmt.Printf("i: %v %T\n", i, i) // %vは値、%Tは型
	fmt.Printf("i: %[1]v %[1]T\nui: %[2]v %[2]T\n", i, ui) // %[1]は1番目の引数、%[2]は2番目の引数

	// いろんな型
	f := 1.23456
	s := "Hello, World!"
	b := true
	fmt.Printf("f: %v %T\ns: %v %T\nb: %v %T\n", f, f, s, s, b, b)

	// 複数の変数を宣言s
	pi, title := 3.14159, "円周率"
	fmt.Printf("pi: %v %T\ntitle: %v %T\n", pi, pi, title, title)

	// 型変換
	x := 10 // int
	y := 1.23 // float64
	z := float64(x) + y
	fmt.Printf("z: %v %T\n", z, z)

	// 定数と一括宣言
	type Os int

	const (
		Mac Os = iota + 1 // iotaは連番を生成する
		Windows
		Linux
	)

	fmt.Printf("Mac: %v\nWindows: %v\nLinux: %v\n", Mac, Windows, Linux)



}