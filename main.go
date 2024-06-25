package main

import (
	"go-poc/jsonexample"
	"go-poc/functional"
	"go-poc/channels"
	"go-poc/chess"
)

func main() {
	jsonexample.Run()
	functional.Run()
	channels.Run()
	chess.Run()
}
