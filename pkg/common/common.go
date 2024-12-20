package common

import (
	"fmt"
)

func Newline() {
	fmt.Println()
}

func Test(str string) {
	fmt.Println(str)
	Newline()
}
