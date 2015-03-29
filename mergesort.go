package main

func merge(left, right []int) []int {
    result := make([]int, len(left) + len(right))
    i, j := 0, 0

    for i < len(left) && j < len(right) {
        leftValue := left[i]
        rightValue := right[j]

        if leftValue <= rightValue {
            result[i + j] = leftValue
            i++
        } else {
            result[i + j] = rightValue
            j++
        }
    }

    for i < len(left) {
        result[i + j] = left[i]
        i++
    }

    for j < len(right) {
        result[i + j] = right[j]
        j++
    }

    return result
}

// http://www.sorting-algorithms.com/merge-sort
// https://www.youtube.com/watch?v=0nlPxaC2lTw
// Time complexity is O(nlog(n))
// Space complexity is O(n) extra memory. This is because we free the auxiliary slices on each stack pop
// However, if we don't free memory at all, then it would be O(nlog(n))
func Mergesort(arr []int) []int {
    if (len(arr) == 1) {
        return arr
    }

    mid := len(arr) / 2
    left := append(arr[:mid])
    right := append(arr[mid:])

    return merge(Mergesort(left), Mergesort(right))
}
