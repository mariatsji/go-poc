package channels

import(
	"time"
	"fmt"
)

func Fn(s string, c chan string) {
	c <- s
}

func Run() {
	ch := make(chan string)
	go Fn("Hello", ch)
	go Fn("World", ch)
	x, y := <-ch, <-ch
	close(ch)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Channel receive: %s %s", x,y)
}