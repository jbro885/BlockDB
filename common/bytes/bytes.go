package bytes

// GetUInt32 get an uint32 from byte array with a start position.
// This is for those little-endian bytes.
func GetUInt32(b []byte, pos int) uint32 {
	return uint32(b[pos]) | uint32(b[pos+1])<<8 | uint32(b[pos+2])<<16 | uint32(b[pos+3])<<24
}

// SetUInt32 set an uint32 into byte array at a position.
func SetUInt32(b []byte, pos int, i uint32) {
	b[pos] = byte(i)
	b[pos+1] = byte(i >> 8)
	b[pos+2] = byte(i >> 16)
	b[pos+3] = byte(i >> 24)
}

// GetInt32 get an int32 from byte array with a start position.
func GetInt32(b []byte, pos int) int32 {
	return int32(b[pos]) | int32(b[pos+1])<<8 | int32(b[pos+2])<<16 | int32(b[pos+3])<<24
}

// SetInt32 set an int32 into byte array at a position.
func SetInt32(b []byte, pos int, i int32) {
	b[pos] = byte(i)
	b[pos+1] = byte(i >> 8)
	b[pos+2] = byte(i >> 16)
	b[pos+3] = byte(i >> 24)
}

// GetInt64 get an int64 from byte array with a start position.
func GetInt64(b []byte, pos int) int64 {
	return int64(b[pos]) | int64(b[pos+1])<<8 | int64(b[pos+2])<<16 | int64(b[pos+3])<<24 |
		int64(b[pos+4])<<32 | int64(b[pos+5])<<40 | int64(b[pos+6])<<48 | int64(b[pos+7])<<56
}
