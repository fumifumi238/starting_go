package main

import (
	"fmt"
)

// 大文字はグローバル、小文字はローカル変数
type Point struct {
	x, y int
}

// Pointer.Render() でアクセスできる。
func (p *Point) Render() {
	fmt.Printf("<%d,%d>", p.x, p.y)
}

// クラスのように扱えます ポインター型にすると参照渡しになります。
func (p *Point) Set(x, y int) {
	p.x = x
	p.y = y
}

type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{Message: "エラーが発生しました", ErrCode: 1234}
}

func main() {
	type Myint int

	var n1 Myint = 5
	n2 := Myint(7)
	fmt.Println(n1)
	fmt.Println(n2)

	var point Point
	point.x = 1
	point.y = 2

	pt := Point{x: 3, y: 4}
	// x = 0が入る
	pt2 := Point{y: 28}
	fmt.Println(point.x, point.y, pt, pt2)

	type Feed struct {
		Name   string
		Amount uint
	}

	type Animal struct {
		Name string
		Feed
	}

	a := Animal{
		Name: "Monkey",
		Feed: Feed{
			Name:   "Banana",
			Amount: 10,
		},
	}

	a.Amount = 15
	fmt.Println(a)
	p := &Point{x: 10, y: 2}

	// コードの実行
	p.Render()

	p1 := Point{}
	p1.Set(4, 5)

	p2 := &Point{}
	p2.Set(8, 10)

	type User struct {
		Id   int
		Name string
	}

	m1 := map[User]string{
		{Id: 1, Name: "Taro"}: "Osaka",
	}

	m2 := map[int]User{
		1: {Id: 1, Name: "Tanaka"},
	}

	fmt.Println(m1, m2)

	err := RaiseError()
	fmt.Println(err.Error())

}
