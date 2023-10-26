package gocombinatorics_test

import (
	"strings"
	"testing"

	"github.com/keep94/gocombinatorics"
	"github.com/stretchr/testify/assert"
)

func TestTCombinations(t *testing.T) {
	stream := gocombinatorics.TCombinations(
		[]string{"red", "orange", "yellow", "green", "blue"}, 3)
	assert.Panics(t, func() { stream.Next(nil) })
	assertTStream(t, stream,
		"red orange yellow", "red orange green", "red orange blue",
		"red yellow green", "red yellow blue", "red green blue",
		"orange yellow green", "orange yellow blue", "orange green blue",
		"yellow green blue")
}

func TestTCombinationsWithReplacement(t *testing.T) {
	stream := gocombinatorics.TCombinationsWithReplacement(
		[]string{"zero", "one"}, 3)
	assert.Panics(t, func() { stream.Next(nil) })
	assertTStream(t, stream,
		"zero zero zero", "zero zero one", "zero one one", "one one one")
}

func TestTPermutations(t *testing.T) {
	stream := gocombinatorics.TPermutations(
		[]string{"alpha", "beta", "gamma", "delta"}, 2)
	assert.Panics(t, func() { stream.Next(nil) })
	assertTStream(t, stream,
		"alpha beta", "alpha gamma", "alpha delta",
		"beta alpha", "beta gamma", "beta delta",
		"gamma alpha", "gamma beta", "gamma delta",
		"delta alpha", "delta beta", "delta gamma")
}

func TestTProduct(t *testing.T) {
	stream := gocombinatorics.TProduct([]string{"blue", "green", "red"}, 2)
	assert.Panics(t, func() { stream.Next(nil) })
	assertTStream(t, stream,
		"blue blue", "blue green", "blue red",
		"green blue", "green green", "green red",
		"red blue", "red green", "red red")
}

func TestZeroTStream(t *testing.T) {
	var stream gocombinatorics.TStream[string]
	assert.Zero(t, stream.TupleSize())
	assert.False(t, stream.Next(nil))
	stream.Reset()
}

// Reads first tuple off stream, resets it, then reads first 2 tuples off
// stream, resets again, then reads first 3 tuples off stream etc. until
// all expected tuples are read off stream.
func assertTStream(
	t *testing.T,
	stream *gocombinatorics.TStream[string],
	results ...string) {
	t.Helper()
	assert := assert.New(t)
	values := make([]string, stream.TupleSize())

	// Go to len(results) + 1 so that we have a chance to reset the stream
	// after exhausing it.
	for i := 0; i <= len(results)+1; i++ {
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
				valueStr := strings.Join(values, " ")
				makeEmpty(values) // Make sure stream has its own copy of values
				if !assert.Equal(results[j], valueStr) {
					return
				}
			}
		}
		stream.Reset()
	}
}

func makeEmpty(values []string) {
	for i := range values {
		values[i] = ""
	}
}
