package main

type HashSet map[int]int

func newValueHashSet(arr []int) HashSet {
	hs := HashSet{}

	for _, n := range arr {
		val, ok := hs[n]
		if ok {
			hs[n] = val + 1
		} else {
			hs[n] = 1
		}
	}

	return hs
}

func (h HashSet) Add(n int) {
	val, ok := h[n]
	if ok {
		h[n] = val + 1
	} else {
		h[n] = 1
	}
}

func (h HashSet) Del(n int) {
	val, ok := h[n]
	if ok {
		if val > 1 {
			h[n] = val - 1
		} else {
			delete(h, n)
		}
	}
}

func (h HashSet) Present(n int) bool {
	_, ok := h[n]
	return ok
}

func (h HashSet) Difference(a HashSet) []int {
	var diff []int

	for k0, v0 := range h {
		v1, ok := a[k0]
		if !ok {
			diff = append(diff, k0)
			continue
		}

		if v0 != v1 {
			delta := 0
			if v0 > v1 {
				delta = v0 - v1
			} else {
				delta = v1 - v0
			}

			for delta != 0 {
				diff = append(diff, k0)
				delta--
			}
		}
	}

	return diff
}

func canBeEqual(target, arr []int) bool {
	ts := newValueHashSet(target)
	as := newValueHashSet(arr)
	return len(ts.Difference(as)) == 0
}

func main() {
}
