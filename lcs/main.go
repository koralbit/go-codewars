package main

import (
	"fmt"
)

/*
Longest CommonWrite a function called LCS that accepts two sequences and returns the longest subsequence common to the passed in sequences.

Subsequence

A subsequence is different from a substring. The terms of a subsequence need not be consecutive terms of the original sequence.

Example subsequence
Subsequences of "abc" = "a", "b", "c", "ab", "ac", "bc" and "abc".

LCS( "abcdef", "abc" ) => returns "abc"
LCS( "abcdef", "acf" ) => returns "acf"
LCS( "132535365", "123456789" ) => returns "12356"

REF: https://www.codewars.com/kata/52756e5ad454534f220001ef/train/go
*/
func main()  {
	fmt.Println(LCS( "132535365", "123456789" ))
	fmt.Println(LCS( "abcdef", "abc" ))

}
func LCS(s1, s2 string) string {
	r, c := len(s1), len(s2)
	m := make([][]string, r + 1)
	for i := range m {
		m[i] = make([]string, c + 1)
	}
	for i := 1 ; i <=  r ; i ++{
		rv := string(s1[i - 1])
		for j := 1 ; j <=  c ; j ++ {
			cv := string(s2[j - 1])
			if(rv == cv ) {
				m[i][j] =  m[i-1][j-1] + s1[i-1:i]
				continue
			} 
			if len(m[i-1][j]) > len(m[i][j-1]) {
				m[i][j] = m[i-1][j]
			} else {
				m[i][j] = m[i][j-1]
			}
		}
	}
	return m[r][c]
}