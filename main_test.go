package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	input := "helloWorld"

	circularBuffer := New(&input)

	result := circularBuffer.String()

	assert.Equal(t, result, input)
}

type RotateTest struct {
	Input  string
	Rotate int
	Expect string
	Name   string
}

var rotateTests = []RotateTest{
	{
		Input:  "helloworld",
		Rotate: 0,
		Expect: "helloworld",
	},
	{
		Input:  "helloworld",
		Rotate: 1,
		Expect: "dhelloworl",
	},
	{
		Input:  "helloworld",
		Rotate: 2,
		Expect: "ldhellowor",
	},
}

func TestRotate(t *testing.T) {
	for _, test := range rotateTests {
		t.Run(test.Name, func(t *testing.T) {
			circularBuffer := New(&test.Input)

			circularBuffer.Rotate(test.Rotate)

			response := circularBuffer.String()

			assert.Equal(t, test.Expect, response)

		})
	}
}
