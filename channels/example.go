package channels

import(
	"fmt"
)

func Fn(c chan string) {
	c <- "Hello"
	c <- "World"
	close(c)
}

func Run() {
	ch := make(chan string)
    go Fn(ch)
	for x := range ch {
		fmt.Printf("%s <- | ", x)
	}
}