package go_tilities

import "math"

type MinHeap struct {
	HeapArray *[]int
	Size      int
	MaxSize   int // Current number of elements in min heap
}

func NewMinHeap(maxSize int) *MinHeap {
	return &MinHeap{
		HeapArray: &[]int{},
		Size:      0,
		MaxSize:   maxSize,
	}
}

func (m *MinHeap) parent(index int) int {
	return (index - 1) / 2
}

func (m *MinHeap) left(index int) int {
	return 2*index + 1
}

func (m *MinHeap) right(index int) int {
	return 2*index + 2
}

func (m *MinHeap) swap(first, second int) {
	(*m.HeapArray)[first], (*m.HeapArray)[second] = (*m.HeapArray)[second], (*m.HeapArray)[first]
}

func (m *MinHeap) InsertKey(item int) {
	if m.Size == m.MaxSize {
		return
	}

	*m.HeapArray = append(*m.HeapArray, item)
	m.Size++
	i := m.Size - 1

	//m.heapifyUp(m.Size - 1)
	for i != 0 && (*m.HeapArray)[i] < (*m.HeapArray)[m.parent(i)] {
		m.swap(i, m.parent(i))
		i = m.parent(i)
	}
}

func (m *MinHeap) decreaseKey(i int, newValue int) {
	(*m.HeapArray)[i] = newValue
	//m.heapifyUp(i)
	for i != 0 && (*m.HeapArray)[i] < (*m.HeapArray)[m.parent(i)] {
		m.swap(i, m.parent(i))
		i = m.parent(i)
	}
}

func (m *MinHeap) DeleteKey(index int) {
	m.decreaseKey(index, math.MinInt)
	m.ExtractMin()
}

func (m *MinHeap) GetMin() int {
	return (*m.HeapArray)[0]
}

func (m *MinHeap) ExtractMin() int {
	if m.Size <= 0 {
		return math.MaxInt
	}

	if m.Size == 1 {
		m.Size--
		return (*m.HeapArray)[0]
	}

	//Remove min from heap
	//Then move the last element to root
	//Perform heapify from root
	root := (*m.HeapArray)[0]
	(*m.HeapArray)[0] = (*m.HeapArray)[m.Size-1]
	*m.HeapArray = (*m.HeapArray)[:(m.Size - 1)]
	m.Size--
	m.minHeapify(0)
	//log.Println("New root:", (*m.HeapArray)[0])
	return root
}

func (m *MinHeap) heapifyUp(index int) {
	for index != 0 && (*m.HeapArray)[index] < (*m.HeapArray)[m.parent(index)] {
		m.swap(index, m.parent(index))
		index = m.parent(index)
	}
}

// A recursive method to heapify a subtree with the root at given index
// This method assumes that the subtrees are already heapified
func (m *MinHeap) minHeapify(i int) {
	l := m.left(i)
	r := m.right(i)
	smallest := i

	if l < m.Size && (*m.HeapArray)[l] < (*m.HeapArray)[i] {
		smallest = l
	}

	if r < m.Size && (*m.HeapArray)[r] < (*m.HeapArray)[smallest] {
		smallest = r
	}

	if smallest != i {
		m.swap(i, smallest)
		m.minHeapify(smallest)
	}
}
