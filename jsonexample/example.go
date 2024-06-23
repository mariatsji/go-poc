package jsonexample

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

func ParseDir(d string) Direction {
	switch strings.ToLower(d) {
	case "north":
		return North
	case "east":
		return East
	case "west":
		return West
	default:
		return South
	}
}

func (d *Direction) UnmarshalJSON(b []byte) error {
	var dir string
	if err := json.Unmarshal(b, &dir); err != nil {
		return err
	}
	*d = ParseDir(dir)
	return nil
}

type Country struct {
	Name      string
	Capital   string
	Direction Direction
}

func Run() {
	indiaByte := []byte(`{
		"Name": "India",
		"Capital": "New Dehli",
		"Direction": "East"
	}`)

	var india Country

	if err := json.Unmarshal(indiaByte, &india); err != nil {
		fmt.Println("error parsing country", err)
	}

	fmt.Println("blob was: ", india)
}
