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
  assertStream(t, stream, 4,
      "1 2 3 4", "1 2 4 4", "1 3 3 4", "1 3 4 4",
      "1 4 4 4", "2 2 3 4", "2 2 4 4", "2 3 3 4",
      "2 3 4 4", "2 4 4 4", "3 3 3 4", "3 3 4 4",
      "3 4 4 4", "4 4 4 4")
  stream = gocombinatorics.OpsPosits(1)
  assertStream(t, stream, 1, "1")
  stream = gocombinatorics.OpsPosits(0)
  assertStream(t, stream, 0, "")
  assert.Panics(func() { gocombinatorics.OpsPosits(-1) })
}

func TestCombinations(t *testing.T) {
  assert := assert.New(t)
  stream := gocombinatorics.Combinations(5, 3)
  assertStream(t, stream, 3,
      "0 1 2", "0 1 3", "0 1 4", "0 2 3", "0 2 4",
      "0 3 4", "1 2 3", "1 2 4", "1 3 4", "2 3 4")
  stream = gocombinatorics.Combinations(5, 5)
  assertStream(t, stream, 5, "0 1 2 3 4")
  stream = gocombinatorics.Combinations(5, 1)
  assertStream(t, stream, 1, "0", "1", "2", "3", "4")
  stream = gocombinatorics.Combinations(5, 0)
  assertStream(t, stream, 0, "")
  stream = gocombinatorics.Combinations(5, 6)
  assertStream(t, stream, 6)
  stream = gocombinatorics.Combinations(0, 0)
  assertStream(t, stream, 0, "")
  stream = gocombinatorics.Combinations(0, 1)
  assertStream(t, stream, 1)
  assert.Panics(func() { gocombinatorics.Combinations(3, -1) })
  assert.Panics(func() { gocombinatorics.Combinations(-1, 3) })
}
  

func TestPermutations(t *testing.T) {
  assert := assert.New(t)
  stream := gocombinatorics.Permutations(4, 2)
  assertStream(t, stream, 2,
      "0 1", "0 2", "0 3",
      "1 0", "1 2", "1 3",
      "2 0", "2 1", "2 3",
      "3 0", "3 1", "3 2")
  stream = gocombinatorics.Permutations(4, 4)
  assertStream(t, stream, 4,
      "0 1 2 3", "0 1 3 2", "0 2 1 3", "0 2 3 1", "0 3 1 2", "0 3 2 1",
      "1 0 2 3", "1 0 3 2", "1 2 0 3", "1 2 3 0", "1 3 0 2", "1 3 2 0",
      "2 0 1 3", "2 0 3 1", "2 1 0 3", "2 1 3 0", "2 3 0 1", "2 3 1 0",
      "3 0 1 2", "3 0 2 1", "3 1 0 2", "3 1 2 0", "3 2 0 1", "3 2 1 0",
  )
  stream = gocombinatorics.Permutations(4, 1)
  assertStream(t, stream, 1, "0", "1", "2", "3")
  stream = gocombinatorics.Permutations(4, 0)
  assertStream(t, stream, 0, "")
  stream = gocombinatorics.Permutations(4, 5)
  assertStream(t, stream, 5)
  stream = gocombinatorics.Permutations(0, 0)
  assertStream(t, stream, 0, "")
  stream = gocombinatorics.Permutations(0, 1)
  assertStream(t, stream, 1)
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
  assertStream(t, stream, 2,
               "0 0", "0 1", "0 2",
               "1 0", "1 1", "1 2",
               "2 0", "2 1", "2 2")
  stream = gocombinatorics.Product(3, 0)
  assertStream(t, stream, 0, "")
  stream = gocombinatorics.Product(0, 0)
  assertStream(t, stream, 0, "")
  stream = gocombinatorics.Product(1, 4)
  assertStream(t, stream, 4, "0 0 0 0")
  stream = gocombinatorics.Product(0, 3)
  assertStream(t, stream, 3)
  assert.Panics(func() { gocombinatorics.Product(-1, 3) })
  assert.Panics(func() { gocombinatorics.Product(3, -1) })
}

func assertStream(
    t *testing.T,
    stream gocombinatorics.Stream,
    tupleSize int,
    results ...string) {
  t.Helper()
  assertStreamOnce(t, stream, tupleSize, results...)
  stream.Reset()
  assertStreamOnce(t, stream, tupleSize, results...)
}

func assertStreamOnce(
    t *testing.T,
    stream gocombinatorics.Stream,
    tupleSize int,
    results ...string) {
  t.Helper()
  assert := assert.New(t)
  values := make([]int, tupleSize)
  idx := 0
  for stream.Next(values) {
    if idx == len(results) {
      assert.Fail("Stream should have quit emitting values")
      return
    }
    valueStr := asString(values)
    makeZero(values)
    assert.Equal(results[idx], valueStr)
    idx++
  }
  assert.Equal(len(results), idx)
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
