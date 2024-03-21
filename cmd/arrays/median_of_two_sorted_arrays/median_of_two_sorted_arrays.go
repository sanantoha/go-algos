package main

import (
	"fmt"
	"math"
)

// O(n + m) time | O(n + m) space
func medianOfTwoSortedArrays(arr1, arr2 []int) float64 {
	if arr1 == nil || len(arr1) == 0 {
		return median(arr2)
	}
	if arr2 == nil || len(arr2) == 0 {
		return median(arr1)
	}

	arr := merge(arr1, arr2)
	return median(arr)
}

func merge(arr1, arr2 []int) []int {
	res := make([]int, len(arr1)+len(arr2))

	idx, i, j := 0, 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			res[idx] = arr1[i]
			idx++
			i++
		} else {
			res[idx] = arr2[j]
			idx++
			j++
		}
	}

	for i < len(arr1) {
		res[idx] = arr1[i]
		i++
		idx++
	}

	for j < len(arr2) {
		res[idx] = arr2[j]
		j++
		idx++
	}

	return res
}

func median(arr []int) float64 {
	if arr == nil || len(arr) == 0 {
		return math.NaN()
	}
	n := len(arr)
	h := n / 2
	if n%2 == 0 {
		return float64(arr[h]+arr[h-1]) / 2
	} else {
		return float64(arr[h])
	}
}

// O(log(min(n, m))) time | O(1) space
func medianOfTwoSortedArrays1(arr1, arr2 []int) float64 {
	if arr1 == nil || len(arr1) == 0 {
		return median(arr2)
	}
	if arr2 == nil || len(arr2) == 0 {
		return median(arr1)
	}

	len1 := len(arr1)
	len2 := len(arr2)
	if len1 > len2 {
		return medianOfTwoSortedArrays1(arr2, arr1)
	}

	l := 0
	r := len1

	for l <= r {
		med1 := (l + r) / 2
		med2 := (len1+len2+1)/2 - med1

		leftMax1 := math.MinInt
		if med1 != 0 {
			leftMax1 = arr1[med1-1]
		}
		rightMin1 := math.MaxInt
		if med1 != len1 {
			rightMin1 = arr1[med1]
		}

		leftMax2 := math.MinInt
		if med2 != 0 {
			leftMax2 = arr2[med2-1]
		}
		rightMin2 := math.MaxInt
		if med2 != len2 {
			rightMin2 = arr2[med2]
		}

		if leftMax1 <= rightMin2 && leftMax2 <= rightMin1 {
			ln := len1 + len2
			if ln%2 == 0 {
				return float64(maxInt(leftMax1, leftMax2)+minInt(rightMin1, rightMin2)) / 2
			} else {
				return float64(maxInt(leftMax1, leftMax2))
			}
		} else if leftMax1 > rightMin2 {
			r = med1 - 1
		} else {
			l = med1 + 1
		}
	}
	return math.NaN()
}

func minInt(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {

	arr1 := []int{1, 2, 3, 4, 5, 6}
	arr2 := []int{7, 8, 9, 10, 11, 12}

	fmt.Println(medianOfTwoSortedArrays(arr1, arr2))
	fmt.Println(medianOfTwoSortedArrays1(arr1, arr2))
}
