package main

import "fmt"

type I interface {
	M()
}

type T struct {
	Name string
	Year int
}

func (t *T) M() {
	fmt.Printf("Name: %v Year: %v\n ", t.Name, t.Year)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {

	var i I = &T{Name: "Mickel", Year: 12}
	var s I = &T{Name: "Kate", Year: 15}
	// panic: runtime error: invalid memory address or nil pointer dereference
	// var u I t.Namet.Yearがないためにエラーとなります。
	// u.M()

	i.M()
	s.M()

	do(21)
	do("hello")
	do(true)

}
