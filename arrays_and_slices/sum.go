package arrays_and_slices


import (
	"fmt"
)


func AddNums(nums []int) int {
	fmt.Printf("input %d", nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
