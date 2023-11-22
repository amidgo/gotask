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

func editing(first, second string, maxMismatchCount int) bool {
	if first == second {
		return true
	}
	firstRuneList := []rune(first)
	secondRuneList := []rune(second)
	lengthDifference := intAbs(len(first) - len(second))
	switch {
	case lengthDifference == 0:
		return replaceCharCase(firstRuneList, secondRuneList, maxMismatchCount)
	case lengthDifference > 0 && lengthDifference <= maxMismatchCount:
		return removeCharCase(firstRuneList, secondRuneList, maxMismatchCount)
	default:
		return false
	}
}

func replaceCharCase(firstRuneList, secondRuneList []rune, maxMismatches int) bool {
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

func removeCharCase(firstRuneList, secondRuneList []rune, maxMismatches int) bool {
	maxLengthRuneList, minLengthRuneList := maxLenghtRuneList(firstRuneList, secondRuneList)
	var offset int
	var mismatchCount int

	for index := range maxLengthRuneList {

	}
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
