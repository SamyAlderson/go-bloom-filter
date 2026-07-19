package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBloom(t *testing.T) {
	bf := NewBloom(10, 0.01)
	assert.Nil(t, bf.Add("hello"))
	assert.Nil(t, bf.Add("world"))
	assert.Nil(t, bf.Add(""))

	// Check false positives
	assert.True(t, bf.MaybeContains("goodbye"))
	assert.True(t, bf.MaybeContains("foo"))
	assert.False(t, bf.MaybeContains("nope"))

	// Check false negatives
	assert.False(t, bf.MaybeContains("hello"))
	assert.False(t, bf.MaybeContains("world"))

	// Clean up
	assert.Nil(t, bf.Clear())

	// Check clear
	assert.False(t, bf.MaybeContains("hello"))
	assert.False(t, bf.MaybeContains("world"))
	assert.False(t, bf.MaybeContains("goodbye"))
}

func TestBloomFull(t *testing.T) {
	bf := NewBloom(100, 0.01)
	words := []string{
		"hello", "world", "foo", "bar", "baz", "qux", "quux", "corge", "grault", "garply",
		"waldo", "fred", "plugh", "thud", "xyzzy", "fjord", "kirk", "spark", "jolt", "tull",
		"glorp", "gurp", "splat", "squirt", "bloop", "blop", "bloop", "bleep", "bloop", "gurgle",
		"splish", "splash", "bloop", "bleep", "splat", "glorp", "fjord", "bloop", "blop", "blop",
	}

	for _, word := range words {
		assert.Nil(t, bf.Add(word))
	}

	for _, word := range words {
		assert.True(t, bf.MaybeContains(word))
	}

	// Try some words that aren't in the filter
	assert.False(t, bf.MaybeContains("goodbye"))
	assert.False(t, bf.MaybeContains("nope"))
}

func TestBloomError(t *testing.T) {
	_, err := NewBloom(-1, 0.01)
	assert.Error(t, err)

	_, err = NewBloom(10, -0.01)
	assert.Error(t, err)
}