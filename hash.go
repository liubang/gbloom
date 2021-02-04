package gbloom

const (
	m uint32 = 0xc6a4a793
	r uint32 = 24
)

func Hash(data []byte, n uint32, seed uint32) uint32 {
	h := seed ^ (n * m)
	var idx uint32
	for {
		if idx+4 > n {
			break
		}
		w := DecodeFixed32(data[idx:])
		idx += 4
		h += w
		h *= m
		h ^= (h >> 16)
	}
	left := n - idx
	switch left {
	case 3:
		h += uint32(data[idx+2]) << 16
		h += uint32(data[idx+1]) << 8
		h += uint32(data[idx+0]) << 0
	case 2:
		h += uint32(data[idx+1]) << 8
		h += uint32(data[idx+0]) << 0
	case 1:
		h += uint32(data[idx+0]) << 0
	}
	if left != 0 {
		h *= m
		h ^= (h >> r)
	}
	return h
}
