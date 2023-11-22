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
	if len(first) == 1 && len(second) == 1 {
		return true
	}
	return editing(first, second, maxOperationCount)
}

func editing(first, second string, maxOperations int) bool {
	if first == second {
		return true
	}
	firstRuneList := []rune(first)
	secondRuneList := []rune(second)
	lengthDifference := intAbs(len(first) - len(second))
	switch {
	case lengthDifference == 0:
		return replaceRuneCase(firstRuneList, secondRuneList, maxOperations)
	case lengthDifference > 0 && lengthDifference <= maxOperations:
		return appendRuneCase(firstRuneList, secondRuneList, maxOperations)
	default:
		return false
	}
}

func replaceRuneCase(firstRuneList, secondRuneList []rune, maxMismatches int) bool {
	var mismatchCount int
	for index := range firstRuneList {
		if index > len(secondRuneList)-1 {
			break
		}
		if firstRuneList[index] != secondRuneList[index] {
			mismatchCount++
		}
	}
	return mismatchCount <= maxMismatches
}

func appendRuneCase(firstRuneList, secondRuneList []rune, maxMismatches int) bool {
	if len(firstRuneList) == 0 || len(secondRuneList) == 0 {
		return true
	}
	maxLengthRuneList, minLengthRuneList := maxLenghtRuneList(firstRuneList, secondRuneList)
	var mismatchesCount int

	for index := range maxLengthRuneList {

		offsetIndex := index - mismatchesCount
		if offsetIndex > len(minLengthRuneList)-1 {
			mismatchesCount++
			continue
		}

		if maxLengthRuneList[index] == minLengthRuneList[offsetIndex] {
			continue
		}

		mismatchesCount++
	}
	if mismatchesCount > maxMismatches {
		return false
	}
	return true
}

func maxLenghtRuneList(firstRuneList, secondRuneList []rune) (max []rune, min []rune) {
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
