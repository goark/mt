package mt

import (
	"encoding/binary"
	"errors"
	"io"
	"sync"
	"testing"
)

func TestNil(t *testing.T) {
	prng := (*PRNG)(nil)
	// prng.Seed(0)
	prng.SeedArray(nil)
	if prng.Uint64() != 0 {
		t.Errorf("PRNG.Uint64() = %v, want %v.", prng.Uint64(), 0)
	}
	if prng.Real(0) != 0.0 {
		t.Errorf("PRNG.Real() = %v, want %v.", prng.Real(0), 0.0)
	}
}

// mockup for test
type testSource struct{}

func (t *testSource) SeedArray(seeds []uint64) {}
func (t *testSource) Uint64() uint64           { return 123456 }
func (t *testSource) Real(mode int) float64    { return 0.123456 }

func TestPRNG(t *testing.T) {
	prng := New(&testSource{})
	prng.SeedArray(nil) //no panic
	if prng.Uint64() != 123456 {
		t.Errorf("PRNG.Uint64() = %v, want %v.", prng.Uint64(), 123456)
	}
	if prng.Real(0) != 0.123456 {
		t.Errorf("PRNG.Real() = %v, want %v.", prng.Real(0), 0.123456)
	}
}

func getBytes(prng *PRNG) (uint64, error) {
	r := prng.NewReader()
	buf := [9]byte{}
	_, err := r.Read(buf[:])
	if err != nil {
		return 0, err
	}
	//fmt.Println(buf[:ct])
	_, err = r.ReadByte()
	if err != nil {
		return 0, err
	}
	ct, err := r.Read(buf[:])
	if err != nil {
		return 0, err
	}
	//fmt.Println(buf[:ct])
	return binary.LittleEndian.Uint64(buf[:ct]), nil
}

func TestReader(t *testing.T) {
	prng := New(&testSource{})
	wg := sync.WaitGroup{}
	res := binary.LittleEndian.Uint64([]byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x40, 0xe2, 0x01, 0x00})
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(id int) {
			n, err := getBytes(prng)
			if err != nil {
				t.Errorf("PRNG.Read() is %v, want nil.", err)
			}
			if n != res {
				t.Errorf("PRNG.Read() = %v, want %v.", n, res)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func TestReaderNil(t *testing.T) {
	r := (*PRNG)(nil).NewReader()
	buf := [8]byte{}
	_, err := r.Read(buf[:])
	if !errors.Is(err, io.ErrUnexpectedEOF) {
		t.Errorf("PRNG.Read() is \"%v\", want \"%v\".", err, io.ErrUnexpectedEOF)
	}
}

func TestNilReader(t *testing.T) {
	r := (*Reader)(nil)
	buf := [8]byte{}
	_, err := r.Read(buf[:])
	if !errors.Is(err, io.ErrUnexpectedEOF) {
		t.Errorf("PRNG.Read() is \"%v\", want \"%v\".", err, io.ErrUnexpectedEOF)
	}
}

/* MIT License
 *
 * Copyright 2019-2024 Spiegel
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
