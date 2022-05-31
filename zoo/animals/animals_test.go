package animals

import (
	"testing"
)

// go test -v starting_go/zoo/animals modの絶対パスを指定
func TestElephantFeed(t *testing.T) {
	expect := "Grass"
	actual := ElephantFeed()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}

func TestMonkeyFeed(t *testing.T) {
	expect := "Banana"
	actual := MonkeyFeed()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}

func TestRabbitFeed(t *testing.T) {
	expect := "Carrot"
	actual := RabbitFeed()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}
