// Package go-bloom-filter provides a Go implementation of a Bloom filter with counting.
package main

import (
	"errors"
	"fmt"
	"math/bits"
)

const (
	// Default seed value for the Mersenne Twister random number generator.
	defaultSeed = 31415
)

// BloomFilter is a Bloom filter with counting.
type BloomFilter struct {
	m       int
	n       int
	k       int
	bits    uint8
	counts  []uint8
	hashes  []func(string) uint32
	rng     *mt19937
}

// NewBloomFilter creates a new Bloom filter with counting.
func NewBloomFilter(m, n, k int) (*BloomFilter, error) {
	if m < 1 {
		return nil, errors.New("m must be greater than 0")
	}
	if n < 1 {
		return nil, errors.New("n must be greater than 0")
	}
	if k < 1 {
		return nil, errors.New("k must be greater than 0")
	}
	if m < n {
		return nil, errors.New("m must be greater than or equal to n")
	}

	bitsPerElement := uint(bits.Len(uint(m)))

	b := make([]uint8, (m+7)/8)
	counts := make([]uint8, n)
	hashes := make([]func(string) uint32, k)
	rng := newMersenneTwister(defaultSeed)

	for i := range hashes {
		hashes[i] = func(s string) (h uint32) {
			for _, b := range s {
				h ^= uint32(b)
			}
			return
		}
	}

	return &BloomFilter{
		m:       m,
		n:       n,
		k:       k,
		bits:    b,
		counts:  counts,
		hashes:  hashes,
		rng:     rng,
	}, nil
}

// Add adds an item to the Bloom filter.
func (bf *BloomFilter) Add(s string) {
	for i := range bf.hashes {
		h := bf.hashes[i](s)
		bitIndex := h % uint32(len(bf.bits))
		bf.bits[bitIndex] |= 1 << uint(i)
		bf.counts[h%uint32(len(bf.counts))]++
	}
}

// Maybe contains an item in the Bloom filter.
func (bf *BloomFilter) Maybe(s string) bool {
	for i := range bf.hashes {
		h := bf.hashes[i](s)
		bitIndex := h % uint32(len(bf.bits))
		if bf.bits[bitIndex]&1<<uint(i) == 0 {
			return false
		}
	}
	return true
}

// Count returns the count of occurrences for an item in the Bloom filter.
func (bf *BloomFilter) Count(s string) uint8 {
	for i := range bf.hashes {
		h := bf.hashes[i](s)
		bitIndex := h % uint32(len(bf.bits))
		if bf.bits[bitIndex]&1<<uint(i) == 0 {
			return 0
		}
	}
	return bf.counts[h%uint32(len(bf.counts))]
}

// newMersenneTwister creates a new Mersenne Twister random number generator.
func newMersenneTwister(seed uint32) *mt19937 {
	var mt [625]uint32
	mt[0] = seed
	for i := 1; i < 624; i++ {
		mt[i] = (1812433253*(mt[i-1]^uint32(i)))&0xffffffff
	}
	return &mt19937{mt: mt}
}

type mt19937 struct {
	mt [625]uint32
}

// next generates the next random number in the sequence.
func (m *mt19937) next() uint32 {
	var x uint32
	x ^= m.mt[0] >> 11
	x ^= m.mt[0] << 7 & 2636928640
	x ^= m.mt[0] << 15 & 4022730752
	m.mt[0] = m.mt[0] * 26975351 + 1
	for i := 1; i < 624; i++ {
		m.mt[i-1] = m.mt[i]
	}
	m.mt[624] = x
	return m.mt[0]
}

func main() {
	// Create a new Bloom filter with counting.
	bf, err := NewBloomFilter(128, 100, 3)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Add some items to the Bloom filter.
	bf.Add("hello")
	bf.Add("world")
	bf.Add("hello")

	// Check if an item is present in the Bloom filter.
	if bf.Maybe("hello") {
		fmt.Println("hello is present")
	} else {
		fmt.Println("hello is not present")
	}

	// Get the count of occurrences for an item in the Bloom filter.
	fmt.Println("Count for hello:", bf.Count("hello"))

	// Check if an item is not present in the Bloom filter.
	if !bf.Maybe("foo") {
		fmt.Println("foo is not present")
	}
}