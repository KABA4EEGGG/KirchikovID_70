package Sum

func Sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func Sum_slice(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
