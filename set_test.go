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
	s.ForEach(func(k int) bool {
		sum += k
		return true
	})
	if sum != 3 {
		t.Fatalf("total sum of elements must be 3")
	}
}

func TestSetClear(t *testing.T) {
	s := NewSet[int](0)

	s.Add(1)
	s.Add(2)
	if sz := s.Len(); sz != 2 {
		t.Fatalf("unexpected set len %d", sz)
	}
	if !s.Has(1) || !s.Has(2) {
		t.Fatalf("set must contain 1 and 2 before Clear()")
	}

	s.Clear()
	if sz := s.Len(); sz != 0 {
		t.Fatalf("unexpected set len %d", sz)
	}
	if s.Has(1) || s.Has(2) {
		t.Fatalf("set must not contain 1 or 2 after Clear()")
	}
}

func TestSetDel(t *testing.T) {
	s := NewSet[int](0)

	s.Add(1)
	s.Add(2)
	if sz := s.Len(); sz != 2 {
		t.Fatalf("unexpected set len %d", sz)
	}
	if !s.Has(1) || !s.Has(2) {
		t.Fatalf("set must contain 1 and 2 before Clear()")
	}

	if found := s.Del(27); found {
		t.Fatalf("set must not contain 27")
	}

	// Delete 2
	if found := s.Del(2); !found {
		t.Fatalf("set must contain 2 on delete")
	}
	if sz := s.Len(); sz != 1 {
		t.Fatalf("unexpected set len %d", sz)
	}
	if s.Has(2) {
		t.Fatalf("set must not contain 2 after Del(2)")
	}
	if found := s.Del(2); found {
		t.Fatalf("set must not contain 2 on seconb deletion")
	}

	// Delete 1
	if found := s.Del(1); !found {
		t.Fatalf("set must contain 1 on delete")
	}
	if sz := s.Len(); sz != 0 {
		t.Fatalf("unexpected set len %d", sz)
	}
	if s.Has(1) {
		t.Fatalf("set must not contain 1 after Del(1)")
	}
	if found := s.Del(1); found {
		t.Fatalf("set must not contain 1 on seconb deletion")
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
	s.ForEach(func(k int) bool {
		count++
		return true
	})
	if count != 0 {
		t.Fatalf("total count of elements in nil set must be 0")
	}

}
