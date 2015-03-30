package algorithms

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
        arr[i] = rand.Intn(100) - 100 / 2 // - range of -50..50
    }

    return arr
}

// http://golang.org/pkg/testing/
// http://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
func benchmarkSort(findMaxSubarraySum func([]int) int, arrSize int, b *testing.B) {
    arr := randomArray(arrSize)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        findMaxSubarraySum(arr)
    }
}

func BenchmarkMergeMSS10(b *testing.B) { benchmarkSort(MaxSubarraySum, 10, b) }
func BenchmarkMergeMSS100(b *testing.B) { benchmarkSort(MaxSubarraySum, 100, b) }
func BenchmarkMergeMSS1000(b *testing.B) { benchmarkSort(MaxSubarraySum, 1000, b) }
func BenchmarkMergeMSS10000(b *testing.B) { benchmarkSort(MaxSubarraySum, 10000, b) }
func BenchmarkMergeMSS100000(b *testing.B) { benchmarkSort(MaxSubarraySum, 100000, b) }

func BenchmarkMergeMSSK10(b *testing.B) { benchmarkSort(MaxSubarraySumKadane, 10, b) }
func BenchmarkMergeMSSK100(b *testing.B) { benchmarkSort(MaxSubarraySumKadane, 100, b) }
func BenchmarkMergeMSSK1000(b *testing.B) { benchmarkSort(MaxSubarraySumKadane, 1000, b) }
func BenchmarkMergeMSSK10000(b *testing.B) { benchmarkSort(MaxSubarraySumKadane, 10000, b) }
func BenchmarkMergeMSSK100000(b *testing.B) { benchmarkSort(MaxSubarraySumKadane, 100000, b) }
