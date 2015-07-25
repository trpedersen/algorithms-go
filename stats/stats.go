package stats

import (
	"fmt"
)

type Accumulator struct {
	total float64
	N     int
}

func NewAccumulator() *Accumulator {
	return &Accumulator{}
}

func (this *Accumulator) AddFloat(value float64) {
	this.total += value
	this.N++
}

func (this *Accumulator) AddInt64(value int64) {
	this.AddFloat(float64(value))
}

func (this *Accumulator) Mean() float64 {
	return this.total / float64(this.N)
}

func (this *Accumulator) String() string {
	return fmt.Sprintf("Mean(%d values): %f", this.N, this.Mean())
}

type Int64Accumulator struct {
	total int64
	N     int64
}

func NewInt64Accumulator() *Int64Accumulator {
	return &Int64Accumulator{}
}

func (this *Int64Accumulator) Add(value int64) {
	this.total += value
	this.N++
}

func (this *Int64Accumulator) Mean() int64 {
	return this.total / this.N
}

func (this *Int64Accumulator) String() string {
	return fmt.Sprintf("Total: %d N: %d Mean: %d", this.total, this.N, this.Mean())
}
