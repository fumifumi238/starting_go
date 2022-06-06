package main

import (
	"fmt"
)

// 任意の数を受け取れるsum
func sum(s ...int) int {
	n := 0

	// sum(1,2,3) = [1,2,3]の配列のループ _は番号
	for _, v := range s {
		n += v
	}

	return n
}
func main() {
	s := make([]int, 10)
	fmt.Println(s)
	s[0] = 1
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	s = make([]int, 5, 10)

	fmt.Println(s)
	s[0] = 1
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	// [5]のように指定しない
	a := []int{1, 2, 3, 4, 5}
	b := a[0:2]
	fmt.Println(b)
	a = append(a, 6, 7, 8)
	fmt.Println(a)

	s0 := []int{1, 2, 3, 4}
	s1 := []int{5, 6, 7, 8, 9}
	s2 := append(s0, s1...)
	// s0にある要素をs1に上書きします nはコピーされた要素数をしめします
	n := copy(s1, s0)
	fmt.Println(s2)
	fmt.Println(n, s1)

	c := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c1 := c[2:4]
	// cap = 元の配列-最小値　10 - 2 = 8
	fmt.Println(len(c1), cap(c1))

	c2 := c[2:4:5]

	// cap = 5 - 2 = 3  high <= max <= cap(a)
	fmt.Println(len(c2), cap(c2))
	fmt.Println(sum())
	fmt.Println(sum(5, 6, 7, 8, 10))
	fmt.Println(sum(s0...))

	m := make(map[int]string)

	m[1] = "japan"
	m[5] = "America"
	m[10] = "China"

	fmt.Println(m, m[1])

	r := map[int]string{
		1: "taro",
		2: "jiro",
		3: "saburo",
	}

	taro, ok := r[1]
	fmt.Println(taro, ok)

	// 存在しない要素だとokがfalseになります。
	sirou, ok := r[4]
	fmt.Println(sirou, ok)
	for k, v := range r {
		fmt.Printf("%d=>%s\n", k, v)
	}

	delete(r, 3)

	for k, v := range r {
		fmt.Printf("%d=>%s\n", k, v)
	}

}
