package intmap

// Set is a specialization of Map modelling a set of integers.
// Like Map, methods that read from the set are valid on the nil Set.
// This include Has, Len, and ForEach.
type Set[K IntKey] Map[K, struct{}]

// NewSet creates a new Set with a given initial capacity.
func NewSet[K IntKey](capacity int) *Set[K] {
	return (*Set[K])(New[K, struct{}](capacity))
}

// Add an element to the set. Returns true if the element was not already present.
func (s *Set[K]) Add(k K) bool {
	_, found := (*Map[K, struct{}])(s).PutIfNotExists(k, struct{}{})
	return found
}

// Has returns true if the key is in the set.
// If the set is nil this method always return false.
func (s *Set[K]) Has(k K) bool {
	return (*Map[K, struct{}])(s).Has(k)
}

// Len returns the number of elements in the set.
// If the set is nil this method return 0.
func (s *Set[K]) Len() int {
	return (*Map[K, struct{}])(s).Len()
}

// ForEach iterates the elements in the set.
// This method returns immediately if the set is nil.
func (s *Set[K]) ForEach(visit func(k K) bool) {
	(*Map[K, struct{}])(s).ForEach(func(k K, _ struct{}) bool {
		return visit(k)
	})
}
