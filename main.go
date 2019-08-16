package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

const (
	size int = 10
)

func timeTaken(t time.Time, name string) {
	elapsed := time.Since(t)
	log.Printf("TIME: %s\t\t%s\n", name, elapsed)
}
func swap(i int, j int) (int, int) {
	return j, i
}

func generateRandomSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(199)
	}
	return slice
}

func bubbleSort(inArr []int) {
	defer timeTaken(time.Now(), "bubbleSort")
	arr := inArr
	n := 1
	for n < len(arr) {
		for i := 0; i < len(arr)-n; i++ {
			// compare arr[0] vs arr[1]
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = swap(arr[i], arr[i+1])
			}
		}
		n++
	}
	// fmt.Printf("Sorted: %v\n", arr)
}

func selectSort(arr []int) {
	defer timeTaken(time.Now(), "selectSort")
	for i := 0; i < len(arr); i++ {
		min := arr[i]
		for j := i; j < len(arr); j++ {
			if arr[j] < min {
				min = arr[j]
				arr[i], arr[j] = swap(arr[i], arr[j])
			}
		}
	}
	// fmt.Printf("Sorted: %v\n", arr)
}

func insertionSort(arr []int) {
	defer timeTaken(time.Now(), "insertionSort")
	for i := 0; i < len(arr); i++ {
		cursor := arr[i]
		pos := i

		for pos > 0 && arr[pos-1] > cursor {
			arr[pos] = arr[pos-1]
			pos = pos - 1
		}
		arr[pos] = cursor
	}
	// fmt.Printf("Sorted: %v\n", arr)
}

func mergeSortWrapper(arr []int) []int {
	fmt.Printf("Initial array: %v\n", arr)
	defer timeTaken(time.Now(), "mergeSort")
	merged := mergeSort(arr)
	fmt.Printf("Merge sorted: %v\n", merged)
	return merged
}
func mergeSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left, right := mergeSort(arr[:mid]), mergeSort(arr[mid:])
	merged := merge(left, right)
	// fmt.Printf("Input: %v\nMid: %v\nLeft: %v\nRight: %v\nMerged: %v\n\n", arr, mid, left, right, merged)
	return merged
}

func merge(l []int, r []int) []int {
	merged := make([]int, len(l)+len(r))

	leftCursor, rightCursor := 0, 0
	for leftCursor < len(l) && rightCursor < len(r) {

		if l[leftCursor] < r[rightCursor] {
			merged[leftCursor+rightCursor] = l[leftCursor]
			leftCursor++
		} else {
			merged[leftCursor+rightCursor] = r[rightCursor]
			rightCursor++
		}
	}
	for leftCursor < len(l) {
		merged[leftCursor+rightCursor] = l[leftCursor]
		leftCursor++
	}

	for rightCursor < len(r) {
		merged[leftCursor+rightCursor] = r[rightCursor]
		rightCursor++
	}
	return merged
}

func main() {
	randomInts := generateRandomSlice(size)

	bubbleSort(randomInts)

	randomInts = generateRandomSlice(size)
	selectSort(randomInts)

	randomInts = generateRandomSlice(size)

	insertionSort(randomInts)
	randomInts = generateRandomSlice(size)
	fmt.Printf("Unsorted: %v\n", randomInts)
	mergeSortWrapper(randomInts)
}
