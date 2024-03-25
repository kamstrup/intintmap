package intmap

import "testing"

func TestSet(t *testing.T) {
	s := NewSet[int](0)
	if sz := s.Len(); sz != 0 {
		t.Fatalf("length of set must be 0: %d", sz)
	}

	if added := s.Add(1); !added {
		t.Fatalf("1 must be added")
	}
	if sz := s.Len(); sz != 1 {
		t.Fatalf("length of set must be 1: %d", sz)
	}

	if added := s.Add(1); added {
		t.Fatalf("1 must not be added")
	}
	if sz := s.Len(); sz != 1 {
		t.Fatalf("length of set must be 1: %d", sz)
	}

	if added := s.Add(2); !added {
		t.Fatalf("2 must be added")
	}
	if sz := s.Len(); sz != 2 {
		t.Fatalf("length of set must be 2: %d", sz)
	}

	if !(s.Has(1) && s.Has(2)) {
		t.Fatalf("set must have both 1 and 2")
	}

	sum := 0
	s.ForEach(func(k int) {
		sum += k
	})
	if sum != 3 {
		t.Fatalf("total sum of elements must be 3")
	}
}

func TestNilSet(t *testing.T) {
	var s *Set[int]

	if sz := s.Len(); sz != 0 {
		t.Fatalf("length of nil set must be 0: %d", sz)
	}

	if s.Has(0) || s.Has(1) {
		t.Fatalf("nil set must not have 0 or 1")
	}

	count := 0
	s.ForEach(func(k int) {
		count++
	})
	if count != 0 {
		t.Fatalf("total count of elements in nil set must be 0")
	}

}
