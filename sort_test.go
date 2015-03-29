package main

import (
    "testing"
    "reflect"
    "math/rand"
    "time"
)

// http://golang.org/pkg/testing/
func TestSort(t *testing.T) {
    var tests = []struct{
        arr, want []int
    }{
        {
            []int{2, 6, 345, 54, 23, 34, 3, 3, 434, 432, 1},
            []int{1, 2, 3, 3, 6, 23, 34, 54, 345, 432, 434},
        },
    }

    for _, c := range tests {
        got := Mergesort(c.arr)
        if !reflect.DeepEqual(got, c.want) {
            t.Errorf("Mergesort(%v) == %v, wanted %v", c.arr, got, c.want)
        }
    }
}

func randomArray(size int) []int {
    arr := make([]int, size)
    t := time.Now()
    rand.Seed(int64(t.Nanosecond()))

    for i:= 0; i < size; i++ {
        arr[i] = rand.Intn(100)
    }

    return arr
}

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
