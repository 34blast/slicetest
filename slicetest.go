package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	fmt.Println("slicetest.main() : starting execution")
	fmt.Println()

	// Deal with int Arrays
	fmt.Println("slicetest.main() : Deal with int Arrays tests")
	var intArray2 []int
	intArray := []int{5, 4, 3, 2, 1}

	fmt.Println("slicetest.main() : print int slices")
	printIntArray(intArray2)
	/* create a slice numbers1 with double the capacity of intArray */
	intArray2 = make([]int, len(intArray), (cap(intArray))*2)

	copy(intArray2, intArray)
	printIntArray(intArray2)
	intArray2 = append(intArray2, 10, 20, 30)
	printIntArray(intArray2)

	// remove the 4th element 2
	intArray2 = append(intArray2[:3], intArray2[4], intArray2[5], intArray2[6], intArray2[7])
	fmt.Printf("[")
	for idx := range intArray2 {
		fmt.Printf("%d ", intArray2[idx])
	}
	fmt.Printf("]\n")

	printIntArray(intArray2)
	fmt.Println()

	fmt.Println("int Array (return new array) before sort = ", intArray)
	intArray = sortIntArray(intArray)
	fmt.Println("int Array (return new array)after  sort = ", intArray)
	fmt.Println()

	// Deal with String Arrays
	fmt.Println("slicetest.main() : Deal with String Arrays")
	stringArray := []string{"Rich", "Scott"}

	// Print the String Array 3 ways
	printStringArray(stringArray)

	fmt.Println()
	fmt.Println("slicetest.main() : ending execution")
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

func printIntArray(x []int) {
	fmt.Printf("len = %d cap = %d slice = %v\n", len(x), cap(x), x)
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
