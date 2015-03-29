package sort

import (
    "testing"
    "math/rand"
    "time"
)

func randomArray(size int) []int {
    arr := make([]int, size)
    t := time.Now()
    rand.Seed(int64(t.Nanosecond()))

    for i:= 0; i < size; i++ {
        arr[i] = rand.Intn(100)
    }

    return arr
}

// http://golang.org/pkg/testing/
// http://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
func benchmarkSort(sort func([]int) []int, arrSize int, b *testing.B) {
    arr := randomArray(arrSize)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        sort(arr)
    }
}

func BenchmarkMergeSort10(b *testing.B) { benchmarkSort(Mergesort, 10, b) }
func BenchmarkMergeSort100(b *testing.B) { benchmarkSort(Mergesort, 100, b) }
func BenchmarkMergeSort1000(b *testing.B) { benchmarkSort(Mergesort, 1000, b) }
func BenchmarkMergeSort10000(b *testing.B) { benchmarkSort(Mergesort, 10000, b) }
func BenchmarkMergeSort100000(b *testing.B) { benchmarkSort(Mergesort, 100000, b) }
