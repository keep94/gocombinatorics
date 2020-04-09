package gocombinatorics

import (
  "fmt"
  "strconv"
  "strings"
)

// intSet represents an ordered set of integers.
type intSet []bool

// newIntSet returns a new, empty intSet with a capcity of max. That is the
// intSet can contain integers from 0 up to but not including max. If max=0,
// newIntSet returns nil which represents an empty intSet with capacity of 0.
func newIntSet(max int) intSet {
  if max <= 0 {
    return nil
  }
  return make([]bool, max)
}

// Add adds x to this set. Add panics if x is less than 0 or greater than
// or equal to s.Cap(). Add runs in O(1) time.
func (s intSet) Add(x int) {
  if x < 0 || x >= len(s) {
    panic("Value out of range")
  }
  s[x] = true
}

// Remove removes x from this set. Remove runs in O(1) time.
func (s intSet) Remove(x int) {
  if x < 0 || x >= len(s) {
    return
  }
  s[x] = false
}

// Cap returns the capacity of this set.
func (s intSet) Cap() int {
  return len(s)
}

// Len returns the number of integers in this set. Len runs in O(N) time
// where N is the capacity of this set.
func (s intSet) Len() int {
  result := 0
  for i := range s {
    if s[i] {
      result++
    }
  }
  return result
}

// Contains returns true if this set contains x. Contains runs in O(1) time
func (s intSet) Contains(x int) bool {
  if x < 0 || x >= len(s) {
    return false
  }
  return s[x]
}

// Next returns the smallest integer in this set that is greater than or
// equal to x. If there is no integer greater than or equal to x in this
// set, Next returns -1. Next operates in near constant time if this set
// is nearly full. If this set is sparse, Next operates in O(N) worst
// case where N is the capacity of this set.
func (s intSet) Next(x int) int {
  if x < 0 {
    x = 0
  }
  for ; x < len(s); x++ {
    if s[x] {
      return x
    }
  }
  return -1
}

func (s intSet) String() string {
  strs := make([]string, 0, s.Len())
  for num := s.Next(0); num != -1; num = s.Next(num + 1) {
    strs = append(strs, strconv.Itoa(num))
  }
  return fmt.Sprintf("{%s}", strings.Join(strs, " "))
}
