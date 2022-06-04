package main

import "fmt"

func Hello() {
	fmt.Println("Hello")
}

func loop(x int) int {
	if x > 10 {
		return 1
	}
	return x
}
func integers() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	fmt.Println("Hello World")
	Hello()

	// x := 2

	var (
		y, z int
		name string
	)

	const (
		// import できます
		GLOBAL = 100
		// := ではない

		// import できません
		local = 1
	)

	fmt.Println(GLOBAL, local)
	// var b bool
	// var x int8 => -128 <= x <= 127
	// var x int16 => -32768 <= x <= 32767
	// var x int32 => -2147483648 <= x <= 214743647
	// var x int64 => -9223372036854775808 <= x <= 9223372036854775807 int

	// var x unit8(byte) 0 <= x <= 255
	// var x unit16 0<= x <= 65535
	// var x unit32 0<= x <= 42944967295
	// var x unit64 0 <= x <= 18446744073709551615 unit
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{}
	c := [...]int{1, 2, 3}
	z = 2
	y = 3
	name = "AAAA"
	// 無名関数
	f := func(x, y int) int {
		return x + y
	}
	f(2, 3)
	fmt.Println(z, y, a, b, c, name)

	ints := integers()
	fmt.Println(ints()) // 1
	fmt.Println(ints()) // 2
	fmt.Println(ints()) // 3

	otherInts := integers()
	fmt.Println(otherInts()) // 1

	// go run hello.go

	// 実行ファイル(fmtなどを直接作るファイル。重い)

	// go build -o hello hello.go
	// ./hello

	fmt.Println(loop(5))

	breakInt := 0
	for {
		breakInt++
		if breakInt > 5 {
			break
		}
	}
	fmt.Println(breakInt)

	whileInt := 0

	// while(i < 10)と同じ
	for whileInt < 10 {
		whileInt++
	}

	fmt.Println(whileInt)

	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Println(i)
	}

	fruits := [3]string{"Banana", "Orange", "Apple"}

	for i, s := range fruits {
		// i 配列の番号　s 要素
		// fruits[0] = "Banana"
		fmt.Printf("fruits[%d]=%s\n", i, s)
	}
}
