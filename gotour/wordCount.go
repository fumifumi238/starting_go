package main

import (
	"fmt"
	"math"
	"strings"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
func WordCount(s string) map[string]int {

	list := map[string]int{}
	m := strings.Split(s, " ")
	for _, value := range m {
		_, ok := list[value]
		if !ok {
			list[value] = 1
		} else {
			list[value] += 1
		}
	}
	return list
}

func main() {
	fmt.Println(WordCount("I go to school to meet you"))
	fmt.Println(WordCount("I ate a donut. Then I ate another donut."))

	v := Vertex{3, 4}
	// ポインター参照なので値自体が変わります。
	v.Scale(10)
	fmt.Println(v.Abs())

}
