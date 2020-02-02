package testcase

func SumSlice(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum = sum + v
	}
	return sum
}

func SumArray(nums [10]int) int {
	sum := 0
	for _, v := range nums {
		sum = sum + v
	}
	return sum
}

func SumArrayP(nums *[10]int) int {
	sum := 0
	for _, v := range nums {
		sum = sum + v
	}
	return sum
}
