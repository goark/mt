package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/spiegel-im-spiegel/mt"
	"github.com/spiegel-im-spiegel/mt/mt19937"
)

func main() {
	start := time.Now()
	prng := mt.New(mt19937.NewSource(19650218))
	wg := sync.WaitGroup{}
	//r := prng.NewReader()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			r := prng.NewReader()
			for i := 0; i < 10000; i++ {
				buf := [8]byte{}
				//ct, err := r.Read(buf[:])
				_, err := r.Read(buf[:])
				if err != nil {
					return
				}
				//fmt.Println(binary.LittleEndian.Uint64(buf[:ct]))
			}
		}()
	}
	wg.Wait()
	fmt.Println("Time:", time.Now().Sub(start))
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
