package mt

import (
	"context"
	"io"
	"math/rand"
	"sync"
)

//Source represents a source of uniformly-distributed
type Source interface {
	rand.Source64
	SeedArray([]uint64)
	Real(int) float64
}

// PRNG is class of pseudo random number generator.
type PRNG struct {
	source Source
	mutexR *sync.Mutex
	mutexW *sync.Mutex
	readCh <-chan byte
	cancel context.CancelFunc
}

var _ io.ReadCloser = (*PRNG)(nil) //PRNG is compatible with io.ReadCloser interface
var _ io.ByteReader = (*PRNG)(nil) //PRNG is compatible with io.ByteReader interface

//New returns new PRNG instance
func New(s Source) *PRNG {
	return &PRNG{source: s, mutexR: &sync.Mutex{}, mutexW: &sync.Mutex{}, readCh: nil, cancel: nil}
}

//Seed initializes Source.mt with a seed
func (prng *PRNG) Seed(seed int64) {
	if prng == nil {
		return
	}
	prng.mutexW.Lock()
	prng.source.Seed(seed)
	prng.mutexW.Unlock()
}

//SeedArray initializes Source.mt with seeds array
func (prng *PRNG) SeedArray(seeds []uint64) {
	if prng == nil {
		return
	}
	prng.mutexW.Lock()
	prng.source.SeedArray(seeds)
	prng.mutexW.Unlock()
}

//Uint64 generates a random number on [0, 2^64-1]-interval
func (prng *PRNG) Uint64() (n uint64) {
	if prng == nil {
		return
	}
	prng.mutexW.Lock()
	n = prng.source.Uint64()
	prng.mutexW.Unlock()
	return
}

//Int63 generates a random number on [0, 2^63-1]-interval
func (prng *PRNG) Int63() (n int64) {
	if prng == nil {
		return
	}
	prng.mutexW.Lock()
	n = prng.source.Int63()
	prng.mutexW.Unlock()
	return
}

//Real generates a random number
// on [0,1)-real-interval if mode==1,
// on (0,1)-real-interval if mode==2,
// on [0,1)-real-interval others
func (prng *PRNG) Real(mode int) (f float64) {
	if prng == nil {
		return
	}
	prng.mutexW.Lock()
	f = prng.source.Real(mode)
	prng.mutexW.Unlock()
	return
}

//Open triggers goroutine for generator.
func (prng *PRNG) Open() io.ReadCloser {
	if prng == nil {
		return prng
	}
	prng.mutexR.Lock()
	defer prng.mutexR.Unlock()
	if prng.cancel != nil {
		return prng
	}
	const sz = 8
	ch := make(chan byte, sz)
	prng.readCh = ch
	ctx, cancel := context.WithCancel(context.Background())
	prng.cancel = cancel
	go func() {
		defer close(ch)
		ct := sz
		buf := [sz]byte{}
		for {
			if ct >= sz {
				n := prng.Uint64()
				buf[0] = byte(n)
				buf[1] = byte(n >> 8)
				buf[2] = byte(n >> 16)
				buf[3] = byte(n >> 24)
				buf[4] = byte(n >> 32)
				buf[5] = byte(n >> 40)
				buf[6] = byte(n >> 48)
				buf[7] = byte(n >> 56)
				ct = 0
			}
			select {
			case <-ctx.Done():
				return
			default:
				ch <- buf[ct]
				ct++
			}
		}
	}()
	return prng
}

//ReadByte reads byte data from generator (compatible with io.ByteReader interface)
func (prng *PRNG) ReadByte() (byte, error) {
	if prng == nil || prng.readCh == nil {
		return 0, io.ErrUnexpectedEOF
	}
	b, ok := <-prng.readCh
	if !ok {
		return b, io.EOF
	}
	return b, nil
}

//Read reads bytes data from generator (compatible with io.ReadCloser interface)
func (prng *PRNG) Read(buf []byte) (int, error) {
	l := len(buf)
	if l == 0 {
		return 0, nil
	}
	prng.mutexR.Lock()
	defer prng.mutexR.Unlock()
	for i := 0; i < l; i++ {
		b, err := prng.ReadByte()
		if err != nil {
			return i, err
		}
		buf[i] = b
	}
	return l, nil
}

//Close closes goroutine for generator (compatible with io.ReadCloser interface).
//It always returns nil error.
func (prng *PRNG) Close() error {
	if prng == nil {
		return nil
	}
	prng.mutexR.Lock()
	defer prng.mutexR.Unlock()
	if prng.cancel != nil {
		prng.cancel()
		prng.cancel = nil
	}
	return nil
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
