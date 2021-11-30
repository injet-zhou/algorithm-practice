package stack

import "testing"

func TestInitStack(t *testing.T) {
	s := InitStack(1,2,3,4,5,6)
	PrintStack(s)
}
