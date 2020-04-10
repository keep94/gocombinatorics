package gocombinatorics_test

import (
  "strconv"
  "strings"
  "testing"

  "github.com/keep94/gocombinatorics"
  "github.com/stretchr/testify/assert"
)

func TestOpsPosits(t *testing.T) {
  assert := assert.New(t)
  stream := gocombinatorics.OpsPosits(4)
  assertStream(t, stream,
      "1 2 3 4", "1 2 4 4", "1 3 3 4", "1 3 4 4",
      "1 4 4 4", "2 2 3 4", "2 2 4 4", "2 3 3 4",
      "2 3 4 4", "2 4 4 4", "3 3 3 4", "3 3 4 4",
      "3 4 4 4", "4 4 4 4")
  stream = gocombinatorics.OpsPosits(1)
  assertStream(t, stream, "1")
  stream = gocombinatorics.OpsPosits(0)
  assertStream(t, stream, "")
  assert.Panics(func() { gocombinatorics.OpsPosits(-1) })
}

func TestCombinations(t *testing.T) {
  assert := assert.New(t)
  stream := gocombinatorics.Combinations(5, 3)
  assertStream(t, stream,
      "0 1 2", "0 1 3", "0 1 4", "0 2 3", "0 2 4",
      "0 3 4", "1 2 3", "1 2 4", "1 3 4", "2 3 4")
  stream = gocombinatorics.Combinations(5, 5)
  assertStream(t, stream, "0 1 2 3 4")
  stream = gocombinatorics.Combinations(5, 1)
  assertStream(t, stream, "0", "1", "2", "3", "4")
  stream = gocombinatorics.Combinations(5, 0)
  assertStream(t, stream, "")
  stream = gocombinatorics.Combinations(5, 6)
  assertStream(t, stream)
  stream = gocombinatorics.Combinations(0, 0)
  assertStream(t, stream, "")
  stream = gocombinatorics.Combinations(0, 1)
  assertStream(t, stream)
  assert.Panics(func() { gocombinatorics.Combinations(3, -1) })
  assert.Panics(func() { gocombinatorics.Combinations(-1, 3) })
}

func TestCombinationsWithReplacement(t *testing.T) {
  assert := assert.New(t)
  stream := gocombinatorics.CombinationsWithReplacement(2, 3)
  assertStream(t, stream,
      "0 0 0", "0 0 1", "0 1 1", "1 1 1")
  stream = gocombinatorics.CombinationsWithReplacement(3, 4)
  assertStream(t, stream,
      "0 0 0 0", "0 0 0 1", "0 0 0 2", "0 0 1 1", "0 0 1 2", "0 0 2 2",
      "0 1 1 1", "0 1 1 2", "0 1 2 2", "0 2 2 2", "1 1 1 1", "1 1 1 2",
      "1 1 2 2", "1 2 2 2", "2 2 2 2")
  stream = gocombinatorics.CombinationsWithReplacement(4, 2)
  assertStream(t, stream,
      "0 0", "0 1", "0 2", "0 3", "1 1", "1 2", "1 3", "2 2", "2 3", "3 3")
  stream = gocombinatorics.CombinationsWithReplacement(4, 1)
  assertStream(t, stream, "0", "1", "2", "3")
  stream = gocombinatorics.CombinationsWithReplacement(0, 5)
  assertStream(t, stream)
  stream = gocombinatorics.CombinationsWithReplacement(0, 0)
  assertStream(t, stream, "")
  stream = gocombinatorics.CombinationsWithReplacement(1, 0)
  assertStream(t, stream, "")
  assert.Panics(func() { gocombinatorics.CombinationsWithReplacement(3, -1) })
  assert.Panics(func() { gocombinatorics.CombinationsWithReplacement(-1, 3) })
}

