package functional

import (
	"fmt"
)

type Maybe[A any] struct {
	some A
	none bool // Indicates whether there is a value or not.
}

func Some[A any](value A) Maybe[A] {
	return Maybe[A]{some: value, none: false}
}

func None[A any]() Maybe[A] {
	return Maybe[A]{none: true}
}

func Empty[A any](m Maybe[A]) bool {
	return m.none
}

func Get[A any](m Maybe[A]) A {
	return m.some
}

func DoIf[A any](m Maybe[A], fn func(A)) {
	if !m.none {
		fn(m.some)
	}
}

func Fmap[A any, B any](fn func(A) B, s []A) []B {
	retVal := []B{}
	for i := 0; i < len(s); i++ {
		retVal = append(retVal, fn(s[i]))
	}
	return retVal
}

func ForEach[A any](fn func(A), s []A) {
	for i := 0; i < len(s); i++ {
		fn(s[i])
	}
}

func FoldL[A any, B any](fn func(B, A) B, b B, as []A) B {
	if len(as) == 0 {
		return b
	} else {
		return FoldL(fn, fn(b, as[0]), as[1:])
	}
}

func Filter[A any](a []A, fn func(A) bool) (ret []A) {
	for _, v := range a {
		if fn(v) {
			ret = append(ret, v)
		}
	}
	return ret
}

func Run() {
	var double func(int) int = func(x int) int { return x * 2 }
	var slice []int = []int{1, 2, 3}
	fmt.Println(Fmap(double, slice))

	for _, i := range new([100]int)[0:50] {
		fmt.Println(i)
	}

	cool := [...]string{"yo", "dude"}

	for _, s := range cool {
		fmt.Println(s)
	}

	xs := new([10]int)[0:3]

	for i := range xs {
		fmt.Println(i)
	}

	type Reader struct {
		things []string
	}

	myThings := [...]string{"database", "connections"}

	r := &Reader{myThings[:]}

	r.things[0] = "floppydisk"

	for _, thing := range r.things {
		fmt.Println(thing)
	}

	ints := []int{1, 2, 3, 4, 5}

	fmt.Printf("Folded: %s", FoldL(func(r string, i int) string { return r + fmt.Sprint(i) }, "", ints))

}
