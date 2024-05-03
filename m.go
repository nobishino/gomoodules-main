package m

import "fmt"

// Mainモジュールだからといってmainパッケージを持つ必要はないです
func MainModulesFunction() {
	fmt.Println("This is gomodules-m. I am developing this module.")
}

func Add(x, y int) int {
	return x + y
}
