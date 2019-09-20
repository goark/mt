package mt19937

import (
	"math/rand"
)

const (
	nn        = 312
	mm        = 156
	matrixA   = 0xB5026F5AA96619E9
	upperMask = 0xFFFFFFFF80000000 //Most significant 33 bits
	lowerMask = 0x000000007FFFFFFF //Least significant 31 bits
)

//Source is a source of random numbers.
type Source struct {
	mt  [nn]uint64 //The array for the state vector
	mti int        //mti==nn+1 means mt[nn] is not initialized
}

//New returns a new pseudo-random source seeded with the given value.
func New(seed int64) *Source {
	rng := &Source{mt: [nn]uint64{}, mti: nn + 1}
	rng.Seed(seed)
	return rng
}

//NewWithArray returns a new pseudo-random source seeded with the given values.
func NewWithArray(seeds []uint64) *Source {
	rng := &Source{mt: [nn]uint64{}, mti: nn + 1}
	rng.SeedArray(seeds)
	return rng
}

//NewSource returns a new pseudo-random source (compatible with rand.Source interface)
func NewSource(seed int64) rand.Source {
	return New(seed)
}

//Seed initializes Source.mt with a seed
func (s *Source) Seed(seed int64) {
	if s == nil {
		return
	}
	s.mt[0] = uint64(seed)
	for s.mti = 1; s.mti < nn; s.mti++ {
		s.mt[s.mti] = 6364136223846793005*(s.mt[s.mti-1]^(s.mt[s.mti-1]>>62)) + uint64(s.mti)
	}
}

//SeedArray initializes Source.mt with seeds array
func (s *Source) SeedArray(seeds []uint64) {
	if s == nil {
		return
	}
	s.Seed(19650218)
	k := len(seeds)
	if k == 0 {
		return
	}
	if nn > k {
		k = nn
	}
	i := 1
	j := 0
	for ; k > 0; k-- {
		s.mt[i] = (s.mt[i] ^ ((s.mt[i-1] ^ (s.mt[i-1] >> 62)) * 3935559000370003845)) + seeds[j] + uint64(j) // non linear
		i++
		if i >= nn {
			s.mt[0] = s.mt[nn-1]
			i = 1
		}
		j++
		if j >= len(seeds) {
			j = 0
		}
	}
	for k = nn - 1; k > 0; k-- {
		s.mt[i] = (s.mt[i] ^ ((s.mt[i-1] ^ (s.mt[i-1] >> 62)) * 2862933555777941757)) - uint64(i) // non linear
		i++
		if i >= nn {
			s.mt[0] = s.mt[nn-1]
			i = 1
		}
	}
	s.mt[0] = 1 << 63 //MSB is 1; assuring non-zero initial array
}

//Uint64 generates a random number on [0, 2^64-1]-interval
func (s *Source) Uint64() uint64 {
	if s == nil {
		return 0
	}
	if s.mti >= nn {
		if s.mti >= 1+nn {
			s.Seed(5489) // a default initial seed is used
		}
		for i := 0; i < nn-1; i++ {
			x := (s.mt[i] & upperMask) | (s.mt[i+1] & lowerMask)
			if i < (nn - mm) {
				s.mt[i] = s.mt[i+mm] ^ (x >> 1) ^ ((x & 0x01) * matrixA)
			} else {
				s.mt[i] = s.mt[i+(mm-nn)] ^ (x >> 1) ^ ((x & 0x01) * matrixA)
			}
		}
		x := (s.mt[nn-1] & upperMask) | (s.mt[0] & lowerMask)
		s.mt[nn-1] = s.mt[mm-1] ^ (x >> 1) ^ ((x & 0x01) * matrixA)
		s.mti = 0
	}

	x := s.mt[s.mti]
	s.mti++
	x ^= (x >> 29) & 0x5555555555555555
	x ^= (x << 17) & 0x71D67FFFEDA60000
	x ^= (x << 37) & 0xFFF7EEE000000000
	x ^= (x >> 43)
	return x
}

//Int63 generates a random number on [0, 2^63-1]-interval
func (s *Source) Int63() int64 {
	return (int64)(s.Uint64() >> 1)
}

/* MIT License
 *
 * Copyright 2019 Spiegel
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
