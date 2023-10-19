package main

import (
	"fmt"
	"io"
	"math/bits"
	"math/rand"
	"slices"
)

func pots(teams []team) []*team {
	pots := make([]*team, 0, 16)

	// Add teams in pot1.
	for i, t := range teams {
		if t.pot1 {
			pots = append(pots, &teams[i])
		}
	}
	// Add remaining teams
	for i, t := range teams {
		if !t.pot1 {
			pots = append(pots, &teams[i])
		}
	}

	slices.SortFunc(pots[4:], func(a, b *team) int {
		x, y := a.coeff, b.coeff
		switch {
		case x < y:
			return 1
		case x > y:
			return -1
		default:
			return 0
		}
	})

	return pots
}

func printPots(w io.Writer, pots []*team) {
	for i := 0; i < 4; i++ {
		fmt.Fprintf(w, "Pot %d: ", i+1)
		for j := 0; j < 4; j++ {
			if j > 0 {
				fmt.Fprint(w, ", ")
			}
			fmt.Fprintf(w, "%s", pots[i*4+j].id)
		}
		fmt.Fprintln(w)
	}
}

const (
	red    uint8 = 0b0011
	blue   uint8 = 0b1100
	groupA uint8 = 0b0001
	groupB uint8 = 0b0010
	groupC uint8 = 0b0100
	groupD uint8 = 0b1000
)

type drawer struct {
	pots   []*team
	rng    *rand.Rand
	groups []uint8
}

func randomSelect(rng *rand.Rand, s uint16) int {
	n := bits.OnesCount16(s)
	if n == 0 {
		panic("no bits to select")
	}
	i := rng.Intn(n)
	for {
		j := bits.TrailingZeros16(s)
		if i == 0 {
			return j
		}
		i--
		s &= s - 1
	}
}

func newDrawer(pots []*team, seed int64) *drawer {
	d := &drawer{
		pots: pots,
		rng:  rand.New(rand.NewSource(seed)),
	}

	return d
}

// draw draws the group for all teams observing the constraints.
func (d *drawer) draw() [][]int {
	d.groups = make([]uint8, len(d.pots))
	for i := range d.groups {
		d.groups[i] = 0b1111
	}
	for i := 0; i < 4; i++ {
		d.drawPot(i)
	}

	r := make([][]int, 4)

	for j, g := range d.groups {
		switch g {
		case groupA:
			r[0] = append(r[0], j)
		case groupB:
			r[1] = append(r[1], j)
		case groupC:
			r[2] = append(r[2], j)
		case groupD:
			r[3] = append(r[3], j)
		default:
			panic("invalid group")
		}
	}

	return r
}

// drawPot draws all tests from a pot. Pots are numbered 0 to 3.
func (d *drawer) drawPot(pot int) {
	p := 0b1111 << (pot * 4)
	for i := 0; i < 4; i++ {
		j := randomSelect(d.rng, uint16(p))
		p &^= 1 << j
		g := randomSelect(d.rng, uint16(d.groups[j]))
		d.groups[j] = 1 << g
		d.handle(j)
	}
}

func (d *drawer) handle(j int) {
	g := d.groups[j]
	n := bits.OnesCount8(g)
	if n > 2 {
		return
	}
	if n == 2 {
		d.handle2(j)
		return
	}
	if n == 0 {
		panic(fmt.Errorf("team %d no group left", j))
	}
	d.handlePot(j)
	d.handleCountry(j)
	d.handleColor(j)
}

func (d *drawer) handle2(j int) {
	g := d.groups[j]
	n := bits.OnesCount8(g)
	if n != 2 {
		return
	}
	p := j / 4
	for i := p * 4; i < (p+1)*4; i++ {
		if i == j {
			continue
		}
		if g == d.groups[i] {
			for k := p * 4; k < (p+1)*4; k++ {
				if k == j || k == i {
					continue
				}
				if bits.OnesCount8(d.groups[k]) <= 2 {
					continue
				}
				d.groups[k] &^= g
				d.handle(k)
			}
		}
	}
}

func (d *drawer) handlePot(j int) {
	p := j / 4
	for i := p * 4; i < (p+1)*4; i++ {
		if i == j {
			continue
		}
		if bits.OnesCount8(d.groups[i]) == 1 {
			continue
		}
		d.groups[i] &^= d.groups[j]
		d.handle(i)
	}
}

func (d *drawer) handleCountry(j int) {
	country := d.pots[j].country
	g := d.groups[j]
	for i := range d.groups {
		if i == j {
			continue
		}
		if d.pots[i].country == country {
			if bits.OnesCount8(d.groups[i]) == 1 {
				continue
			}
			d.groups[i] &^= g
			d.handle(i)
		}
	}
}

func (d *drawer) findPaired(j int) int {
	paired := d.pots[j].paired
	if paired == "" {
		return -1
	}
	for i := range d.pots {
		if i == j {
			continue
		}
		if d.pots[i].id == paired {
			return i
		}
	}
	return -1
}

func (d *drawer) handleColor(j int) {
	i := d.findPaired(j)
	if i < 0 {
		return
	}
	if bits.OnesCount8(d.groups[i]) == 1 {
		return
	}
	if d.groups[j]&red != 0 {
		d.groups[i] &= blue
	} else {
		d.groups[i] &= red
	}
	d.handle(i)
}

func (d *drawer) printGroups(w io.Writer, groups [][]int) {
	for g, y := range groups {
		fmt.Fprintf(w, "%c:", 'A'+g)
		for k, z := range y {
			if k > 0 {
				fmt.Fprintf(w, ", ")
			}
			fmt.Fprintf(w, " %s", d.pots[z].id)
		}
		fmt.Fprintln(w)
	}
}
