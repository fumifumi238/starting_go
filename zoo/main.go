package main

import (
	"fmt"

	// ./animalsは非推奨
	// go mod init fooと指定 ファイルの絶対パスを指定　foo/bar/baz
	"starting_go/zoo/animals"
)

func main() {
	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.MonkeyFeed())
	fmt.Println(animals.RabbitFeed())

	// go run zoo/main.go zoo/app.go または　zoo/*.goと指定
	fmt.Println(AppName())

}
