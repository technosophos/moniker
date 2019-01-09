package main

import (
	"fmt"
	".."
)

func main() {
	n := moniker.New()
	fmt.Printf("Your name is %q\n", n.Name())
	fmt.Printf("Your name is %q\n", n.NameWithOptions(" ", "á", true))
	fmt.Printf("Your name is %q\n", n.NameWithOptions(" ", "à", false))
}
