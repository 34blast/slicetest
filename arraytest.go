package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	fmt.Println("arraytest.main() : starting execution")
	fmt.Println()

	// Show how to pass by Reference to change a value
	fmt.Println("arraytest.main() : Deal with pass by value tests for string")
	var myString = "unchanged"
	fmt.Println("Original String = ", myString)
	changeString(&myString)
	fmt.Println("changed String = ", myString)
	fmt.Println()

	// Show how to pass by Reference to change a value
	fmt.Println("arraytest.main() : Deal with pass by value tests for int")
	var myInt = 111
	fmt.Println("Original int = ", myInt)
	changeInt(&myInt)
	fmt.Println("changed int = ", myInt)
	fmt.Println()
	
	// Deal with int Arrays
	fmt.Println("arraytest.main() : Deal with int Arrays tests")
	intArray := []int{5, 4, 3, 2, 1}

	fmt.Println("int Array (return new array) before sort = ", intArray)
	intArray = sortIntArray(intArray)
	fmt.Println("int Array (return new array)after  sort = ", intArray)
	fmt.Println()
	
	intArray = []int{5, 4, 3, 2, 1}
	fmt.Println("int Array (pass by reference) before  sort ")
	printIntArray(intArray)
	sortIntArrayByRef(intArray)
	fmt.Println("int Array (pass by reference) after  sort ")
	printIntArray(intArray)
	fmt.Println()

	// Deal with String Arrays
	fmt.Println("arraytest.main() : Deal with String Arrays")
	stringArray := []string{"Rich", "Scott"}
	sValue := reflect.ValueOf(stringArray)

	// Print the String Array 3 ways
	printValue(sValue)
	fmt.Println()
	printStringArray(stringArray)
	fmt.Println()
	printSlice(sValue)
	fmt.Println()

	fmt.Println("arraytest.main() : ending execution")
}

func changeString(pString *string) {
	*pString = "Changed"
}

func changeInt(pInt *int) {
	*pInt = 999
}

func sortIntArrayByRef(pArray []int) {
	sort.Ints(pArray)
}

func sortIntArray(pArray []int) []int {
	retArray := pArray
	sort.Ints(pArray)
	return retArray
}

func printIntArray(pArray []int) {

	for idx := 0; idx < len(pArray); idx++ {
		fmt.Printf("index %d, value = %v\n", idx, pArray[idx])
	}

}

func printStringArray(pArray []string) {

	for idx := 0; idx < len(pArray); idx++ {
		fmt.Printf("index %d, value = %v\n", idx, pArray[idx])
	}

}

func printValue(pValue reflect.Value) {

	fmt.Println(pValue)

}

func printSlice(pValue reflect.Value) {

	fmt.Println()

	if pValue.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	for idx := 0; idx < pValue.Len(); idx++ {
		fmt.Printf("index %d, value = %v\n", idx, pValue.Index(idx).Interface())
	}

}
