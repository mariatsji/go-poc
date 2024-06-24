package functional

import (
	"fmt"
)

func Fmap[A any, B any](fn (func (A) B), s []A) []B {
	retVal := []B{}
	for i := 0; i < len(s); i++ {
		retVal = append(retVal, fn(s[i]))
	}
	return retVal
}

func Run() {
	var double func (int) int = func(x int) int { return x * 2 }
	var slice []int = []int{1,2,3}
	fmt.Println(Fmap(double, slice))
}