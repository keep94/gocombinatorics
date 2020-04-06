package gocombinatorics

import (
  "fmt"
  "strconv"
  "strings"
)

// intSet represents an ordered set of integers.
type intSet struct {
  layers [][]int
}

// newIntSet returns a new, empty intSet with a capcity of max. That is the
// intSet can contain integers from 0 up to but not including max. If max=0,
// newIntSet returns nil which represents an empty intSet with capacity of 0.
func newIntSet(max int) *intSet {
  if max <= 0 {
    return nil
  }
  var layers [][]int
  for max > 1 {
    layer := make([]int, max)
    layers = append(layers, layer)
    max = (max + 1) / 2
  }
  layers = append(layers, make([]int, 1))
  return &intSet{layers: layers}
}

// Add adds x to this set. Add panics if x is less than 0 or greater than
// or equal to s.Cap(). Add runs in O(log N) time where N is the capacity
// of this set.
func (s *intSet) Add(x int) {
  if s == nil || x < 0 || x >= len(s.layers[0]) {
    panic("Value out of range")
  }
  if s.layers[0][x] == 1 {
    return
  }
  layerNo := 0
  for layerNo < len(s.layers) {
    s.layers[layerNo][x]++
    x /= 2
    layerNo++
  }
}

// Remove removes x from this set. Remove runs in O(log N) time where N is
// the capacity of this set.
func (s *intSet) Remove(x int) {
  if s == nil || x < 0 || x >= len(s.layers[0]) {
    return
  }
  if s.layers[0][x] == 0 {
    return
  }
  layerNo := 0
  for layerNo < len(s.layers) {
    s.layers[layerNo][x]--
    x /= 2
    layerNo++
  }
}

// Cap returns the capacity of this set.
func (s *intSet) Cap() int {
  if s == nil {
    return 0
  }
  return len(s.layers[0])
}

// Len returns the number of integers in this set. Len runs in O(1) time.
func (s *intSet) Len() int {
  if s == nil {
    return 0
  }
  length := len(s.layers)
  return s.layers[length-1][0]
}

// Contains returns true if this set contains x. Contains runs in O(1) time
func (s *intSet) Contains(x int) bool {
  if s == nil || x < 0 || x >= len(s.layers[0]) {
    return false
  }
  return s.layers[0][x] == 1
}

// Next returns the smallest integer in this set that is greater than or
// equal to x. If there is no integer greater than or equal to x in this
// set, Next returns -1. Next operates in near constant time if this set
// is nearly full. If this set is sparse, Next operates in O(log N) worst
// case where N is the capacity of this set.
func (s *intSet) Next(x int) int {
  if x < 0 {
    x = 0
  }
  if s == nil || x >= len(s.layers[0]) {
    return -1
  }
  layerNo := 0
  for s.layers[layerNo][x] == 0 {
    if x + 1 == len(s.layers[layerNo]) {
      return -1
    }
    if x % 2 == 1 {
      x++
    } else {
      x /= 2
      layerNo++
    }
  }

  for layerNo > 0 {
    layerNo--
    x *= 2
    if s.layers[layerNo][x] == 0 {
      x++
    }
  }
  return x
}

func (s *intSet) String() string {
  strs := make([]string, 0, s.Len())
  for num := s.Next(0); num != -1; num = s.Next(num + 1) {
    strs = append(strs, strconv.Itoa(num))
  }
  return fmt.Sprintf("{%s}", strings.Join(strs, " "))
}
