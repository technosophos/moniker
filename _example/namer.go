package main

import (
	"fmt"

	".."
)

func main() {
	n := moniker.NewAlliterator()
	fmt.Printf("Your name is %q\n", n.NameSepPrefix(" ", "q"))
}
