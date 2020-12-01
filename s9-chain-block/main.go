package main

import "fmt"

//最简单的
func main() {

	c := make(chan int)

	defer close(c)
	go func() {
		c <- 1
	}()

	v, ok := <-c
	fmt.Println(v, ok)
}

//又一个小例子
//func compute(i int, c chan int) {
//	i = i + 1
//	c <- i
//}
//func main() {
//	c := make(chan int)
//	go compute(1, c)
//	go compute(2, c)
//	x, y := <-c, <-c
//	fmt.Println(x, y, x*y)
//}

// range 读取
//func main() {
//
//	c := make(chan int)
//	go func() {
//		for i := 0; i < 10; i = i + 1 {
//			fmt.Println("input :", i)
//			c <- i
//		}
//		close(c)
//	}()
//	for i := range c {
//		fmt.Println("output:", i)
//	}
//	fmt.Println("Finished")
//}

//开关信号
//func main() {
//
//	c := make(chan int)
//
//	defer close(c)
//	go func() {
//		fmt.Println("	休眠")
//		time.Sleep(1 * time.Second)
//		fmt.Println("	完成休眠")
//		c <- 1
//	}()
//	fmt.Println("开始")
//	<-c
//	//fmt.Println(v, ok)
//	fmt.Println("退出")
//}
