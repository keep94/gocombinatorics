// Package gocombinatorics contains routines useful for combinatorics.
package gocombinatorics

const (
  kSliceTooSmall = "Slice passed to Next is too small."
)

// Stream represents a finite stream of tuples
type Stream interface {

  // Next populates values with the next tuple and returns true.
  // If there are no more tuples, Next returns false and leaves values
  // unchanged. Caller must pass in a slice big enough to hold a tuple.
  Next(values []int) bool

  // Reset resets this stream to the state it had when it was first created.
  // After calling Reset, Next will yield the first tuple.
  Reset()
}

// Permutations yields all the ways you can pick k ints from 0 to n-1
// inclusive where order matters. The returned Stream's Next method will
// yield k-tuples.
//
// For instance, Permutations(4,2) yields
// (0,1), (0,2), (0,3), (1,0), (1,2), (1,3),
// (2,0), (2,1), (2,3), (3,0), (3,1), (3,2)
func Permutations(n, k int) Stream {
  if n < 0 {
    panic("n must be greater than or equal to 0")
  }
  if k < 0 {
    panic("k must be greater than or equal to 0")
  }
  var unused intSet
  var values []int
  if k > 0 {
    unused = newIntSet(n)
    values = make([]int, k)
  }
  result := &permutations{
      unused: unused,
      values: values,
      n: n,
      k: k,
  }
  result.Reset()
  return result
}

// OpsPosits yields all the possible positions of k binary operators in a
// postfix expression as a sequence of k-tuples. A value of 0 means the
// operator comes after the first number; 1 means the operator comes after
// the second number etc. Each value in a tuple is the same or greater
// than the previous value. Moreover the 1st value in a tuple must be
// at least 1; the 2nd value in a tuple must be at least 2 etc.
// None of the values in a tuple can be greater than k.
//
// For instance, OpsPosits(4) yeilds
// (1,2,3,4), (1,2,4,4), (1,3,3,4), (1,3,4,4), (1,4,4,4), (2,2,3,4)
// (2,2,4,4), (2,3,3,4), (2,3,4,4), (2,4,4,4), (3,3,3,4), (3,3,4,4)
// (3,4,4,4), (4,4,4,4)
func OpsPosits(k int) Stream {
  if k < 0 {
    panic("k must be greater than or equal to 0")
  }
  result := &opsPosits{values: make([]int, k), k: k}
  result.Reset()
  return result
}

// Product is like Permutations except that the returned tuples may contain
// duplicates. For instance, Product(3, 2) yields
// (0,0), (0,1), (0,2), (1,0), (1,1), (1,2), (2,0), (2,1), (2,2)
func Product(n, k int) Stream {
  if n < 0 {
    panic("n must be greater than or equal to 0")
  }
  if k < 0 {
    panic("k must be greater than or equal to 0")
  }
  result := &product{values: make([]int, k), n: n, k: k}
  result.Reset()
  return result
}

type opsPosits struct {
  values []int
  k int
  done bool
}

func (o *opsPosits) Next(values []int) bool {
  if len(values) < o.k {
    panic(kSliceTooSmall)
  }
  if o.done {
    return false
  }
  copy(values, o.values)
  o.increment()
  return true
}

func (o *opsPosits) Reset() {
  o.done = false
  for i := 0; i < o.k; i++ {
    o.values[i] = i + 1
  }
}

func (o *opsPosits) increment() {
  idx := o.k - 1
  for idx >= 0 && o.values[idx] == o.k {
    idx--
  }
  if idx < 0 {
    o.done = true
    return
  }
  o.values[idx]++
  for i := idx + 1; i < o.k; i++ {
    o.values[i] = max(o.values[idx], i + 1)
  }
}

func max(i, j int) int {
  if i > j {
    return i
  }
  return j
}

type permutations struct {
  // Everything except the values preceding the value being changed.
  // But if k = 0 unused is nil, the empty set.
  unused intSet

  // The values of the current tuple
  values []int

  n int
  k int
  done bool
}

func (p *permutations) Next(values []int) bool {
  if len(values) < p.k {
    panic(kSliceTooSmall)
  }
  if p.done {
    return false
  }
  copy(values, p.values)
  p.increment()
  return true
}

func (p *permutations) Reset() {
  p.done = p.k > p.n
  if p.done || p.k == 0 {
    return
  }
  for i := 0; i < p.k; i++ {
    p.values[i] = i
  }

  // The last value will get changed so unused to contain everything
  // except the values preceding the last value in the tuple.
  for i := 0; i < p.k - 1; i++ {
    p.unused.Remove(i)
  }
  for i := p.k - 1; i < p.n; i++ {
    p.unused.Add(i)
  }
}

func (p *permutations) increment() {

  // Special case: when k=0 there is one permutation so we are done as
  // soon as we increment.
  if p.k == 0 {
    p.done = true
    return
  }
  idx := p.k - 1

  // Increment the last value
  p.values[idx] = p.unused.Next(p.values[idx]+1)

  // If we reached the end, try to increment the previous value while
  // keeping the invariant that p.unused is everything except the values
  // preceding the value being incremented.
  for p.values[idx] == -1 {

    // If we reached the end when incrementing the very first value then
    // we are done.
    if idx == 0 {
      p.done = true
      return
    }
    idx--
    p.unused.Add(p.values[idx])
    p.values[idx] = p.unused.Next(p.values[idx]+1)
  }

  // After we have successfully incremented a value, fill in the slots
  // that come after that value with the smallest possible values.
  last := -1
  for idx < p.k - 1 {
    p.unused.Remove(p.values[idx])
    idx++
    p.values[idx] = p.unused.Next(last+1)

    // We know the next value we fill in has to be at least one greater
    // then the this value
    last = p.values[idx]
  }
}

type product struct {
  values []int
  n int
  k int
  done bool
}

func (p *product) Next(values []int) bool {
  if len(values) < p.k {
    panic(kSliceTooSmall)
  }
  if p.done {
    return false
  }
  copy(values, p.values)
  p.increment()
  return true
}

func (p *product) Reset() {
  p.done = p.n == 0 && p.k > 0
  for i := 0; i < p.k; i++ {
    p.values[i] = 0
  }
}

func (p *product) increment() {
  idx := p.k - 1
  for idx >= 0 && p.values[idx] == p.n - 1 {
    p.values[idx] = 0
    idx--
  }
  if idx < 0 {
    p.done = true
    return
  }
  p.values[idx]++
}
