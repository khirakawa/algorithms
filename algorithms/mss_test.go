package algorithms

import "testing"

var tests = []struct{
    arr []int
    want int
}{
    {[]int{1, -3, 2, -5, 7, 6, -1, -4, 11, -23}, 19,},
    {[]int{-1, -3, -2, -5, -7, -23, 0}, 0,},
    {[]int{-4, -3, -8, -1}, -1,},
    {[]int{4, 8, 20}, 32,},
}

// http://golang.org/pkg/testing/
func TestMaxSubarraySum(t *testing.T) {
    for _, c := range tests {
        got := MaxSubarraySum(c.arr)
        if got != c.want {
            t.Errorf("MaxSubarraySum(%v) == %v, wanted %v", c.arr, got, c.want)
        }
    }
}

func TestMaxSubarraySumKadane(t *testing.T) {
    for _, c := range tests {
        got := MaxSubarraySumKadane(c.arr)
        if got != c.want {
            t.Errorf("MaxSubarraySumKadane(%v) == %v, wanted %v", c.arr, got, c.want)
        }
    }
}

