package main

func Item6(pArray []int, pSize int, pValue int) (bool, int) {
	for index := 0; index < pSize; index++ {
		if pArray[index] == pValue {
			return true, index
		}
	}
	return false, 0
}
