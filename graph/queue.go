package graph

import "errors"

type queue struct {
	elements []*node
}

var (
	ErrEmptyQueue = errors.New("Queue is empty")
)

func NewQueue() *queue {
	return &queue{
		elements: []*node{},
	}
}

func (q *queue) Length() int {
	return len(q.elements)
}

func (q *queue) Contains(element *node) bool {
	for _, k := range q.elements {
		if element.GetName() == k.GetName() {
			return true
		}
	}

	return false
}

func (q *queue) Enqueue(element *node) {
	q.elements = append(q.elements, element)
}

func (q *queue) Dequeue() (*node, error) {
	var element *node

	if q.Length() > 0 {
		element = q.elements[0]
		q.elements = q.elements[1:]
		return element, nil
	}

	return nil, ErrEmptyQueue
}
