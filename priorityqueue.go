// 优先队列（非线程安全）
package Algorithm

import (
	"container/heap"
	"fmt"
	"reflect"
)

// An Item is something we manage in a priority queue.
type Item struct {
	value    interface{} // The value of the item; arbitrary.
	priority int         // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A queue implements heap.Interface and holds Items.
type queue []*Item

func (q queue) Len() int { return len(q) }

func (q queue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return q[i].priority > q[j].priority
}

func (q queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

func (q *queue) Push(x interface{}) {
	n := len(*q)
	item := x.(*Item)
	item.index = n
	*q = append(*q, item)
}

func (q *queue) Pop() interface{} {
	old := *q
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*q = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (q *queue) update(item *Item, value interface{}, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(q, item.index)
}

type PriorityQueue struct {
	q         queue
	valueType reflect.Type
}

func NewPriorityQueue(valueType reflect.Type) *PriorityQueue {
	return &PriorityQueue{q: make([]*Item, 0), valueType: valueType}
}

func (pq *PriorityQueue) Len() int {
	return len(pq.q)
}

func (pq *PriorityQueue) isAccept(item *Item) bool {
	if item == nil || reflect.TypeOf(item.value) != pq.valueType {
		return false
	}
	return true
}

func (pq *PriorityQueue) Push(x *Item) {
	if pq.isAccept(x) {
		heap.Push(&pq.q, x)
	}
}

func (pq *PriorityQueue) Pop() *Item {
	if pq.Len() == 0 {
		return nil
	}
	return heap.Pop(&pq.q).(*Item)
}

func (pq *PriorityQueue) Update(item *Item, value interface{}, priority int) {
	if reflect.TypeOf(value) == pq.valueType {
		pq.q.update(item, value, priority)
	}
}

func (pq *PriorityQueue) ValueType() reflect.Type {
	return pq.valueType
}

// This example creates a PriorityQueue with some items, adds and manipulates an item,
// and then removes the items in priority order.
func main() {
	// Some items and their priorities.
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4, "apple1": 1,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := NewPriorityQueue(reflect.TypeOf(""))
	i := 0
	for value, priority := range items {
		pq.Push(&Item{
			value:    value,
			priority: priority,
			index:    i,
		})
		i++
	}
	// Insert a new item and then modify its priority.
	item := &Item{
		value:    "orange",
		priority: 1,
	}
	pq.Push(item)
	pq.Update(item, item.value, 5)

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := pq.Pop()
		fmt.Printf("%.2d:%s ", item.priority, item.value)
	}
}
