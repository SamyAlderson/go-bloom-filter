package main

import (
	"errors"
	"fmt"
	"math/bits"
)

type CountingBloomFilter struct {
	m       []uint8
	k       int
	m       int
	n       int
	estimatedN int
}

func NewCountingBloomFilter(k, m, n int) (*CountingBloomFilter, error) {
	if k <= 0 || m <= 0 {
		return nil, errors.New("k and m must be positive integers")
	}
	if n <= 0 {
		return nil, errors.New("n must be a positive integer")
	}
	if k > bits.Len64(uint64(n)) {
		return nil, errors.New("k cannot be larger than the number of bits required to represent n")
	}
	if m > n {
		return nil, errors.New("m cannot be larger than n")
	}
	cf := &CountingBloomFilter{
		k:       k,
		m:       m,
		n:       n,
	}
	cf.estimatedN = 0
	cf.m = make([]uint8, m)
	return cf, nil
}

func (cf *CountingBloomFilter) Add(key []byte) {
	hash := make([]uint64, cf.k)
	for i := range hash {
		hash[i] = fnv1aHash(key)
	}
	for i, h := range hash {
		index := int(h % uint64(cf.n))
		cf.m[index]++
		if cf.m[index] == 1 {
			cf.estimatedN++
		}
	}
}

func (cf *CountingBloomFilter) MightContain(key []byte) bool {
	hash := make([]uint64, cf.k)
	for i := range hash {
		hash[i] = fnv1aHash(key)
	}
	for _, h := range hash {
		index := int(h % uint64(cf.n))
		if cf.m[index] == 0 {
			return false
		}
	}
	return true
}

func (cf *CountingBloomFilter) EstimatedN() int {
	return cf.estimatedN
}

func fnv1aHash(key []byte) uint64 {
	h := uint64(2166136261)
	for _, b := range key {
		h = (h ^ uint64(b)) * 16777219
	}
	return h
}

func main() {
	// Example usage
	bf, err := NewCountingBloomFilter(10, 100, 1000)
	if err != nil {
		fmt.Println(err)
		return
	}
	bf.Add([]byte("hello"))
	fmt.Println(bf.MightContain([]byte("hello"))) // true
	fmt.Println(bf.MightContain([]byte("world"))) // might be true or false
	fmt.Println(bf.EstimatedN()) // returns the estimated number of unique elements
}