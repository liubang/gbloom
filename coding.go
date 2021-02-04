package gbloom

func DecodeFixed32(buffer []byte) uint32 {
	return uint32(buffer[0]) |
		(uint32(buffer[1]) << 8) |
		(uint32(buffer[2]) << 16) |
		(uint32(buffer[3]) << 24)
}

func DecodeFixed64(buffer []byte) uint64 {
	return uint64(buffer[0]) |
		(uint64(buffer[1]) << 8) |
		(uint64(buffer[2]) << 16) |
		(uint64(buffer[3]) << 24) |
		(uint64(buffer[4]) << 32) |
		(uint64(buffer[5]) << 40) |
		(uint64(buffer[6]) << 48) |
		(uint64(buffer[7]) << 56)
}
