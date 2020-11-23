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
			`4`, `[[2,1],[3,1],[1,4]]`, `2`, 
			`3`,
		},
		{
			`5`, `[[2,1],[3,1],[4,1],[1,5]]`, `2`, 
			`4`,
		},
		{
			`11`, `[]`, `2`, 
			`6`,
		},
		
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, minNumberOfSemesters, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/biweekly-contest-29/problems/parallel-courses-ii/