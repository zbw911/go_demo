///关于异常处理的 Demo
package main

import (
	"fmt"
)

func g() {
	defer func() {
		fmt.Println("b")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("d")
	}()

	f()

	fmt.Println("e")
}

func f() {
	fmt.Println("a")
	panic("c.panic exception here ............")
	fmt.Println("c")
}

func main() {
	g()
	fmt.Println("e")
}
