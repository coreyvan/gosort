package main

import (
	"log"
	"math/rand"
	"time"
)

const (
	size int = 10000
)

// timeTaken measures the time taken for a function to return
func timeTaken(t time.Time, name string) {
	elapsed := time.Since(t)
	log.Printf("TIME: %s\t\t%s\n", name, elapsed)
}

// swap returns the inputs in swapped order
func swap(i int, j int) (int, int) {
	return j, i
}

// generateRandomSlice generates a slice of ints of the size specified from 0-199
func generateRandomSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(199)
	}
	return slice
}

//bubbleSort sorts the input array with a bubble sort algorithm
func bubbleSort(inArr []int) {
	defer timeTaken(time.Now(), "bubbleSort")
	arr := inArr
	n := 1
	for n < len(arr) {
		for i := 0; i < len(arr)-n; i++ {
			// compare arr[0] vs arr[1] and swap
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = swap(arr[i], arr[i+1])
			}
		}
		n++
	}
}

// selectSort sorts the input array with a select sort algorithm
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
}

// insertionSort sorts the input array with an insertion sort algorithm
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
}

// mergeSortWrapper wraps the recursive function mergeSort
func mergeSortWrapper(arr []int) {
	defer timeTaken(time.Now(), "mergeSort")
	_ = mergeSort(arr)
}

// mergeSort recursively sorts the input array using a merge sort algorithm
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left, right := mergeSort(arr[:mid]), mergeSort(arr[mid:])
	merged := merge(left, right)
	return merged
}

// merge returns a merged sorted array using the two input arrays
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
	log.Printf("Input array size:\t\t%d", size)

	randomInts := generateRandomSlice(size)
	bubbleSort(randomInts)

	randomInts = generateRandomSlice(size)
	selectSort(randomInts)

	randomInts = generateRandomSlice(size)
	insertionSort(randomInts)

	randomInts = generateRandomSlice(size)
	mergeSortWrapper(randomInts)
}
