package main

import (
	"fmt"

	"github.com/technosophos/moniker"
)

func main() {
	n := moniker.NewWith(moniker.Got, moniker.Descriptors)
	fmt.Printf("Your name is %q\n", n.Name())
}
