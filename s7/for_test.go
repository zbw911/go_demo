package s7

import (
	"fmt"
	"testing"
)

func Test_for1(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func Test_for2(t *testing.T) {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func Test_for3(t *testing.T) {
	for {
		fmt.Println(1)
	}
}
