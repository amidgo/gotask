package gotask

// Написать метод (класс и импорты не нужны) на вход которого приходит две строки.
// На выходе надо проверить можно ли получить одну строку из другой за 1 исправление:
// * замена одного символа в одной строке
// * вставка/удаление одного символа из одной строки
// Примеры тестовых сценариев:
//   first = "a", second = "b" -> true
//   first = "ab", second = "b" -> true
//   first = "ab", second = "cb" -> true
//   first = "ab", second = "ba" -> false

/*
   a b c
   b b c

   count ++

   true

   a b

   axbc
   abc

   abc
   defz



*/

func Editing(first, second string, maxOperationCount int) bool {
	if len(first) == maxOperationCount && len(second) == maxOperationCount {
		return true
	}
	return editing(first, second, maxOperationCount)
}

func editing(first, second string, maxOperations int) bool {
	if first == second {
		return true
	}
	editor := NewStringEditor(first, second, maxOperations)
	return editor.CanEdit()
}

type StringEditor struct {
	maxOperations     int
	maxLengthRuneList []rune
	minLengthRuneList []rune

	insertOperationsCount int
	offset                int
}

func NewStringEditor(first, second string, maxOperations int) *StringEditor {
	maxLengthRuneList, minLengthRuneList := maxMinLenghtRuneList([]rune(first), []rune(second))
	return &StringEditor{
		maxLengthRuneList: maxLengthRuneList,
		minLengthRuneList: minLengthRuneList,
		maxOperations:     maxOperations,
	}
}

func (s *StringEditor) CanEdit() bool {
	lengthDifference := s.lengthDifference()
	switch {
	case lengthDifference == 0:
		return replaceOperationsCount(s.maxLengthRuneList, s.minLengthRuneList) <= s.maxOperations
	case lengthDifference > 0 && lengthDifference <= s.maxOperations:
		if len(s.maxLengthRuneList) == 0 || len(s.minLengthRuneList) == 0 {
			return true
		}
		return s.canEdit()
	default:
		return false
	}
}

func (s *StringEditor) canEdit() bool {
	lastMinLengthRuneListIndex := len(s.minLengthRuneList) - 1
	for index := 0; index < len(s.maxLengthRuneList); index++ {
		offsetIndex := s.offsetIndex(index)
		if offsetIndex > lastMinLengthRuneListIndex {
			for i := 0; i < offsetIndex-lastMinLengthRuneListIndex; i++ {
				s.addInsertOperation()
			}
			break
		}
		if s.maxLengthRuneList[index] == s.minLengthRuneList[offsetIndex] {
			continue
		}
		if s.canReplaceIndex(index) {
			return true
		}
		if !s.canAddInsertOperation() {
			return false
		}
		s.addInsertOperation()
	}
	return s.operationsCount() <= s.maxOperations
}

func (s *StringEditor) offsetIndex(index int) int {
	return index - s.offset
}

func (s *StringEditor) canReplaceIndex(index int) bool {
	index++
	offsetIndex := s.offsetIndex(index)

	operationsCount := s.operationsCount() + 1
	editor := StringEditor{
		maxOperations:     s.maxOperations - operationsCount,
		maxLengthRuneList: s.maxLengthRuneList[index:],
		minLengthRuneList: s.minLengthRuneList[offsetIndex:],
	}
	return editor.CanEdit()
}

func (s *StringEditor) canReplaceRemainFromIndex(index int) bool {
	lastMaxLengthRuneList := len(s.maxLengthRuneList) - 1
	lastMinLengthRuneList := len(s.minLengthRuneList) - 1

	replaceOperationsCount := s.replaceOperationsCountFromIndex(index)
	currentLastOffsetIndex := lastMaxLengthRuneList - s.offset
	replaceOperationsCount += currentLastOffsetIndex - lastMinLengthRuneList

	canReplace := replaceOperationsCount+s.insertOperationsCount <= s.maxOperations
	return canReplace
}

func (s *StringEditor) replaceOperationsCountFromIndex(index int) int {
	offsetIndex := s.offsetIndex(index)
	itemsLeft := len(s.minLengthRuneList[offsetIndex:])

	return replaceOperationsCount(
		s.maxLengthRuneList[index:index+itemsLeft],
		s.minLengthRuneList[offsetIndex:],
	)
}

func (s *StringEditor) canAddInsertOperation() bool {
	return s.insertOperationsCount < s.lengthDifference()
}

func (s *StringEditor) lengthDifference() int {
	return len(s.maxLengthRuneList) - len(s.minLengthRuneList)
}

func (s *StringEditor) addInsertOperation() {
	s.insertOperationsCount++
	s.offset++
}

func (s *StringEditor) operationsCount() int {
	return s.insertOperationsCount
}

func replaceOperationsCount(firstRuneList, secondRuneList []rune) (operationCount int) {
	for index := range firstRuneList {
		if firstRuneList[index] != secondRuneList[index] {
			operationCount++
		}
	}
	return operationCount
}

func maxMinLenghtRuneList(firstRuneList, secondRuneList []rune) (max []rune, min []rune) {
	if len(firstRuneList) >= len(secondRuneList) {
		max = firstRuneList
		min = secondRuneList
	} else {
		max = secondRuneList
		min = firstRuneList
	}
	return
}

func intAbs(i int) int {
	if i >= 0 {
		return i
	} else {
		return -i
	}
}
