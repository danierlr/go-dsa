package heap

import (
	"cmp"
	"container/heap"
)

type goHeap[T any] struct {
	items       []T
	checkIsLess func(smaller, bigger T) bool
}

func (h *goHeap[T]) Len() int { return len(h.items) }
func (h *goHeap[T]) Less(indexA, indexB int) bool {
	return h.checkIsLess(h.items[indexA], h.items[indexB])
}
func (h *goHeap[T]) Swap(indexA, indexB int) {
	h.items[indexA], h.items[indexB] = h.items[indexB], h.items[indexA]
}

func (h *goHeap[T]) Push(item any) {
	h.items = append(h.items, item.(T))
}

func (h *goHeap[T]) Pop() any {
	oldItems := h.items
	num := len(oldItems)
	item := oldItems[num-1]
	var defaultValue T
	oldItems[num-1] = defaultValue
	h.items = oldItems[0 : num-1]
	return item
}

type Heap[T any] struct {
	goHeap *goHeap[T]
}

func (h *Heap[T]) Len() int {
	return h.goHeap.Len()
}

func (h *Heap[T]) Push(item T) {
	heap.Push(h.goHeap, item)
}

func (h *Heap[T]) Pop() (T, bool) {
	if h.Len() == 0 {
		var defaultValue T
		return defaultValue, false
	}

	item := heap.Pop(h.goHeap)
	return item.(T), true
}

func (h *Heap[T]) Peek() (T, bool) {
	if h.Len() == 0 {
		var defaultValue T
		return defaultValue, false
	}

	return h.goHeap.items[0], true
}

func newHeapCmp[T any](initialItems []T, checkIsLess func(smaller, bigger T) bool) *Heap[T] {
	goHeap := &goHeap[T]{
		items:       initialItems,
		checkIsLess: checkIsLess,
	}

	heap.Init(goHeap)

	h := &Heap[T]{
		goHeap: goHeap,
	}

	return h
}

func NewMinHeapCmp[T any](initialItems []T, checkIsLess func(smaller, bigger T) bool) *Heap[T] {
	return newHeapCmp(initialItems, checkIsLess)
}

func NewMaxHeapCmp[T any](initialItems []T, checkIsLess func(smaller, bigger T) bool) *Heap[T] {
	return newHeapCmp(initialItems, func(smaller, bigger T) bool {
		return checkIsLess(bigger, smaller)
	})
}

func defaultComparator[T cmp.Ordered](smaller, bigger T) bool {
	return smaller < bigger
}

func NewMinHeap[T cmp.Ordered](initialItems []T) *Heap[T] {
	return NewMinHeapCmp(initialItems, defaultComparator)
}

func NewMaxHeap[T cmp.Ordered](initialItems []T) *Heap[T] {
	return NewMaxHeapCmp(initialItems, defaultComparator)
}
