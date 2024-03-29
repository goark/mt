//go:build run
// +build run

package main

import (
	"fmt"
	"math"
	"math/rand/v2"

	"github.com/goark/mt/v2/mt19937"
)

func main() {
	rnd := rand.New(mt19937.New(rand.Int64()))
	points := []float64{}
	max := 0.0
	min := 1.0
	sum := 0.0
	for range 10000 {
		point := rnd.NormFloat64()
		points = append(points, point)
		min = math.Min(min, point)
		max = math.Max(max, point)
		sum += point
	}
	n := float64(len(points))
	ave := sum / n
	d2 := 0.0
	for _, p := range points {
		d2 += (p - ave) * (p - ave)
	}
	fmt.Println("           minimum: ", min)
	fmt.Println("           maximum: ", max)
	fmt.Println("           average: ", ave)
	fmt.Println("standard deviation: ", math.Sqrt(d2/n))
}

/* MIT License
 *
 * Copyright 2024 Spiegel
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
