package main

import (
	"fmt"

	"github.com/faiface/pixel"
)

func main() {
	rect := pixel.R(1, 1, 7, 7)
	fmt.Println(rect.W(), rect.H())
	fmt.Println(rect.Size())
}
