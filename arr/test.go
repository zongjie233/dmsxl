package arr

func test(nums []int, target int) int {
	i := 0
	length := len(nums)
	res := length + 1
	sum := 0
	for j := 0; j < length; j++ {
		sum += nums[j]
		for sum >= target {
			sublength := j - i + 1
			if sublength < res {
				res = sublength
			}
			sum -= nums[i]
			i++
		}
	}
	if res == length+1 {
		return 0
	} else {
		return res
	}
}
