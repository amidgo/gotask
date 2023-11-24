package gotask_test

import (
	"fmt"
	"gotask"
	"testing"

	"github.com/stretchr/testify/assert"
)

type EditingCase struct {
	Name               string
	First, Second      string
	ExpectedResult     bool
	MaxOperationsCount int
}

var operationCountEqual_1_Cases = []*EditingCase{
	{
		Name:               "fist char empty, second is single char",
		First:              "",
		Second:             "b",
		ExpectedResult:     true,
		MaxOperationsCount: 1,
	},
	{
		Name:               "single char replace",
		First:              "a",
		Second:             "b",
		ExpectedResult:     true,
		MaxOperationsCount: 1,
	},
	{
		Name:               "one char remove",
		First:              "ab",
		Second:             "b",
		ExpectedResult:     true,
		MaxOperationsCount: 1,
	},
	{
		Name:               "char replace (string length 2)",
		First:              "ab",
		Second:             "cb",
		ExpectedResult:     true,
		MaxOperationsCount: 1,
	},
	{
		Name:               "char in beginning of string",
		First:              "xabc",
		Second:             "abc",
		ExpectedResult:     true,
		MaxOperationsCount: 1,
	},
	{
		Name:               "char in end of string",
		First:              "abcx",
		Second:             "abc",
		ExpectedResult:     true,
		MaxOperationsCount: 1,
	},
	{
		Name:               "char in middle of string",
		First:              "abc",
		Second:             "ahc",
		ExpectedResult:     true,
		MaxOperationsCount: 1,
	},
	{
		Name:               "wrong char order (string length 2)",
		First:              "ab",
		Second:             "ba",
		ExpectedResult:     false,
		MaxOperationsCount: 1,
	},
	{
		Name:               "2 different operations append to end and replace in middle",
		First:              "abcde",
		Second:             "bbde",
		ExpectedResult:     false,
		MaxOperationsCount: 1,
	},
	{
		Name:               "2 different operations append to start and replace in middle",
		First:              "abcde",
		Second:             "bcdd",
		ExpectedResult:     false,
		MaxOperationsCount: 1,
	},
	{
		Name:               "need remove 2 char",
		First:              "aba",
		Second:             "a",
		ExpectedResult:     false,
		MaxOperationsCount: 1,
	},
	{
		Name:               "need replace 2 char",
		First:              "asdj",
		Second:             "bsdk",
		ExpectedResult:     false,
		MaxOperationsCount: 1,
	},
}

var operationCountEqual_2_Cases = []*EditingCase{
	{
		Name:               "fist char empty, second is single char",
		First:              "",
		Second:             "b",
		ExpectedResult:     true,
		MaxOperationsCount: 1,
	},
	{
		Name:               "single char replace",
		First:              "a",
		Second:             "b",
		ExpectedResult:     true,
		MaxOperationsCount: 1,
	},
	{
		Name:               "one char remove",
		First:              "ab",
		Second:             "b",
		ExpectedResult:     true,
		MaxOperationsCount: 2,
	},
	{
		Name:               "char replace (string length 2)",
		First:              "ab",
		Second:             "cb",
		ExpectedResult:     true,
		MaxOperationsCount: 2,
	},
	{
		Name:               "char in beginning of string",
		First:              "xabc",
		Second:             "abc",
		ExpectedResult:     true,
		MaxOperationsCount: 2,
	},
	{
		Name:               "char in end of string",
		First:              "abcx",
		Second:             "abc",
		ExpectedResult:     true,
		MaxOperationsCount: 2,
	},
	{
		Name:               "char in middle of string",
		First:              "abc",
		Second:             "ahc",
		ExpectedResult:     true,
		MaxOperationsCount: 2,
	},
	{
		Name:               "fist char empty, second is two chars",
		First:              "",
		Second:             "ba",
		ExpectedResult:     true,
		MaxOperationsCount: 2,
	},
	{
		Name:               "need replace 2 chars",
		First:              "ba",
		Second:             "ab",
		ExpectedResult:     true,
		MaxOperationsCount: 2,
	},
	{
		Name:               "need append 2 chars in end",
		First:              "a",
		Second:             "abb",
		ExpectedResult:     true,
		MaxOperationsCount: 2,
	},
	{
		Name:               "need append 2 chars in start",
		First:              "a",
		Second:             "bba",
		ExpectedResult:     true,
		MaxOperationsCount: 2,
	},
	{
		Name:               "need append chars in middle and end",
		First:              "abcde",
		Second:             "acd",
		ExpectedResult:     true,
		MaxOperationsCount: 2,
	},
	{
		Name:               "2 different operations, replace first char and insert into end",
		First:              "a",
		Second:             "bc",
		ExpectedResult:     true,
		MaxOperationsCount: 2,
	},
	{
		Name:               "2 different operations, append into start and replace in middle",
		First:              "abcdef",
		Second:             "btdef",
		ExpectedResult:     true,
		MaxOperationsCount: 2,
	},
	{
		Name:               "2 different operations, replace in middle and append into end",
		First:              "abcdef",
		Second:             "abtde",
		ExpectedResult:     true,
		MaxOperationsCount: 2,
	},

	{
		Name:               "2 different operations, replace in start and append into end",
		First:              "abcdef",
		Second:             "bbcde",
		ExpectedResult:     true,
		MaxOperationsCount: 2,
	},

	{
		Name:               "2 different operations, append in start and replace into end",
		First:              "abcdef",
		Second:             "bcdet",
		ExpectedResult:     true,
		MaxOperationsCount: 2,
	},
	{
		Name:               "3 different operations, replace and 2 append",
		First:              "a",
		Second:             "bcd",
		ExpectedResult:     false,
		MaxOperationsCount: 2,
	},

	{
		Name:               "need replace 3 chars",
		First:              "abc",
		Second:             "bca",
		ExpectedResult:     false,
		MaxOperationsCount: 2,
	},

	{
		Name:               "need append 3 chars",
		First:              "",
		Second:             "abc",
		ExpectedResult:     false,
		MaxOperationsCount: 2,
	},
}

func Test_Editing(t *testing.T) {
	cases := make([]*EditingCase, 0)
	cases = append(cases, operationCountEqual_1_Cases...)
	cases = append(cases, operationCountEqual_2_Cases...)
	cases = append(cases, GenerateCasesForMaxOperationCount(3, 20)...)

	for _, cs := range cases {
		name := fmt.Sprintf("%s, max op count %d", cs.Name, cs.MaxOperationsCount)
		t.Run(name, func(t *testing.T) {
			testEditingCase(t, cs)
		})
	}
}

func testEditingCase(t *testing.T, editingCase *EditingCase) {
	result := gotask.Editing(editingCase.First, editingCase.Second, editingCase.MaxOperationsCount)
	assert.Equal(t, editingCase.ExpectedResult, result, "first string: %s, second string: %s", editingCase.First, editingCase.Second)
}
