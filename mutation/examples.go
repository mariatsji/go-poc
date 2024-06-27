package mutation

import (
	"fmt"
)

// immutable!
func Primitives() {
	var a int = 3
	var b = a
	a = 2 
	fmt.Println(a) // a mutated
	fmt.Println(b) // b is still the same

	i := &a
	fmt.Println(*i)
	i = &b

	fmt.Println(*i)
}

// immutable!
func Arrays() {
	a := [2]int{1,2}
	var b = a
	a[0] = 0
	fmt.Println(a)
	fmt.Println(b)

}

// mutable reference!
func Slices() {
	a := [2]int{1,2}
	var s = a[:]
	a[0] = 0
	fmt.Println(a)
	fmt.Println(s)
}

func Run() {
	Slices()
}