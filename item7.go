package main

func Item7(pArray []int, pSize int, pValue int, pComparisons int) (bool, int) {
	if pSize == 0 {
		return false, 0
	}
	if pSize <= 2 {
		if (pArray[0] == pValue) || (pArray[1] == pValue) {
			return true, pComparisons
		}
		return false, pComparisons
	}
	pivot := pSize / 2
	if pArray[pivot] == pValue {
		return true, pComparisons
	}
	if pArray[pivot] > pValue {
		return Item7(pArray[:pivot], pivot, pValue, pComparisons+1)
	}
	return Item7(pArray[pivot:], pSize-pivot, pValue, pComparisons+1)
}
