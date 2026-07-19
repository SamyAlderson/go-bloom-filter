// Package main provides the main entry point for the Go Bloom filter implementation.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/stretchr/testify/assert"
)

// Main is the entry point for the application.
func Main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go-bloom-filter <file>")
		return
	}

	fileName := os.Args[1]

	bf := NewBloomFilter(10, 0.01)
	bf.Add([]byte("hello"))
	bf.Add([]byte("world"))
	bf.Add([]byte("foo"))

	countBf := NewCountBloomFilter(10, 0.01)
	countBf.Add([]byte("hello"), 1)
	countBf.Add([]byte("world"), 2)
	countBf.Add([]byte("foo"), 3)

	// Test the counting Bloom filter
	testCountBloomFilter(countBf)

	// Test the standard Bloom filter
	testBloomFilter(bf)

	// Load the input file
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		b := []byte(line)
		if bf.Test(b) {
			fmt.Println(line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	Main()
}

// testBloomFilter tests the standard Bloom filter.
func testBloomFilter(bf *BloomFilter) {
	// Test that the Bloom filter returns true for known elements
	assert.True(bf.Test([]byte("hello")))
	assert.True(bf.Test([]byte("world")))
	assert.True(bf.Test([]byte("foo")))

	// Test that the Bloom filter returns false for unknown elements
	assert.False(bf.Test([]byte("bar")))
}

// testCountBloomFilter tests the counting Bloom filter.
func testCountBloomFilter(countBf *CountBloomFilter) {
	// Test that the counting Bloom filter returns the correct count for known elements
	assert.Equal(countBf.Count([]byte("hello")), 1)
	assert.Equal(countBf.Count([]byte("world")), 2)
	assert.Equal(countBf.Count([]byte("foo")), 3)
}