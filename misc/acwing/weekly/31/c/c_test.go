// Code generated by copypasta/template/acwing/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [c]")
	testCases := [][2]string{
		{
			`2 1 1
0 0
3 3
2 0 0
1 2`,
			`1`,
		},
		{
			`2 3 1
0 0
4 4
1 0 0
2 0 0
3 0 0
1 2`,
			`3`,
		},
	}
	target := 0 // -1
	testutil.AssertEqualStringCase(t, testCases, target, run)
}
