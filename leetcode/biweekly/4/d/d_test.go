// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	examples := [][]string{
		{
			`[1,2,2,3,3,4,4]`, `3`, 
			`true`,
		},
		{
			`[5,6,6,7,8]`, `3`, 
			`false`,
		},
		// TODO 测试入参最小的情况
		
	}
	targetCaseNum := 0 // -1
	if err := testutil.RunLeetCodeFuncWithExamples(t, canDivideIntoSubsequences, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-4/problems/divide-array-into-increasing-sequences/