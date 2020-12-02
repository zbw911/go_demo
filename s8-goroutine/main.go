package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 100; i++ {
		//如果没有这行，打印会出问题
		//j := i
		//go func() {
		//	fmt.Println(j)
		//}()

		//go func(x int) {
		//	fmt.Println(x)
		//}(i)

		go func() {
			time.Sleep(10 * time.Millisecond)
			fmt.Println(i)
		}()

		//这个没有问题，值传递，
		//go myRoutine(strconv.Itoa(i))
		//go myIntRoutine(i)
	}
	fmt.Println("after   loop")
	//time.Sleep(3 * time.Second)

}

func myIntRoutine(input int) {

	fmt.Println(input)
	time.Sleep(1 * time.Microsecond)
}

func myRoutine(input string) {
	fmt.Println(input)
}
