package main

import "sort"

type set map[val]meta

func newSet(entries ...entry) (s set) {
	s = set{}
	for _, e := range entries {
		s[e.Val] = e.Meta
	}
	return
}

func (s *set) add(entries ...entry) {
	for _, e := range entries {
		(*s)[e.Val] = e.Meta
	}
}

func (s set) entries() []entry {
	entries := make([]entry, 0, len(s))
	for k, v := range s {
		e := entry{
			Val:  k,
			Meta: v,
		}
		entries = append(entries, e)
	}
	// sort entries w.r.t last modified time
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Meta.LastModified.After(entries[j].Meta.LastModified)
	})
	return entries
}
