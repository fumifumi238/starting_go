package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	// x := 2

	var (
		x, y int
		name string
	)

	// var b bool
	// var x int8 => -128 <= x <= 127
	// var x int16 => -32768 <= x <= 32767
	x = 2
	y = 3
	name = "AAAA"
	fmt.Println(x, y, name)
}

// go run hello.go

// 実行ファイル(fmtなどを直接作るファイル。重い)

// go build -o hello hello.go
// ./hello
