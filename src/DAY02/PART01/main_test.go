package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInputD2P1(t *testing.T) {

	t.Run("Test Input Day 2 Part 1", func(t *testing.T) {
		fileName := "../_testcase/input.txt"
		expected := 2331
		actual := Calculate(fileName)
		assert.Equal(t, expected, actual)
	})

}
