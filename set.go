package main

type set[E comparable] map[E]struct{}

func newSet[E comparable]() set[E] {
	return set[E]{}
}

func (s *set[E]) add(vals ...E) {
	for _, v := range vals {
		(*s)[v] = struct{}{}
	}
}

func (s set[E]) values() []E {
	vals := make([]E, 0, len(s))
	for v := range s {
		vals = append(vals, v)
	}
	return vals
}
