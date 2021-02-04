package gbloom

import (
	"fmt"
	"testing"
)

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
		res := bloomFilter.KeyMayMatch(filter, bys)
		if !res {
			t.Errorf("the key %s match result is error, result: %v", string(bys), res)
		}
	}

	keys_not_in_filters := [][]byte{
		[]byte("rpshHpE3CPInGPeb"),
		[]byte("T40#TEoxh*3DA*Ad"),
	}
	for _, bys := range keys_not_in_filters {
		res := bloomFilter.KeyMayMatch(filter, bys)
		fmt.Printf("key:%s, res:%v\n", string(bys), res)
	}
}
