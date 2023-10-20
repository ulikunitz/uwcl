package main

import (
	"fmt"
	"io"
	"sort"
)

type stats struct {
	pots   []*team
	n      int
	counts [][]int
}

func newStats(p []*team) *stats {
	s := &stats{
		pots:   p,
		counts: make([][]int, len(p)),
	}
	for i := range s.counts {
		s.counts[i] = make([]int, len(p))
	}
	return s
}

func (s *stats) addGroup(g []int) {
	for i, a := range g {
		for j, b := range g {
			if i == j {
				continue
			}
			s.counts[a][b]++
		}
	}
}

func (s *stats) add(d [][]int) {
	for _, g := range d {
		s.addGroup(g)
	}
	s.n++
}

type prob struct {
	team string
	prob float64
}

func (s *stats) teamIndex(team string) int {
	for i, t := range s.pots {
		if t.id == team {
			return i
		}
	}
	panic(fmt.Errorf("unknown team %q", team))
}

func (s *stats) getProbs(team string) []prob {
	t := s.teamIndex(team)
	p := make([]prob, 0, len(s.pots))
	r := s.counts[t]
	for i, t := range r {
		if t == 0 {
			continue
		}
		p = append(p,
			prob{
				team: s.pots[i].id,
				prob: float64(t) / float64(s.n),
			})
	}
	sort.Slice(p, func(i, j int) bool {
		return p[i].prob > p[j].prob
	})
	return p
}

func printProbs(w io.Writer, p []prob) {
	for i, t := range p {
		if i != 0 {
			fmt.Fprint(w, " ")
		}
		fmt.Fprintf(w, "%s %.0f%%", t.team, t.prob*100)
	}
}
