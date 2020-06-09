package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	
	var n int
	fmt.Printf("Size of array to sort? (≤1000) ")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if n > 1000 {
		fmt.Println("Error:", n, "> 1000 - too long!")
		return
	} else if n < 1 {
		fmt.Println("Error:", n, "< 1")
		return
	}
	
	array := make([]int, n)
	rand.Seed(time.Now().UTC().UnixNano())
	
	fmt.Println("Original Array was:")
	for i := range array {
		array[i] = rand.Intn(10000)
		fmt.Printf("%d\t", array[i])
	}
	fmt.Println()
	
	bubbleSortedArray := BubbleSort(array)
	
	fmt.Println("Bubble Sorted Array is: ")
	for i := range bubbleSortedArray {
		fmt.Printf("%d\t", bubbleSortedArray[i])
	}
	fmt.Println()
	
	mergeSortedArray := MergeSort(array)
	
	fmt.Println("Merge Sorted Array is: ")
	for i := range mergeSortedArray {
		fmt.Printf("%d\t", mergeSortedArray[i])
	}
	fmt.Println()
}

// Реализация пузырьковой сортировки
func BubbleSort(inputArray []int) (sortedArray []int) {
	array := make([]int, len(inputArray))
	copy(array, inputArray)
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-1; j++ {
			if array[j] > array[j+1] {
				// Idiomatic Go swap
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

// Реализация сортировки слиянием
func MergeSort(array []int) (sortedArray []int) {
	sortedArray = make([]int, len(array), cap(array))
	splitMerge(array, sortedArray)
	return array
}

func splitMerge(A []int, B []int) {
	if len(A) < 2 { // if size == 1, считаем что отсортировано
		return
	}
// рекурсивное разбиваем на две половины до тех пор, пока размер не станет = 1,
// затем объединяем и возвращаем.
	iMiddle := len(A) / 2              // середина
	splitMerge(A[:iMiddle], B[:iMiddle])  // левая половина
	splitMerge(A[iMiddle:], B[iMiddle:])  // правая половина
	merge(A, B, iMiddle)  // объединяем две половины
	copy(A, B)             // копируем объединенные половины обратно в А
}

func merge(A []int, B []int, iMiddle int) {
	i0 := 0
	i1 := len(A) / 2
	iEnd := len(A)
	
	for j := range A {
		// если левая <= правой половины
		if i0 < iMiddle && (i1 >= iEnd || A[i0] <= A[i1]) {
			B[j] = A[i0]
			i0++  // Увеличиваем i0 после использования в качестве индекса.
		} else {
			B[j] = A[i1]
			i1++  // Увеличиваем значение i1 после использования в качестве индекса.
		}
	}
	
}
