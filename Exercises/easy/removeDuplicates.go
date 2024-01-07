func removeDuplicates(nums []int) int {

	var count = 0

	for  i := 0; i < len(nums) - 1 - count; i++ {
		if nums[i] != nums[i+1] {
			count = count + 1
		} else {
			for j := i+1; j < len(nums) - 1; j++ {
				nums[j] = nums[j+1]
			}
		}
	}

	return count

}