package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInputD1P2(t *testing.T) {

	t.Run("Test Input Day 1 Part 2", func(t *testing.T) {
		fileName := "../_testcase/input.txt"
		expected := 56324
		actual := Calculate(fileName)
		assert.Equal(t, expected, actual)
	})

}
