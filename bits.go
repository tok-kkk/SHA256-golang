package sha256

import "math/bits"

func ch(x, y, z uint32) uint32 {
	return (x & y) ^ (^x & z)
}

func ma(x, y, z uint32) uint32 {
	return (x & y) ^ (x & z) ^ (y & z)
}

func sig0(x uint32) uint32 {
	return (bits.RotateLeft32(x, -2)) ^ (bits.RotateLeft32(x, -13)) ^ (bits.RotateLeft32(x, -22))
}

func sig1(x uint32) uint32 {
	return (bits.RotateLeft32(x, -6)) ^ (bits.RotateLeft32(x, -11)) ^ (bits.RotateLeft32(x, -25))
}

func omg0(x uint32) uint32 {
	return (bits.RotateLeft32(x, -7)) ^ (bits.RotateLeft32(x, -18)) ^ (x >> 3)
}

func omg1(x uint32) uint32 {
	return (bits.RotateLeft32(x, -17)) ^ (bits.RotateLeft32(x, -19)) ^ (x >> 10)
}
