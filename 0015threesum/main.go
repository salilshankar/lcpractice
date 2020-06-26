func threeSum(nums []int) [][]int {
    var ans [][]int
    if len(nums) == 3 && (nums[0]+nums[1]+nums[2]==0) {
        ans = append(ans, nums)
        return ans
    }
    sort.Ints(nums)
    low := 0
    high := 0
    sum := 0
    end := len(nums) - 2 
    for i:= 0; i<end; i++ {
        if (i == 0) || (nums[i] != nums[i-1]) {  
            low = i + 1
            high = len(nums) - 1
            sum = 0-nums[i]
            
            for low < high {
                if nums[low] + nums[high] == sum {
                    ans = append(ans, []int{nums[i], nums[low], nums[high]})
                    for (low < high) && (nums[low] == nums[low+1]) {
                        low++
                    }
                    for (low < high) && (nums[high-1] == nums[high]) {
                        high--
                    }
                    low++
                    high--
                } else if (nums[low] + nums[high] > sum) {
                    high--
                } else {
                    low++
                }
            }
        }
    }
    return ans   
}