package main

import (
	"fmt"
	"go-poc/jsonexample"
	"go-poc/functional"
)

func main() {
	jsonexample.Run()
	var double func (int) int = func(x int) int { return x * 2 }
	var slice []int = []int{1,2,3}
	fmt.Println(functional.Fmap(double, slice))
}
