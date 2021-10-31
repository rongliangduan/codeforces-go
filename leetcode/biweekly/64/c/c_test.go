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
			`"**|**|***|"`, `[[2,5],[5,9]]`, 
			`[2,3]`,
		},
		{
			`"***|**|*****|**||**|*"`, `[[1,17],[4,5],[14,17],[5,11],[15,16]]`, 
			`[9,0,0,0,0]`,
		},
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, platesBetweenCandles, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-64/problems/plates-between-candles/
