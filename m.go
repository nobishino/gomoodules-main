package m

import (
	"fmt"

	a "github.com/nobishino/gomodules-a"
	b "github.com/nobishino/gomodules-b"
)

// Mainモジュールだからといってmainパッケージを持つ必要はないです
func MainModulesFunction() {
	fmt.Println("This is gomodules-m. I am developing this module.")
}

func Add(x, y int) int {
	return x + y
}

func init() {
	var a a.A
	_ = a.Number
	_ = a.Print
	_ = b.B{}
}
