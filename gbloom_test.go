package gbloom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmall(t *testing.T) {
	bloomFilter := NewBloomFilter(16 * 8)
	keys := [][]byte{
		[]byte("hello"),
		[]byte("world"),
	}
	filter := bloomFilter.CreateFilter(keys)
	assert.True(t, bloomFilter.KeyMayMatch(filter, []byte("hello")))
	assert.True(t, bloomFilter.KeyMayMatch(filter, []byte("world")))
	assert.False(t, bloomFilter.KeyMayMatch(filter, []byte("x")))
	assert.False(t, bloomFilter.KeyMayMatch(filter, []byte("foo")))
}

func TestBloomFilter(t *testing.T) {
	bloomFilter := NewBloomFilter(16 * 8)
	keys := [][]byte{
		[]byte("OatoXR1oWQ27aghr"),
		[]byte("FsuHfYSASnISaLXK"),
		[]byte("0u56dghh6A2VpoqW"),
		[]byte("$Q$0pjppVzm367Vr"),
		[]byte("5wA#jO9nuLEUBKWf"),
	}
	filter := bloomFilter.CreateFilter(keys)
	for _, bys := range keys {
		assert.True(t, bloomFilter.KeyMayMatch(filter, bys))
	}
}
