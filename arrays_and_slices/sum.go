package arrays_and_slices

func AddNums(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var res []int
	for _, num := range numbersToSum {
		res = append(res, AddNums(num))
	}
	return res

}

func SumAllTails(nums ...[]int) []int {
	var res []int
	for _, num := range nums {
		if len(num) == 0 {
			res = append(res, 0)
		} else {
			tail := num[1:]
			res = append(res, AddNums(tail))
		}
	}
	return res
}
