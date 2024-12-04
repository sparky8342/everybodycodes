package utils

func Abs(n int) int {
	if n < 0 {
		return n * -1
	} else {
		return n
	}
}

func Min(nums []int) int {
	n := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < n {
			n = nums[i]
		}
	}
	return n
}

func Max(nums []int) int {
	n := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > n {
			n = nums[i]
		}
	}
	return n
}
