package gocombinatorics

import (
  "fmt"
  "strconv"
  "strings"
)

// intSet represents an ordered set of integers.
type intSet struct {
  membership []bool
}

// newIntSet returns a new, empty intSet with a capcity of max. That is the
// intSet can contain integers from 0 up to but not including max. If max=0,
// newIntSet returns nil which represents an empty intSet with capacity of 0.
func newIntSet(max int) *intSet {
  if max <= 0 {
    return nil
  }
  membership :=  make([]bool, max)
  return &intSet{membership: membership}
}

// Add adds x to this set. Add panics if x is less than 0 or greater than
// or equal to s.Cap(). Add runs in O(1) time.
func (s *intSet) Add(x int) {
  if s == nil || x < 0 || x >= len(s.membership) {
    panic("Value out of range")
  }
  s.membership[x] = true
}

// Remove removes x from this set. Remove runs in O(1) time.
func (s *intSet) Remove(x int) {
  if s == nil || x < 0 || x >= len(s.membership) {
    return
  }
  s.membership[x] = false
}

// Cap returns the capacity of this set.
func (s *intSet) Cap() int {
  if s == nil {
    return 0
  }
  return len(s.membership)
}

// Len returns the number of integers in this set. Len runs in O(N) time
// where N is the capacity of this set.
func (s *intSet) Len() int {
  if s == nil {
    return 0
  }
  result := 0
  for i := range s.membership {
    if s.membership[i] {
      result++
    }
  }
  return result
}

// Contains returns true if this set contains x. Contains runs in O(1) time
func (s *intSet) Contains(x int) bool {
  if s == nil || x < 0 || x >= len(s.membership) {
    return false
  }
  return s.membership[x]
}

// Next returns the smallest integer in this set that is greater than or
// equal to x. If there is no integer greater than or equal to x in this
// set, Next returns -1. Next operates in near constant time if this set
// is nearly full. If this set is sparse, Next operates in O(N) worst
// case where N is the capacity of this set.
func (s *intSet) Next(x int) int {
  if x < 0 {
    x = 0
  }
  if s == nil {
    return -1
  }
  for ; x < len(s.membership); x++ {
    if s.membership[x] {
      return x
    }
  }
  return -1
}

func (s *intSet) String() string {
  strs := make([]string, 0, s.Len())
  for num := s.Next(0); num != -1; num = s.Next(num + 1) {
    strs = append(strs, strconv.Itoa(num))
  }
  return fmt.Sprintf("{%s}", strings.Join(strs, " "))
}