func TestPermutations(t *testing.T) {
  assert := assert.New(t)
  stream := gocombinatorics.Permutations(4, 2)
  assertStream(t, stream,
      "0 1", "0 2", "0 3",
      "1 0", "1 2", "1 3",
      "2 0", "2 1", "2 3",
      "3 0", "3 1", "3 2")
  stream = gocombinatorics.Permutations(4, 4)
  assertStream(t, stream,
      "0 1 2 3", "0 1 3 2", "0 2 1 3", "0 2 3 1", "0 3 1 2", "0 3 2 1",
      "1 0 2 3", "1 0 3 2", "1 2 0 3", "1 2 3 0", "1 3 0 2", "1 3 2 0",
      "2 0 1 3", "2 0 3 1", "2 1 0 3", "2 1 3 0", "2 3 0 1", "2 3 1 0",
      "3 0 1 2", "3 0 2 1", "3 1 0 2", "3 1 2 0", "3 2 0 1", "3 2 1 0",
  )
  stream = gocombinatorics.Permutations(4, 1)
  assertStream(t, stream, "0", "1", "2", "3")
  stream = gocombinatorics.Permutations(4, 0)
  assertStream(t, stream, "")
  stream = gocombinatorics.Permutations(4, 5)
  assertStream(t, stream)
  stream = gocombinatorics.Permutations(0, 0)
  assertStream(t, stream, "")
  stream = gocombinatorics.Permutations(0, 1)
  assertStream(t, stream)
  assert.Panics(func() { gocombinatorics.Permutations(3, -1) })
  assert.Panics(func() { gocombinatorics.Permutations(-1, 3) })
}

func BenchmarkPermutations(b *testing.B) {
  stream := gocombinatorics.Permutations(50, 50)
  values := make([]int, 50)
  for i := 0; i < b.N; i++ {
    if !stream.Next(values) {
      stream.Reset()
    }
  }
}

func BenchmarkCombinations(b *testing.B) {
  stream := gocombinatorics.Permutations(100, 50)
  values := make([]int, 50)
  for i := 0; i < b.N; i++ {
    if !stream.Next(values) {
      stream.Reset()
    }
  }
}

func TestProduct(t *testing.T) {
  assert := assert.New(t)
  stream := gocombinatorics.Product(3, 2)
  assertStream(t, stream,
               "0 0", "0 1", "0 2",
               "1 0", "1 1", "1 2",
               "2 0", "2 1", "2 2")
  stream = gocombinatorics.Product(3, 0)
  assertStream(t, stream, "")
  stream = gocombinatorics.Product(0, 0)
  assertStream(t, stream, "")
  stream = gocombinatorics.Product(1, 4)
  assertStream(t, stream, "0 0 0 0")
  stream = gocombinatorics.Product(0, 3)
  assertStream(t, stream)
  assert.Panics(func() { gocombinatorics.Product(-1, 3) })
  assert.Panics(func() { gocombinatorics.Product(3, -1) })
}

func TestCartesian(t *testing.T) {
  assert := assert.New(t)
  stream := gocombinatorics.Cartesian(3, 2, 4)
  assertStream(t, stream,
      "0 0 0", "0 0 1", "0 0 2", "0 0 3", "0 1 0", "0 1 1", "0 1 2", "0 1 3",
      "1 0 0", "1 0 1", "1 0 2", "1 0 3", "1 1 0", "1 1 1", "1 1 2", "1 1 3",
      "2 0 0", "2 0 1", "2 0 2", "2 0 3", "2 1 0", "2 1 1", "2 1 2", "2 1 3")
  stream = gocombinatorics.Cartesian()
  assertStream(t, stream, "")
  stream = gocombinatorics.Cartesian(3, 0, 4)
  assertStream(t, stream)
  assert.Panics(func() { gocombinatorics.Cartesian(3, -1) })
}

// Reads first tuple off stream, resets it, then reads first 2 tuples off
// stream, resets again, then reads first 3 tuples off stream etc. until
// all expected tuples are read off stream.
func assertStream(
    t *testing.T,
    stream gocombinatorics.Stream,
    results ...string) {
  t.Helper()
  assert := assert.New(t)
  values := make([]int, stream.TupleSize())

  // Go to len(results) + 1 so that we have a chance to reset the stream
  // after exhausing it.
  for i := 0; i <= len(results) + 1; i++ {
    for j := 0; j <= i; j++ {
      hasMore := stream.Next(values)
      if j >= len(results) {
        if !assert.False(hasMore, "There shouldn't be more tuples") {
          return
        }
      } else {
        if !assert.True(hasMore, "There should be more tuples") {
          return
        }
        valueStr := asString(values)
        makeZero(values)  // Make sure stream has its own copy of values
        if !assert.Equal(results[j], valueStr) {
          return
        }
      }
    }
    stream.Reset()
  }
}

func makeZero(values []int) {
  for i := range values {
    values[i] = 0
  }
}

func asString(values []int) string {
  strs := make([]string, len(values))
  for i := range values {
    strs[i] = strconv.Itoa(values[i])
  }
  return strings.Join(strs, " ")
}
