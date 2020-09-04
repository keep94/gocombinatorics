package gocombinatorics

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntSet(t *testing.T) {
	assertIntSet(t, 11, 7, 1, 4, 3, 10, 5, 0, 2, 8, 9, 6)
	assertIntSet(t, 5, 3, 1, 4, 2, 0)
	assertIntSet(t, 1, 0)
	assertIntSet(t, 16, 10, 14, 2, 8, 9)
}

func TestString(t *testing.T) {
	assert := assert.New(t)
	s := newIntSet(10)
	assert.Equal("{}", s.String())
	s.Add(7)
	s.Add(2)
	s.Add(5)
	s.Add(3)
	assert.Equal("{2 3 5 7}", s.String())
	assert.Equal("{}", newIntSet(0).String())
}

func TestZero(t *testing.T) {
	assert := assert.New(t)
	s := newIntSet(0)
	assert.Nil(s)
	assert.Equal(-1, s.Next(0))
	assert.False(s.Contains(0))
	assert.Panics(func() { s.Add(0) })
	s.Remove(0)
	assert.Equal(0, s.Len())
	assert.Equal(0, s.Cap())
}

func TestMisc(t *testing.T) {
	assert := assert.New(t)
	s := newIntSet(3)
	assert.Equal(3, s.Cap())
	s.Add(0)
	s.Add(0)
	s.Add(2)
	s.Add(2)
	assert.True(s.Contains(2))
	assert.False(s.Contains(1))
	assert.True(s.Contains(0))
	assert.False(s.Contains(-1))
	assert.False(s.Contains(3))
	assert.Equal(2, s.Len())
	s.Remove(1)
	assert.Equal(2, s.Len())

	assert.Equal(-1, s.Next(3))
	assert.Equal(2, s.Next(2))
	assert.Equal(2, s.Next(1))
	assert.Equal(0, s.Next(0))
	assert.Equal(0, s.Next(-1))

	s.Remove(2)
	s.Remove(2)
	assert.Equal(1, s.Len())
	assert.False(s.Contains(2))
	assert.False(s.Contains(1))
	assert.True(s.Contains(0))

	assert.Panics(func() { s.Add(3) })
	assert.Panics(func() { s.Add(-1) })
	s.Remove(3)
	s.Remove(-1)
}

// Tests by adding values[0], values[1], ... to an intSet making sure the
// intSet can be correctly traversed after each add. Then it removes
// values[0], values[1], ... making sure the intSet can still be correctly
// traversed after each remove.
func assertIntSet(t *testing.T, size int, values ...int) {
	t.Helper()
	assert := assert.New(t)
	s := newIntSet(size)
	assert.Equal(0, s.Len())
	assert.Equal([]int{}, asSlice(s))
	for i := range values {
		s.Add(values[i])
		assert.Equal(sorted(values[0:i+1]), asSlice(s))
		assert.Equal(i+1, s.Len())
	}
	length := len(values)
	for i := range values {
		s.Remove(values[i])
		assert.Equal(sorted(values[i+1:length]), asSlice(s))
		assert.Equal(length-i-1, s.Len())
	}
}

func sorted(aSlice []int) []int {
	result := make([]int, len(aSlice))
	copy(result, aSlice)
	sort.Ints(result)
	return result
}

func asSlice(s intSet) []int {
	result := make([]int, 0, s.Len())
	for i := s.Next(0); i != -1; i = s.Next(i + 1) {
		result = append(result, i)
	}
	return result
}
