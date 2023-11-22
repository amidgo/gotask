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

	operationsCount int
	offset          int
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
	lengthDifference := len(s.maxLengthRuneList) - len(s.minLengthRuneList)
	switch {
	case lengthDifference == 0:
		return replaceRuneMismatches(s.maxLengthRuneList, s.minLengthRuneList) <= s.maxOperations
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
		if offsetIndex > len(s.minLengthRuneList)-1 {
			s.operationsCount += offsetIndex - lastMinLengthRuneListIndex
			break
		}
		if s.maxLengthRuneList[index] == s.minLengthRuneList[offsetIndex] {
			continue
		}
		if s.canReplaceRemainFromIndex(index) {
			return true
		}
		s.addAppendOperation()
	}
	return s.operationsCount <= s.maxOperations
}

func (s *StringEditor) offsetIndex(index int) int {
	return index - s.offset
}

func (s *StringEditor) canReplaceRemainFromIndex(index int) bool {
	lastMaxLengthRuneList := len(s.maxLengthRuneList) - 1
	lastMinLengthRuneList := len(s.minLengthRuneList) - 1

	offsetIndex := s.offsetIndex(index)
	itemsLeft := len(s.minLengthRuneList[offsetIndex:])

	replaceMismatches := replaceRuneMismatches(
		s.maxLengthRuneList[index:index+itemsLeft],
		s.minLengthRuneList[offsetIndex:],
	)

	currentLastOffsetIndex := lastMaxLengthRuneList - s.offset
	replaceMismatches += currentLastOffsetIndex - lastMinLengthRuneList

	return replaceMismatches+s.operationsCount <= s.maxOperations
}

func (s *StringEditor) addAppendOperation() {
	s.operationsCount++
	s.offset++
}

func replaceRuneMismatches(firstRuneList, secondRuneList []rune) (mismatchCount int) {
	for index := range firstRuneList {
		if firstRuneList[index] != secondRuneList[index] {
			mismatchCount++
		}
	}
	return mismatchCount
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
