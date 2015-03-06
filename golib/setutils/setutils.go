package setutils

type IntSet struct {
	Set map[int]bool
}

func NewIntSet() *IntSet {
	return &IntSet{Set: make(map[int]bool)}
}

func (set *IntSet) Add(n int) bool {
	_, found := set.Set[n]
	set.Set[n] = true
	return !found
}

func (set *IntSet) Remove(n int) bool {
	_, found := set.Set[n]
	if found {
		delete(set.Set, n)
	}
	return found
}

func (set *IntSet) Contains(n int) bool {
	_, found := set.Set[n]
	return found
}
