package main

import (
	"fmt"

	"github.com/nifle3/gosay/internal/image"
)

func main() {
	_, err := fmt.Print(image.Gopher)
	if err != nil {
		fmt.Print("OOPS")
	}
}
