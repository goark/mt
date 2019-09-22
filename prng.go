package mt

import (
	"math/rand"
	"sync"
)

//Source represents a source of uniformly-distributed
type Source interface {
	rand.Source64
	SeedArray([]uint64)
	Real(int) float64
}

//PRNG is class of pseudo random number generator.
type PRNG struct {
	source Source
	mutex  *sync.Mutex
}

var _ rand.Source64 = (*PRNG)(nil) //PRNG is compatible with rand.Source and rand.Source64 interface

//New returns new PRNG instance
func New(s Source) *PRNG {
	return &PRNG{source: s, mutex: &sync.Mutex{}}
}

//Seed initializes Source.mt with a seed
func (prng *PRNG) Seed(seed int64) {
	if prng == nil {
		return
	}
	prng.mutex.Lock()
	prng.source.Seed(seed)
	prng.mutex.Unlock()
}

//SeedArray initializes Source.mt with seeds array
func (prng *PRNG) SeedArray(seeds []uint64) {
	if prng == nil {
		return
	}
	prng.mutex.Lock()
	prng.source.SeedArray(seeds)
	prng.mutex.Unlock()
}

//Uint64 generates a random number on [0, 2^64-1]-interval
func (prng *PRNG) Uint64() (n uint64) {
	if prng == nil {
		return 0
	}
	prng.mutex.Lock()
	n = prng.source.Uint64()
	prng.mutex.Unlock()
	return
}

//Int63 generates a random number on [0, 2^63-1]-interval
func (prng *PRNG) Int63() (n int64) {
	if prng == nil {
		return 0
	}
	prng.mutex.Lock()
	n = prng.source.Int63()
	prng.mutex.Unlock()
	return
}

//Real generates a random number
// on [0,1)-real-interval if mode==1,
// on (0,1)-real-interval if mode==2,
// on [0,1]-real-interval others
func (prng *PRNG) Real(mode int) (f float64) {
	if prng == nil {
		return 0
	}
	prng.mutex.Lock()
	f = prng.source.Real(mode)
	prng.mutex.Unlock()
	return
}

//NewReader returns new Reader instance.
func (prng *PRNG) NewReader() *Reader {
	return &Reader{prng: prng}
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
