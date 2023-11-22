package gotask_test

import (
	"gotask"
	"testing"

	"github.com/stretchr/testify/assert"
)

type EditingCase struct {
	name           string
	first, second  string
	expectedResult bool
}

func Test_Editing_MaxMismatch(t *testing.T) {
	cases := []EditingCase{
		{
			name:           "fist char empty, second is single char",
			first:          "",
			second:         "b",
			expectedResult: true,
		},
		{
			name:           "single char replace",
			first:          "a",
			second:         "b",
			expectedResult: true,
		},
		{
			name:           "one char remove",
			first:          "ab",
			second:         "b",
			expectedResult: true,
		},
		{
			name:           "char replace (string length 2)",
			first:          "ab",
			second:         "cb",
			expectedResult: true,
		},
		{
			name:           "char in beginning of string",
			first:          "xabc",
			second:         "abc",
			expectedResult: true,
		},
		{
			name:           "char in end of string",
			first:          "abcx",
			second:         "abc",
			expectedResult: true,
		},
		{
			name:           "char in middle of string",
			first:          "abc",
			second:         "ahc",
			expectedResult: true,
		},
		{
			name:           "wrong char order (string length 2)",
			first:          "ab",
			second:         "ba",
			expectedResult: false,
		},
		{
			name:           "need remove 2 char",
			first:          "aba",
			second:         "a",
			expectedResult: false,
		},
		{
			name:           "need replace 2 char",
			first:          "asdj",
			second:         "bsdk",
			expectedResult: false,
		},
	}

	for _, cs := range cases {
		t.Run(cs.name, func(t *testing.T) {
			testEditingCase(t, cs)
		})
	}
}

func testEditingCase(t *testing.T, editingCase EditingCase) {
	result := gotask.Editing(editingCase.first, editingCase.second, 1)
	assert.Equal(t, result, editingCase.expectedResult)
}
