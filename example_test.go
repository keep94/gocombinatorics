package gocombinatorics_test

import (
	"fmt"

	"github.com/keep94/gocombinatorics"
)

func ExampleTCombinations() {
	// Print out all the ways you can choose 2 marbles from an urn
	// containing a red, green, yellow, and blue marble without replacement
	// and where order doesn't matter.
	stream := gocombinatorics.TCombinations(
		[]string{"red", "green", "yellow", "blue"}, 2)
	picked := make([]string, stream.TupleSize())
	for stream.Next(picked) {
		fmt.Println(picked)
	}
	// Output:
	// [red green]
	// [red yellow]
	// [red blue]
	// [green yellow]
	// [green blue]
	// [yellow blue]
}
