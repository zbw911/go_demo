package main

//退出 使用wait grop
//func main() {
//
//	l := 10
//	var group sync.WaitGroup
//	group.Add(l)
//
//	for i := 0; i < l; i++ {
//		//如果没有这行，打印会出问题
//		j := i
//		go func() {
//			time.Sleep(1 * time.Second)
//			fmt.Println(j)
//			group.Done()
//		}()
//	}
//	fmt.Println("after loop")
//	group.Wait()
//	//time.Sleep(1 * 30 * time.Second)
//}
