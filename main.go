package main

import (
    "fmt"
    "github.com/khirakawa/algorithms/algorithms"
)

func main() {
    sum := algorithms.MaxSubarraySum([]int{3, -2, 5, -1})
    fmt.Printf("max sum is %d\n", sum)
}
