package main

import (
	"fmt"
	"golang/pkg/basic_concepts"
	"golang/pkg/common"
)

func main() {
	basic_concepts.Introduction()
	common.Newline()
	basic_concepts.BasicSyntax()
	common.Test("Test 2: Who are you?")
	common.Test("Test 3: I am a programmer")
	common.Test("Test 4: I am a golang programmer")
	fmt.Println("Finished")
}
