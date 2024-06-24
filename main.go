package main

import (
	"go-poc/jsonexample"
	"go-poc/functional"
	"go-poc/channels"
)

func main() {
	jsonexample.Run()
	functional.Run()
	channels.Run()
}
