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
	editor := NewStringEditor(first, second, maxOperationCount)
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
	requiredInsertOperations := s.lengthDifference()
	replaceOperationsOnly := requiredInsertOperations == 0
	switch {
	case replaceOperationsOnly:
		if len(s.maxLengthRuneList) <= s.maxOperations {
			return true
		}
		operationsCount := replaceOperationsCount(s.maxLengthRuneList, s.minLengthRuneList)
		return operationsCount <= s.maxOperations
	case requiredInsertOperations > 0 && requiredInsertOperations <= s.maxOperations:
		if len(s.maxLengthRuneList) == 0 || len(s.minLengthRuneList) == 0 {
			return true
		}
		return s.canEdit()
	default:
		return false
	}
}

// n = length of maxLenghtList
// mp = max operations count
// O = n + mp(n-1 + mp-1(n-2 + mp-2(...)))
func (s *StringEditor) canEdit() bool {
	lastMinLengthRuneListIndex := len(s.minLengthRuneList) - 1
	for index := 0; index < len(s.maxLengthRuneList); index++ {
		offsetIndex := s.offsetIndex(index)
		if offsetIndex > lastMinLengthRuneListIndex {
			s.addRemainInsertOperations(index)
			break
		}
		if s.maxLengthRuneList[index] == s.minLengthRuneList[offsetIndex] {
			continue
		}
		if s.operationsCount() == s.maxOperations {
			return false
		}
		if s.canEditItemsAfterIndex(index) {
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

func (s *StringEditor) canEditItemsAfterIndex(index int) bool {
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

func (s *StringEditor) canAddInsertOperation() bool {
	return s.insertOperationsCount < s.lengthDifference()
}

func (s *StringEditor) lengthDifference() int {
	return len(s.maxLengthRuneList) - len(s.minLengthRuneList)
}

func (s *StringEditor) addRemainInsertOperations(index int) {
	lastMinLengthRuneListIndex := len(s.minLengthRuneList) - 1
	offsetIndex := s.offsetIndex(index)
	for i := 0; i < offsetIndex-lastMinLengthRuneListIndex; i++ {
		s.addInsertOperation()
	}
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
