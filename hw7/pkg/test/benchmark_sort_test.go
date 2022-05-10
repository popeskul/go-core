package test

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

var table = []int{100, 1160, 74382, 382399}

func BenchmarkSingleSortInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		nums := rand.Perm(1000)
		b.StartTimer()

		sort.Ints(nums)
	}
}

func BenchmarkSortInt(b *testing.B) {
	for _, n := range table {
		b.Run(fmt.Sprintf("sort_table_int_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				nums := rand.Perm(n)
				b.StartTimer()

				sort.Ints(nums)
			}
		})
	}
}

func BenchmarkSortFloat64s(b *testing.B) {
	for _, n := range table {
		b.Run(fmt.Sprintf("sort_table_float64s_%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				nums := randFloats(n)
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
