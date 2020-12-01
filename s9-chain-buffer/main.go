package main

import "fmt"

//最简单的
func main() {

	c := make(chan int, 10)

	c <- 1
	fmt.Println("len=", len(c))
	c <- 2
	fmt.Println("len=", len(c))
	c <- 3
	fmt.Println("len=", len(c))

	close(c)

	//c <- 4
	//fmt.Println("len=", len(c))

	v, ok := <-c
	fmt.Println(v, ok)

	for i := range c {
		fmt.Println(i)
	}

	fmt.Println("out")
}

//又一个小例子
//func print(c chan int) {
//	for i := range c {
//		fmt.Println(i)
//		//time.Sleep(1 * time.Second)
//	}
//}
//func main() {
//	c := make(chan int, 10)
//	go print(c)
//
//	for i := 0; i < 100; i++ {
//		c <- i
//	}
//}

//func worker(done chan bool) {
//	time.Sleep(time.Second)
//	// 通知任务已完成
//	done <- true
//}
//func main() {
//	done := make(chan bool, 1)
//	go worker(done)
//	// 等待任务完成
//	<-done
//}
