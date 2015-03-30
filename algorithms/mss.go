package algorithms

func findMaxSubarrayMidSum(left, right []int) int {
    leftLen := len(left)
    leftMax := left[leftLen - 1]
    leftSum := leftMax
    for i := leftLen - 2; i >= 0; i-- {
        leftSum += left[i]
        if (leftSum > leftMax) {
            leftMax = leftSum
        }
    }

    rightLen := len(right)
    rightMax := right[0]
    rightSum := rightMax
    for i := 1; i < rightLen; i++ {
        rightSum += right[i]
        if (rightSum > rightMax) {
            rightMax = rightSum
        }
    }

    return leftMax + rightMax
}

// Finds Max Sub-array Sum
// https://www.youtube.com/watch?v=ohHWQf1HDfU
//
// Analysis:
// Time complexity is O(nlog(n))
// Space complexity is O(log(n)) extra memory for the use of stack (from recursion)
//
// Notes:
// - This uses divide and conquer, but you can also brute force this (O(n^2)), or use Kadane's algorithm for O(n)
func MaxSubarraySum(arr []int) int {
    if len(arr) == 1 {
        return arr[0]
    }

    mid := len(arr) / 2
    left := append(arr[:mid])
    right := append(arr[mid:])

    maxLeft := MaxSubarraySum(left)
    maxRight := MaxSubarraySum(right)
    maxMid := findMaxSubarrayMidSum(left, right)

    maxSum := maxLeft
    if maxRight > maxSum {
        maxSum = maxRight
    }
    if maxMid > maxSum {
        maxSum = maxMid
    }

    return maxSum
}

// Finds Max Sub-array Sum using Kadane Algorithm
// https://www.youtube.com/watch?v=ohHWQf1HDfU
//
// Analysis:
// Time complexity is O(n)
// Space complexity is O(1)
//
// Notes:
// - This uses divide and conquer, but you can also brute force this (O(n^2)), or use Kadane's algorithm for O(n)
func MaxSubarraySumKadane(arr []int) int {
    sum, max, i := 0, 0, 0

    // Check if all values are negative
    for ; i < len(arr); i++ {
        v := arr[i]
        if v > 0 {
            break
        }

        if i == 0 {
            max = v
        } else if (v > max) {
            max = v
        }
    }

    if i == len(arr) {
        return max
    }

    max = 0
    for i := 0; i < len(arr); i++ {
        sum += arr[i]
        if sum < 0 {
            sum = 0
        } else if sum > max {
            max = sum
        }
    }

    return max
}
