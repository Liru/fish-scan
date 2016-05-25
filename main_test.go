package main

import (
	"fmt"
)

func ExampleOhioMackerel() {
	resultChan := make(chan result, 1)
	Compare("mackerels", states, resultChan)
	close(resultChan)
	fmt.Println(<-resultChan)
	// Output:
	// ohio is the only state that doesn’t share any letters with the word "mackerels".
}
