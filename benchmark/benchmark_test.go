package benchmark

import (
	"math/rand/v2"
	"testing"

	"github.com/goark/mt/v2"
	"github.com/goark/mt/v2/mt19937"
)

var seed1, seed2, seed3 = rand.Uint64(), rand.Uint64(), rand.Int64()

func BenchmarkRandomPCG(b *testing.B) {
	rnd := rand.NewPCG(seed1, seed2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = rnd.Uint64()
	}
}

func BenchmarkRandomMT19917(b *testing.B) {
	rnd := mt19937.New(seed3)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = rnd.Uint64()
	}
}

func BenchmarkRandomPCGRand(b *testing.B) {
	rnd := rand.New(rand.NewPCG(seed1, seed2))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = rnd.Uint64()
	}
}

func BenchmarkRandomMT19917Rand(b *testing.B) {
	rnd := rand.New(mt19937.New(seed3))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = rnd.Uint64()
	}
}

func BenchmarkRandomChaCha8Locked(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = rand.Uint64()
	}
}

func BenchmarkRandomMT19917Locked(b *testing.B) {
	rnd := mt.New(mt19937.New(seed3))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = rnd.Uint64()
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
