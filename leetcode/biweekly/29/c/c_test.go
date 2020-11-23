// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [c]")
	examples := [][]string{
		{
			`[1,1,0,1]`, 
			`3`,
		},
		{
			`[0,1,1,1,0,1,1,0,1]`, 
			`5`,
		},
		{
			`[1,1,1]`, 
			`2`,
		},
		{
			`[1,1,0,0,1,1,1,0,1]`, 
			`4`,
		},
		{
			`[0,0,0]`, 
			`0`,
		},
		{
			`[1,0,0,0,0]`,
			`1`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, longestSubarray, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-29/problems/longest-subarray-of-1s-after-deleting-one-element/