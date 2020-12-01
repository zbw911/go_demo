package s4_const

import (
	"fmt"
	"testing"
)

const (
	a = iota + 1
	B
	c
)

func Test_const(t *testing.T) {
	fmt.Println(a)
	fmt.Println(B)
	fmt.Println(c)
}

const (
	x = 1 << iota
	y
	z
)

func Test_bits(t *testing.T) {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(x | y | z)
}

const (
	i = "i_1asdfasd"
	j = "j"
	k = "k"
)

func Test_string_const(t *testing.T) {
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
}
