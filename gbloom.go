package gbloom

import (
	"math"
)

func BloomHash(key []byte) uint32 {
	return Hash(key, uint32(len(key)), 0xbc9f1d34)
}

type BloomFilter struct {
	bits_per_key uint8
	k            uint8
}

func NewBloomFilter(bits_per_key uint8) *BloomFilter {
	k := uint8(math.Floor(float64(bits_per_key) * 0.69))
	if k < 1 {
		k = 1
	} else if k > 30 {
		k = 30
	}
	return &BloomFilter{
		bits_per_key: bits_per_key,
		k:            k,
	}
}

func (bloom *BloomFilter) Name() string {
	return "gleveldb.BuiltinBloomFilter2"
}

func (bloom *BloomFilter) CreateFilter(keys [][]byte) []byte {
	n := len(keys)
	bits := uint32(n) * uint32(bloom.bits_per_key)
	if bits < 64 {
		bits = 64
	}
	bytes := (bits + 7) / 8
	bits = bytes * 8
	new_arr := make([]byte, bytes+1)
	new_arr[bytes] = byte(bloom.k)
	for i := 0; i < n; i++ {
		h := BloomHash(keys[i])
		delta := (h >> 17) | (h << 15)
		for j := 0; uint8(j) < bloom.k; j++ {
			bitpos := h % uint32(bits)
			new_arr[bitpos/8] |= (1 << (bitpos % 8))
			h += delta
		}
	}
	return new_arr
}

func (bloom *BloomFilter) KeyMayMatch(bloom_filter, key []byte) bool {
	length := len(bloom_filter)
	if length < 2 {
		return false
	}
	bits := uint32((length - 1) * 8)
	k := int32(bloom_filter[length-1])
	if k > 30 {
		return true
	}
	h := BloomHash(key)
	delta := (h >> 17) | (h << 15)
	for j := 0; int32(j) < k; j++ {
		bitpos := h % bits
		if bloom_filter[bitpos/8]&(1<<(bitpos%8)) == 0 {
			return false
		}
		h += delta
	}
	return true
}
