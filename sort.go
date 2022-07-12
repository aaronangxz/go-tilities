package go_tilities

import (
	"math/rand"
)

/* Bubble Sort */

func BubbleSort(slice []int) {
	swapped := false
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-i-1; j++ {
			if slice[j] > slice[j+1] {
				Swap(&slice[j], &slice[j+1])
				swapped = true
			}
			if !swapped {
				break
			}
		}
	}
}

/* Heap Sort */

func heapify(s []int, n, i int) {
	//root is the largest
	largest := i
	l := 2*i + 1
	r := 2*i + 2

	//if left child is larger than root
	if l < n && s[l] > s[largest] {
		largest = l
	}

	//if right child is larger than root
	if r < n && s[r] > s[largest] {
		largest = r
	}

	//if at the end largest is no longer root
	if largest != i {
		Swap(&s[largest], &s[i])
		heapify(s, len(s), largest)
	}
}

func HeapSort(s []int) {
	//create max heap starting from one level above the leaf nodes -> n/2 - 1
	for i := len(s)/2 - 1; i >= 0; i-- {
		heapify(s, len(s), i)
	}

	//create max heap again on the reduced array
	for i := len(s) - 1; i > 0; i-- {
		Swap(&s[0], &s[i])
		//the last element is sorted, can ignore for now
		heapify(s[:i], i, 0)
	}
}

/* Insertion Sort */

func InsertionSort(s []int) {
	//treat first element as sorted
	for i := 1; i < len(s); i++ {
		curr := s[i]
		j := i - 1

		//Whenever current element is smaller than previous elements,
		//move back one position
		for j >= 0 && curr < s[j] {
			s[j+1] = s[j]
			j--
		}

		//eventually insert into the next sorted position
		s[j+1] = curr
	}
}

/* Merge Sort */

func MergeSort(s []int) []int {
	//when array is sliced until it has only one element, stop slicing
	if len(s) < 2 {
		return s
	}
	mid := len(s) / 2

	//merge the left and right portions
	return merge(MergeSort(s[:mid]), MergeSort(s[mid:]))
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))
	insertIdx := 0
	i, j := 0, 0

	//two pointers to iterate through both sides
	//insert the smaller element and move its pointer forward
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result[insertIdx] = left[i]
			i++
		} else {
			result[insertIdx] = right[j]
			j++
		}
		insertIdx++
	}

	//at least one of the sides is empty now
	//fill the remaining elements to the back of sorted array
	for i < len(left) {
		result[insertIdx] = left[i]
		i++
		insertIdx++
	}

	for j < len(right) {
		result[insertIdx] = right[j]
		j++
		insertIdx++
	}
	return result
}

/* Quick Sort */

func median(a, aIdx, b, bIdx, c, cIdx int) (int, int) {
	if a < b {
		switch {
		case b < c:
			return b, bIdx
		case a < c:
			return c, cIdx
		default:
			return a, aIdx
		}
	}
	switch {
	case a < c:
		return a, aIdx
	case b < c:
		return c, cIdx
	default:
		return b, bIdx
	}
}

func getPivot(v []int) (int, int) {
	n := len(v)
	r1 := rand.Intn(n)
	r2 := rand.Intn(n)
	r3 := rand.Intn(n)
	return median(v[r1], r1, v[r2], r2, v[r3], r3)
}

func QuickSort(s []int) {
	if len(s) < 2 {
		return
	}

	low, high := 0, len(s)-1

	//get pivot: median of 3 random elements
	pivot, pIdx := getPivot(s)

	//swap pivot to the last
	s[pIdx], s[high] = s[high], s[pIdx]

	//move elements smaller than pivot to the front
	for i := range s {
		if s[i] < pivot {
			s[low], s[i] = s[i], s[low]
			low++
		}
	}

	//no smaller elements, swap pivot back
	//this is the sorted position
	s[low], s[high] = s[high], s[low]

	//sort left and right partition, excluding sorted position
	QuickSort(s[:low])
	QuickSort(s[low+1:])
}

/* Selection Sort */

func SelectionSort(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		min := i
		for j := i + 1; j < len(slice); j++ {
			if slice[j] < slice[min] {
				min = j
			}
		}
		Swap(&slice[i], &slice[min])
	}
}
