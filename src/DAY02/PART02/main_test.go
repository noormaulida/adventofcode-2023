package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInputDay2Part1(t *testing.T) {

	t.Run("Test Input Day 2 Part 2", func(t *testing.T) {
		fileName := "../_testcase/input.txt"
		expected := 71585
		actual := Calculate(fileName)
		assert.Equal(t, expected, actual)
	})

}
