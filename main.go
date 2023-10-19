package main

import (
	"fmt"
	"os"
)

func main() {
	p := pots(teams)
	printPots(os.Stdout, p)
	fmt.Println()

	d := newDrawer(p, 42)

	for i := 0; i < 20; i++ {
		x := d.draw()

		d.printGroups(os.Stdout, x)
		fmt.Println()
	}
}
