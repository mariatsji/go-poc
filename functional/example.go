package functional

import (
	"fmt"
)

type Maybe[A any] struct {
    some *A
    none bool // Indicates whether there is a value or not.
}

func Some[A any](value A) Maybe[A] {
    return Maybe[A]{some: &value, none: false}
}

func None[A any]() Maybe[A] {
    return Maybe[A]{some: nil, none: true}
}

func Empty[A any](m Maybe[A]) bool {
	return m.none
}

func Get[A any](m Maybe[A]) *A {
	return m.some
}

func DoIf[A any](m Maybe[A], fn (func (A))) {
	if !m.none {
		it := m.some
		fn(*it)
	}
}

func Fmap[A any, B any](fn (func (A) B), s []A) []B {
	retVal := []B{}
	for i := 0; i < len(s); i++ {
		retVal = append(retVal, fn(s[i]))
	}
	return retVal
}

func ForEach[A any](fn (func (A)), s []A) {
	for i := 0; i < len(s); i++ {
		fn(s[i])
	}
}

func Run() {
	var double func (int) int = func(x int) int { return x * 2 }
	var slice []int = []int{1,2,3}
	fmt.Println(Fmap(double, slice))
}