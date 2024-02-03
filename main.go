package main

import (
	"strings"
)

type CircularBuffer struct {
	buffer     *[]rune
	size       int
	frontIndex int
}

func New(input *string) *CircularBuffer {
	inputLength := len(*input)

	buffer := make([]rune, inputLength)

	for i, ch := range *input {
		buffer[i] = ch
	}

	return &CircularBuffer{
		size:       inputLength,
		buffer:     &buffer,
		frontIndex: 0,
	}
}

func (cb *CircularBuffer) String() string {
	var stringBuilder strings.Builder

	iterationCount := 1
	currentIndex := cb.frontIndex

	for iterationCount <= cb.size {
		stringBuilder.WriteRune((*cb.buffer)[currentIndex])

		if currentIndex == cb.size-1 {
			currentIndex = 0
		} else {
			currentIndex++
		}

		iterationCount++
	}

	return stringBuilder.String()
}

// Rotate the Circular Buffer n times
func (cb *CircularBuffer) Rotate(n int) {
	cb.frontIndex = (cb.frontIndex - n + cb.size) % cb.size
}
