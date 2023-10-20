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
	s := newStats(p)

	i := 0
	for {
		x := d.draw()
		s.add(x)
		i++

		if i > 1000000 {
			break
		}
	}

	for _, t := range p {
		id := t.id
		fmt.Printf("%s: ", id)
		printProbs(os.Stdout, s.getProbs(id))
		fmt.Println()
	}
}
