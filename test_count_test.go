package count

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountingBloomFilter(t *testing.T) {
	// Create a new counting Bloom filter
	filter := NewCountingBloomFilter(10, 0.1)

	// Add some elements
	filter.Add("hello")
	filter.Add("world")

	// Check if the elements are present
	assert.True(t, filter.Contains("hello"), "hello should be present")
	assert.True(t, filter.Contains("world"), "world should be present")

	// Check if the elements are not present after reset
	filter.Reset()
	assert.False(t, filter.Contains("hello"), "hello should not be present after reset")
	assert.False(t, filter.Contains("world"), "world should not be present after reset")
}

func TestCountingBloomFilterOverflow(t *testing.T) {
	// Create a new counting Bloom filter with a small capacity
	filter := NewCountingBloomFilter(2, 0.1)

	// Add more elements than the capacity
	filter.Add("hello")
	filter.Add("world")
	filter.Add("foo")
	filter.Add("bar")

	// Check if the filter has overflowed
	assert.Error(t, filter.Contains("hello"), "hello should cause an error due to overflow")
	assert.Error(t, filter.Contains("world"), "world should cause an error due to overflow")
}

func TestCountingBloomFilterEmpty(t *testing.T) {
	// Create a new counting Bloom filter
	filter := NewCountingBloomFilter(10, 0.1)

	// Check if the filter is empty
	assert.True(t, filter.Empty(), "filter should be empty")

	// Add an element
	filter.Add("hello")

	// Check if the filter is not empty
	assert.False(t, filter.Empty(), "filter should not be empty")
}