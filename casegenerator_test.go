package gotask_test

import (
	"fmt"
	"math/rand"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func GenerateCasesForMaxOperationCount(min, max int) []*EditingCase {
	cases := make([]*EditingCase, 0)
	caseGenerator := CaseGenerator{falseExpectResultCaseCount: 5, trueExpectResultCaseCount: 5}
	for maxOperationsCount := min; maxOperationsCount <= max; maxOperationsCount++ {
		cases = append(cases, caseGenerator.GenerateCases(maxOperationsCount)...)
	}
	return cases
}

type CaseGenerator struct {
	falseExpectResultCaseCount, trueExpectResultCaseCount int
}

func (c CaseGenerator) GenerateCases(maxOperationsCount int) []*EditingCase {
	cases := make([]*EditingCase, 0, c.falseExpectResultCaseCount+c.trueExpectResultCaseCount)
	cases = append(cases, c.generateFalseExpectResultCases(maxOperationsCount)...)
	cases = append(cases, c.generateTrueExpectResultCases(maxOperationsCount)...)
	return cases
}

func (c CaseGenerator) generateFalseExpectResultCases(maxOperationsCount int) []*EditingCase {
	cases := make([]*EditingCase, 0, c.falseExpectResultCaseCount)
	for i := 0; i < c.falseExpectResultCaseCount; i++ {
		cases = append(cases, generateFalseExpectResultCase(maxOperationsCount))
	}
	return cases
}

func (c CaseGenerator) generateTrueExpectResultCases(maxOperationsCount int) []*EditingCase {
	cases := make([]*EditingCase, 0, c.trueExpectResultCaseCount)
	for i := 0; i < c.falseExpectResultCaseCount; i++ {
		cases = append(cases, generateTrueExpectResultCase(maxOperationsCount))
	}
	return cases
}

func generateFalseExpectResultCase(maxOperationsCount int) *EditingCase {
	runeListSize := maxOperationsCount * 3
	runeList := randomRuneList(runeListSize)
	opList := generateRandomStringEditOperationsList(maxOperationsCount*2, runeListSize)
	return &EditingCase{
		Name:               opList.Name(),
		First:              string(runeList),
		Second:             opList.Apply(runeList),
		ExpectedResult:     false,
		MaxOperationsCount: maxOperationsCount,
	}
}

func generateTrueExpectResultCase(maxOperationsCount int) *EditingCase {
	runeListSize := maxOperationsCount * 3
	runeList := randomRuneList(runeListSize)
	opList := generateRandomStringEditOperationsList(maxOperationsCount, runeListSize)

	return &EditingCase{
		Name:               opList.Name(),
		First:              string(runeList),
		Second:             opList.Apply(runeList),
		ExpectedResult:     true,
		MaxOperationsCount: maxOperationsCount,
	}
}

func generateRandomStringEditOperationsList(maxOperationsCount int, runeListSize int) *StringEditOperationList {
	operations := make([]StringEditOperation, 0, maxOperationsCount)
	indexSet := NewIndexSet(runeListSize)
	operationGenerator := OperationGenerator{indexSet: indexSet}
	for i := 0; i < maxOperationsCount; i++ {
		operation := operationGenerator.RandomOperation()
		operations = append(operations, operation)
	}
	return NewStringEditOperationList(operations)
}

type OperationGenerator struct {
	indexSet IndexSet
}

func NewOperationGenerator(indexSet IndexSet) OperationGenerator {
	return OperationGenerator{
		indexSet: indexSet,
	}
}

func (o *OperationGenerator) RandomOperation() StringEditOperation {
	randomValue := rand.Int()
	switch randomValue % 2 {
	case 1:
		return o.insertOperation()
	case 0:
		return o.replaceOperation()
	}
	return nil
}

func (o *OperationGenerator) containMaxIndex() bool {
	maxIndex := o.indexSet.maxValue - 1
	_, ok := o.indexSet.indexes[maxIndex]
	return ok
}

func (o *OperationGenerator) removeOperation() RemoveOperation {
	index := o.indexSet.NextIndex()
	o.indexSet.DescreaseIndexes(index)
	return RemoveOperation{
		Index: index,
	}
}

func (o *OperationGenerator) replaceOperation() ReplaceOperation {
	index := o.indexSet.NextIndex()
	return ReplaceOperation{
		Index: index,
		Rune:  randomRune(),
	}
}

func (o *OperationGenerator) insertOperation() InsertOperation {
	index := o.indexSet.NextIndex()
	o.indexSet.IncreaseIndexes(index)
	return InsertOperation{
		Index: index,
		Rune:  randomRune(),
	}
}

type IndexSet struct {
	maxValue int
	indexes  map[int]struct{}
}

func NewIndexSet(maxValue int) IndexSet {
	return IndexSet{maxValue: maxValue, indexes: make(map[int]struct{})}
}

func (i *IndexSet) DescreaseIndexes(index int) {
	i.maxValue--
	for _, key := range i.filterKeysBigger(index) {
		delete(i.indexes, key)
		if key > 0 {
			i.indexes[key-1] = struct{}{}
		}
	}
}

func (i *IndexSet) IncreaseIndexes(index int) {
	i.maxValue++
	for _, key := range i.filterKeysBigger(index) {
		delete(i.indexes, key)
		i.indexes[key+1] = struct{}{}
	}
	i.indexes[index] = struct{}{}
}

func (i *IndexSet) filterKeysBigger(index int) []int {
	keys := make([]int, 0)
	for key := range i.indexes {
		if key > index {
			keys = append(keys, key)
		}
	}
	sort.Sort(sort.IntSlice(keys))
	return keys
}

func (i *IndexSet) NextIndex() int {
	var index int
	for {
		index = rand.Intn(i.maxValue)
		_, ok := i.indexes[index]
		if !ok {
			break
		}
	}
	i.indexes[index] = struct{}{}
	return index
}

var runeSet []rune = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

func randomRuneList(maxLen int) []rune {
	runeSlice := make([]rune, 0, maxLen)
	for i := 0; i < maxLen; i++ {
		runeSlice = append(runeSlice, randomRune())
	}
	return runeSlice
}

func randomRune() rune {
	randomIndex := rand.Intn(len(runeSet))
	return runeSet[randomIndex]
}

type StringEditOperationList struct {
	builder    strings.Builder
	operations []StringEditOperation
}

func NewStringEditOperationList(operations []StringEditOperation) *StringEditOperationList {
	return &StringEditOperationList{operations: operations}
}

func (s *StringEditOperationList) Name() string {
	if s.builder.Len() != 0 {
		return s.builder.String()
	}
	for index, op := range s.operations {
		s.builder.WriteString(op.String())
		if index < len(s.operations)-1 {
			s.builder.WriteString(", ")
		}
	}
	return s.builder.String()
}

func (s *StringEditOperationList) Apply(runeList []rune) string {
	appliedRuneList := make([]rune, len(runeList))
	copy(appliedRuneList, runeList)
	for _, op := range s.operations {
		appliedRuneList = op.Apply(appliedRuneList)
	}
	return string(appliedRuneList)
}

type StringEditOperation interface {
	Apply([]rune) []rune
	String() string
}

type InsertOperation struct {
	Index int
	Rune  rune
}

func (o InsertOperation) Apply(runeList []rune) []rune {
	return slices.Insert(runeList, o.Index, o.Rune)
}

func (o InsertOperation) String() string {
	return fmt.Sprintf("insert into %d index, %s rune", o.Index, strconv.QuoteRune(o.Rune))
}

type ReplaceOperation struct {
	Index int
	Rune  rune
}

func (o ReplaceOperation) Apply(runeList []rune) []rune {
	for runeList[o.Index] == o.Rune {
		o.Rune = randomRune()
	}
	runeList[o.Index] = o.Rune
	return runeList
}

func (o ReplaceOperation) String() string {
	return fmt.Sprintf("replace %d index %s rune", o.Index, strconv.QuoteRune(o.Rune))
}

type RemoveOperation struct {
	Index int
}

func (o RemoveOperation) Apply(runeList []rune) []rune {
	return slices.Delete(runeList, o.Index, o.Index+1)
}

func (o RemoveOperation) String() string {
	return fmt.Sprintf("remove %d index", o.Index)
}
