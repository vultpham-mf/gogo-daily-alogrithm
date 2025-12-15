package leetcode

/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".



Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.


Constraints:

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lowercase English letters if it is non-empty.
*/

// This is kinda slow
func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	longestPrefix := strs[0]
	for i := 1; i < len(strs); i++ {
		temp := ""
		for j := 0; j < len(longestPrefix) && j < len(strs[i]); j++ {
			if longestPrefix[j] == strs[i][j] {
				temp += string(longestPrefix[j])
			} else {
				break
			}
		}

		longestPrefix = temp
	}

	return longestPrefix
}

// This is faster
func longestCommonPrefix2(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	longestPrefix := ""
	for i := range strs[0] {
		currentLetter := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) == i {
				return longestPrefix
			}
			if strs[j][i] != currentLetter {
				return longestPrefix
			}
		}

		longestPrefix += string(currentLetter)
	}

	return longestPrefix
}
