package s5

import (
	"fmt"
	"testing"
)

type mytestInt64 int64

func Test_convert(t *testing.T) {
	var a int = 1
	var b int64

	//b = a   // here is error
	b = int64(a)

	fmt.Println(a, b)

	var c mytestInt64

	//c = b //error , type 不对

	c = mytestInt64(b)

	fmt.Println(b, c)

}

func Test_Point(t *testing.T) {
	a := 1

	aPtr := &a

	//aPtr = aPtr + 4 //不支持

	fmt.Println(aPtr)
}

func TestStringisNil(t *testing.T) {
	//str := "aaaaaa"

	//这里报错
	//if str == nil {
	//
	//}

	var s string

	if s == "" {
		// do .....
	}
}
