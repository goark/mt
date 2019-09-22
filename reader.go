package mt

import (
	"io"
)

//Reader is class of pseudo random number generator with io.Reader interface.
type Reader struct {
	prng *PRNG
	rn   uint64
	pos  int8
}

var _ io.Reader = (*Reader)(nil)     //Reader is compatible with io.Reader interface
var _ io.ByteReader = (*Reader)(nil) //Reader is compatible with io.ByteReader interface

//Read reads bytes data from generator (compatible with io.ReadCloser interface)
func (r *Reader) Read(buf []byte) (int, error) {
	if r == nil || r.prng == nil {
		return 0, io.ErrUnexpectedEOF
	}
	rest := len(buf)
	if rest == 0 {
		return 0, io.ErrUnexpectedEOF
	}
	r.prng.mutex.Lock()
	defer r.prng.mutex.Unlock()

	ct := 0
	for r.pos > 0 {
		buf[ct] = r.getByte(false)
		rest--
		ct++
		if rest == 0 {
			return ct, nil
		}
	}
	rest8 := rest >> 3 //devision by 8
	for rest8 > 0 {
		rn := r.prng.source.Uint64()
		buf[ct] = byte(rn)
		buf[ct+1] = byte(rn >> 8)
		buf[ct+2] = byte(rn >> 16)
		buf[ct+3] = byte(rn >> 24)
		buf[ct+4] = byte(rn >> 32)
		buf[ct+5] = byte(rn >> 40)
		buf[ct+6] = byte(rn >> 48)
		buf[ct+7] = byte(rn >> 56)
		rest8--
		ct += 8
	}
	rest &= 0x7
	for rest > 0 {
		buf[ct] = r.getByte(true)
		rest--
		ct++
	}
	return ct, nil
}

//Read reads bytes data from generator (compatible with io.ReadCloser interface)
func (r *Reader) ReadByte() (byte, error) {
	buf := [1]byte{}
	_, err := r.Read(buf[:])
	return buf[0], err
}

func (r *Reader) getByte(upd bool) (b byte) {
	if r.pos == 0 {
		if !upd {
			return
		}
		r.rn = r.prng.source.Uint64()
		r.pos = 7
	}
	b = byte(r.rn)
	r.rn >>= 8
	r.pos--
	return
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
