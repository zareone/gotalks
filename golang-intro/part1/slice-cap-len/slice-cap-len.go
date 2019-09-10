package main

import (
	"fmt"
)

func main() {
	// START OMIT
	nums := []string{"one", "two", "three"}
	fmt.Printf("nums: %v\n", nums)
	fmt.Printf("cap: %d, len: %d\n", cap(nums), len(nums))

	nums = append(nums, "four") // HL
	fmt.Printf("nums: %v\n", nums)
	fmt.Printf("cap: %d, len: %d\n", cap(nums), len(nums))

	fmt.Printf("first two: %v\n", nums[:2]) // nums[0:2]
	fmt.Printf("middle two: %v\n", nums[1:3])
	fmt.Printf("last two: %v\n", nums[2:]) // nums[2:len(nums)]
	// END OMIT
}
