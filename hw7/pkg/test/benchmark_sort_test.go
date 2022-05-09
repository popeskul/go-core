package test

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

var table = []struct {
	n int
}{
	{n: 100},
	{n: 1160},
	{n: 74382},
	{n: 382399},
}

func BenchmarkSingleSortInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := rand.Perm(1000)
		b.StartTimer()

		sort.Ints(nums)
	}
}

func BenchmarkSortInt(b *testing.B) {
	for _, t := range table {
		b.Run(fmt.Sprintf("sort_table_int_%d", t.n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				nums := rand.Perm(t.n)
				b.StartTimer()

				sort.Ints(nums)
			}
		})
	}
}

func BenchmarkSortFloat64s(b *testing.B) {
	for _, t := range table {
		b.Run(fmt.Sprintf("sort_table_float64s_%d", t.n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				nums := randFloats(100)
				b.StartTimer()

				sort.Float64s(nums)
			}
		})
	}
}

func randFloats(size int) []float64 {
	res := make([]float64, size)

	for i := range res {
		res[i] = 0. + rand.Float64()*(float64(size)-0.)
	}
	return res
}
