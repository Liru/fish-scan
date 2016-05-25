package main

import (
	"fmt"
	"strings"
	"sync"
)

var wg sync.WaitGroup

type result struct {
	state, word string
}

func (r result) String() string {
	return fmt.Sprintf("%s is the only state that doesn’t share any letters with the word \"%s\".",
		r.state, r.word)
}

func MassCompare(words, states []string) <-chan result {
	resultChan := make(chan result, len(Dictionary))
	for _, word := range words {
		Compare(word, states, resultChan)
	}

	close(resultChan)
	return resultChan
}

func Compare(word string, states []string, resultChan chan<- result) {

	stateNote := ""
	for _, state := range states {
		if !strings.ContainsAny(state, word) {
			if stateNote != "" {
				return
			}
			stateNote = state
		}
	}

	if stateNote != "" {
		resultChan <- result{
			state: stateNote,
			word:  word,
		}
	}
}

func main() {
	LoadDictionary()
	resultChan := MassCompare(Dictionary, states)
	for result := range resultChan {
		fmt.Println(result)
	}
}
