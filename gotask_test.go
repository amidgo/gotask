package gotask_test

import (
	"fmt"
	"gotask"
	"testing"

	"github.com/stretchr/testify/assert"
)

type EditingCase struct {
	name               string
	first, second      string
	expectedResult     bool
	maxOperationsCount int
}

var operationCountEqual_1_Cases = []*EditingCase{
	{
		name:               "fist char empty, second is single char",
		first:              "",
		second:             "b",
		expectedResult:     true,
		maxOperationsCount: 1,
	},
	{
		name:               "single char replace",
		first:              "a",
		second:             "b",
		expectedResult:     true,
		maxOperationsCount: 1,
	},
	{
		name:               "one char remove",
		first:              "ab",
		second:             "b",
		expectedResult:     true,
		maxOperationsCount: 1,
	},
	{
		name:               "char replace (string length 2)",
		first:              "ab",
		second:             "cb",
		expectedResult:     true,
		maxOperationsCount: 1,
	},
	{
		name:               "char in beginning of string",
		first:              "xabc",
		second:             "abc",
		expectedResult:     true,
		maxOperationsCount: 1,
	},
	{
		name:               "char in end of string",
		first:              "abcx",
		second:             "abc",
		expectedResult:     true,
		maxOperationsCount: 1,
	},
	{
		name:               "char in middle of string",
		first:              "abc",
		second:             "ahc",
		expectedResult:     true,
		maxOperationsCount: 1,
	},
	{
		name:               "wrong char order (string length 2)",
		first:              "ab",
		second:             "ba",
		expectedResult:     false,
		maxOperationsCount: 1,
	},
	{
		name:               "2 different operations append to end and replace in middle",
		first:              "abcde",
		second:             "bbde",
		expectedResult:     false,
		maxOperationsCount: 1,
	},
	{
		name:               "2 different operations append to start and replace in middle",
		first:              "abcde",
		second:             "bcdd",
		expectedResult:     false,
		maxOperationsCount: 1,
	},
	{
		name:               "need remove 2 char",
		first:              "aba",
		second:             "a",
		expectedResult:     false,
		maxOperationsCount: 1,
	},
	{
		name:               "need replace 2 char",
		first:              "asdj",
		second:             "bsdk",
		expectedResult:     false,
		maxOperationsCount: 1,
	},
}

var operationCountEqual_2_Cases = []*EditingCase{
	{
		name:               "fist char empty, second is single char",
		first:              "",
		second:             "b",
		expectedResult:     true,
		maxOperationsCount: 1,
	},
	{
		name:               "single char replace",
		first:              "a",
		second:             "b",
		expectedResult:     true,
		maxOperationsCount: 1,
	},
	{
		name:               "one char remove",
		first:              "ab",
		second:             "b",
		expectedResult:     true,
		maxOperationsCount: 2,
	},
	{
		name:               "char replace (string length 2)",
		first:              "ab",
		second:             "cb",
		expectedResult:     true,
		maxOperationsCount: 2,
	},
	{
		name:               "char in beginning of string",
		first:              "xabc",
		second:             "abc",
		expectedResult:     true,
		maxOperationsCount: 2,
	},
	{
		name:               "char in end of string",
		first:              "abcx",
		second:             "abc",
		expectedResult:     true,
		maxOperationsCount: 2,
	},
	{
		name:               "char in middle of string",
		first:              "abc",
		second:             "ahc",
		expectedResult:     true,
		maxOperationsCount: 2,
	},
	{
		name:               "fist char empty, second is two chars",
		first:              "",
		second:             "ba",
		expectedResult:     true,
		maxOperationsCount: 2,
	},
	{
		name:               "need replace 2 chars",
		first:              "ba",
		second:             "ab",
		expectedResult:     true,
		maxOperationsCount: 2,
	},
	{
		name:               "need append 2 chars in end",
		first:              "a",
		second:             "abb",
		expectedResult:     true,
		maxOperationsCount: 2,
	},
	{
		name:               "need append 2 chars in start",
		first:              "a",
		second:             "bba",
		expectedResult:     true,
		maxOperationsCount: 2,
	},
	{
		name:               "need append chars in middle and end",
		first:              "abcde",
		second:             "acd",
		expectedResult:     true,
		maxOperationsCount: 2,
	},
	{
		name:               "2 different operations, replace first char and insert into end",
		first:              "a",
		second:             "bc",
		expectedResult:     true,
		maxOperationsCount: 2,
	},
	{
		name:               "2 different operations, append into start and replace in middle",
		first:              "abcdef",
		second:             "btdef",
		expectedResult:     true,
		maxOperationsCount: 2,
	},
	{
		name:               "2 different operations, replace in middle and append into end",
		first:              "abcdef",
		second:             "abtde",
		expectedResult:     true,
		maxOperationsCount: 2,
	},

	{
		name:               "2 different operations, replace in start and append into end",
		first:              "abcdef",
		second:             "bbcde",
		expectedResult:     true,
		maxOperationsCount: 2,
	},

	{
		name:               "2 different operations, append in start and replace into end",
		first:              "abcdef",
		second:             "bcdet",
		expectedResult:     true,
		maxOperationsCount: 2,
	},
	{
		name:               "3 different operations, replace and 2 append",
		first:              "a",
		second:             "bcd",
		expectedResult:     false,
		maxOperationsCount: 2,
	},

	{
		name:               "need replace 3 chars",
		first:              "abc",
		second:             "bca",
		expectedResult:     false,
		maxOperationsCount: 2,
	},

	{
		name:               "need append 3 chars",
		first:              "",
		second:             "abc",
		expectedResult:     false,
		maxOperationsCount: 2,
	},
}

func Test_Editing(t *testing.T) {
	cases := make([]*EditingCase, 0)
	cases = append(cases, operationCountEqual_1_Cases...)
	// cases = append(cases, operationCountEqual_2_Cases...)

	for _, cs := range cases {
		name := fmt.Sprintf("%s, max op count %d", cs.name, cs.maxOperationsCount)
		t.Run(name, func(t *testing.T) {
			testEditingCase(t, cs)
		})
	}
}

func testEditingCase(t *testing.T, editingCase *EditingCase) {
	result := gotask.Editing(editingCase.first, editingCase.second, editingCase.maxOperationsCount)
	assert.Equal(t, result, editingCase.expectedResult)
}
