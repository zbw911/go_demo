package s6

import "testing"

func TestArrayEq(t *testing.T) {
	a := [...]int{1, 2, 3}
	b := [...]int{1, 2, 4}

	t.Log(a == b)

	c := [...]int{1, 2, 4}

	t.Log(b == c)

	//数组维度不同，不能比较
	//d := [...]int{1, 2, 4, 5}
	//t.Log(b == d)
}
