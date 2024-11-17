/*
Write a Go function that takes two strings as input and finds the longest common substring that occurs in both strings.
The function should return the substring that occurs earlier in the first string passed if there are
multiple substrings of the same length.
Expected function signature
func LongestCommonSubstr(str1, str2 string) string {

}
Usage
LongestCommonSubstr("ABCBDAB", "BDCAB")
>> "AB"
*/

package sprint

func LongestCommonSubstr(str1, str2 string) string {
	n, m := len(str1), len(str2)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	maxLength := 0
	endIndex := 0

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1

				if dp[i][j] > maxLength {
					maxLength = dp[i][j]
					endIndex = i
				}
			}
		}
	}

	if maxLength == 0 {
		return ""
	}
	return str1[endIndex-maxLength : endIndex]
}
