package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInputD1P1(t *testing.T) {

	t.Run("Test Input Day 1 Part 1", func(t *testing.T) {
		fileName := "../_testcase/input.txt"
		expected := 55108
		actual := Calculate(fileName)
		assert.Equal(t, expected, actual)
	})

}
