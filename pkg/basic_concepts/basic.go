package basic_concepts

import (
	"fmt"
	"golang/pkg/common"
)

func Introduction() {
	fmt.Println("Introduction to Go Programming")
	common.Newline()
	fmt.Println("What is Go?\n")
	common.Newline()
	fmt.Println("Go is an open-source programming language designed for simplicity, reliability, and efficiency. \nIt is a statically typed, compiled language that runs on a wide range of operating systems and platforms, including Linux, Windows, macOS, and more.\n\nGo was created at Google in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson, and it has since become one of the most popular programming languages in the world.\nGo is often used for building web applications, cloud services, system tools, and more.\n")
	fmt.Println("Installing Go \nFollow the instructions on the official Go website (https://golang.org/doc/install) to install Go on your system. Setting up Go workspace. \nSet up your Go workspace by creating a directory structure and setting the `GOPATH` environment variable. \nThe `GOPATH` environment variable specifies the location of your Go workspace.\n")
	common.Newline()
}

func BasicSyntax() {
	fmt.Println("Variables and Constants")
	var name string = "John"
	const pi float64 = 3.14
	fmt.Println("variable name string is " + name + " | constant variable pi float is " + fmt.Sprintf("%f", pi) + "")
}
