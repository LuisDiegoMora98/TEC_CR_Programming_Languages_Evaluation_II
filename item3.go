package main

//Item 3, receive an int Array and a value to insert, then it compares all elements in the array to the new value
//returns the number of comparisons until the number insertion or to determine the number is already in the array
func item3(arrayTo *[]int, newValue int) int {
	comparisons := 0
	for k := range *arrayTo {
		if (*arrayTo)[k] == newValue {
			return comparisons
		}
		comparisons++
	}
	//arrayTo = append(arrayTo, newValue)
	*arrayTo = append(*arrayTo, newValue)
	//[len(arrayTo)+1] = newValue
	return comparisons
}
