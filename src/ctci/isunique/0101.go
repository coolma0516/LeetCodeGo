// Implement an algorithm to determine if a string has all unique characters. What if you cannot use additional data structures?

// Example 1:

// Input: s = "leetcode"
// Output: false
// Example 2:

// Input: s = "abc"
// Output: true
//

// Note:

// 0 <= len(s) <= 100
package isunique

func isUnique(astr string) bool {
	// assume that character are in a..z
	num := 0

	for _, v := range astr {
		moveBit := v - 'a'
		if num&(1<<moveBit) != 0 {
			return false
		}
		num = num | (1 << moveBit)

	}
	return true
}
