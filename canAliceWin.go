package golangfirststeps

func canAliceWin(nums []int) bool {
    sum1 := 0
    sum2 := 0
    totalSum := 0

    for _, num := range nums {
        if num < 10 {
            sum1 += num
        } else {
            sum2 += num
        }
        totalSum += num
    }

    if len(nums) > 1 && len(nums) <= 100 {
        return sum1 > (totalSum - sum1) || sum2 > (totalSum - sum2)
    }
    return true
}

/*
You are given an array of positive integers nums.

Alice and Bob are playing a game. In the game, Alice can choose either all single-digit numbers or
all double-digit numbers from nums, and the rest of the numbers are given to Bob. Alice wins if
the sum of her numbers is strictly greater than the sum of Bob's numbers.

Return true if Alice can win this game, otherwise, return false.

Example 1:
	Input: nums = [1,2,3,4,10]
	Output: false

	Alice cannot win by choosing either single-digit or double-digit numbers.

Example 2:
	Input: nums = [1,2,3,4,5,14]
	Output: true

	Alice can win by choosing single-digit numbers which have a sum equal to 15.

Example 3:
	Input: nums = [5,5,5,25]
	Output: true

	Alice can win by choosing double-digit numbers which have a sum equal to 25.

Constraints:
	1 <= nums.length <= 100
	1 <= nums[i] <= 99
*/