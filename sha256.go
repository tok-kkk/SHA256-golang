package sha256

import (
	"encoding/binary"
)

var BlockSize = 64

// bit
// byte = 8 bits
// word = 32 bits

type Digest struct {
	h      [8]uint32
	chunks [][64]byte
}

func NewDigest(data []byte) Digest {
	var chunks [][64]byte

	switch len(data) {
	case 64:
		chunks = make([][64]byte, 1)
		copy(chunks[0][:], data[:])
	case 128:
		chunks = make([][64]byte, 2)
		copy(chunks[0][:], data[:64])
		copy(chunks[1][:], data[64:])
	}

	return Digest{
		h:      [8]uint32{h0, h1, h2, h3, h4, h5, h6, h7},
		chunks: chunks,
	}
}

func (digest *Digest) Hash() [32]byte {
	for _, chunk := range digest.chunks {

		a := digest.h[0]
		b := digest.h[1]
		c := digest.h[2]
		d := digest.h[3]
		e := digest.h[4]
		f := digest.h[5]
		g := digest.h[6]
		h := digest.h[7]

		var w [64]uint32
		for i := 0; i < len(chunk); i = i + 4 {
			word := binary.BigEndian.Uint32(chunk[i : i+4])
			w[i/4] = word
		}
		for i := 16; i < 64; i++ {
			s0 := omg0(w[i-15])
			s1 := omg1(w[i-2])
			w[i] = w[i-16] + s0 + w[i-7] + s1
		}
		for i := 0; i < 64; i++ {
			s1 := sig1(e)
			ch := ch(e, f, g)
			temp1 := h + s1 + ch + Rounds[i] + w[i]

			s0 := sig0(a)
			maj := ma(a, b, c)
			temp2 := s0 + maj

			h = g
			g = f
			f = e
			e = d + temp1
			d = c
			c = b
			b = a
			a = temp1 + temp2
		}

		digest.h[0] = digest.h[0] + a
		digest.h[1] = digest.h[1] + b
		digest.h[2] = digest.h[2] + c
		digest.h[3] = digest.h[3] + d
		digest.h[4] = digest.h[4] + e
		digest.h[5] = digest.h[5] + f
		digest.h[6] = digest.h[6] + g
		digest.h[7] = digest.h[7] + h
	}

	hash := make([]byte, 0)
	for i := range digest.h {
		bytes := make([]byte, 4)
		binary.BigEndian.PutUint32(bytes, digest.h[i])
		hash = append(hash, bytes...)
	}
	var hash32 [32]byte
	copy(hash32[:], hash)
	return hash32
}

// After pre-processing we got either 512 or 1024 bits
func PreProcessing(data []byte) []byte {
	length := len(data) * 8
	data = append(data, byte(128))
	for len(data)%BlockSize != 56 {
		data = append(data, byte(0))
	}
	lengthBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(lengthBytes, uint64(length))
	data = append(data, lengthBytes...)

	return data
}

func Sha256(input []byte) [32]byte {
	data := PreProcessing(input)
	digest := NewDigest(data)
	return digest.Hash()
}
