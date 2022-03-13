package secure

import (
	"crypto/rand"
	"encoding/binary"

	"github.com/goark/mt"
)

// Source is random source for mt.Source interface.
type Source struct{}

var _ mt.Source = (*Source)(nil) //Source is compatible with mt.Source interface

// Seed method is dummy function for rand.Source interface.
func (s Source) Seed(seed int64) {}

// SeedArray method is dummy function for mt.Source interface.
func (s Source) SeedArray([]uint64) {}

// Read method generates random bytes, using crypto/rand.Read() function.
func (s Source) Read(buf []byte) (int, error) {
	return rand.Read(buf)
}

// Uint64 method generates a random number in the range [0, 1<<64).
func (s Source) Uint64() uint64 {
	b := [8]byte{}
	ct, _ := s.Read(b[:])
	return binary.BigEndian.Uint64(b[:ct])
}

// Int63 method generates a random number in the range [0, 1<<63).
func (s Source) Int63() int64 {
	return (int64)(s.Uint64() >> 1)
}

//Real generates a random number
// on [0,1)-real-interval if mode==1,
// on (0,1)-real-interval if mode==2,
// on [0,1]-real-interval others
func (s Source) Real(mode int) float64 {
	switch mode {
	case 1: //generates a random number on [0,1)-real-interval
		return (float64)(s.Uint64()>>11) * (1.0 / 9007199254740991.0)
	case 2: //generates a random number on (0,1)-real-interval
		return (float64)(s.Uint64()>>11) * (1.0 / 9007199254740992.0)
	default: //generates a random number on [0,1]-real-interval
		return ((float64)(s.Uint64()>>12) + 0.5) * (1.0 / 4503599627370496.0)
	}
}

/* MIT License
 *
 * Copyright 2021 Spiegel
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */
