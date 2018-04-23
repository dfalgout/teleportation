package graph

import (
	"reflect"
	"testing"
)

func TestLength(t *testing.T) {
	q := NewQueue()

	if q.Length() > 0 {
		t.Log("Queue should be empty")
		t.Fail()
	}

	testNode1 := NewNode("test1")
	testNode2 := NewNode("test2")

	q.Enqueue(testNode1)
	if q.Length() != 1 {
		t.Log("Queue length should be 1")
		t.Fail()
	}

	q.Enqueue(testNode2)
	if q.Length() != 2 {
		t.Log("Queue length should be 2")
		t.Fail()
	}

	_, err := q.Dequeue()
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	if q.Length() != 1 {
		t.Log("Queue length should be 1")
		t.Fail()
	}
}

func TestEnqueue(t *testing.T) {
	q := NewQueue()

	testNode1 := NewNode("test1")
	testNode2 := NewNode("test2")

	shouldBeWith1 := []*node{
		testNode1,
	}

	shouldBeWith2 := []*node{
		testNode1,
		testNode2,
	}

	q.Enqueue(testNode1)

	if !reflect.DeepEqual(q.elements, shouldBeWith1) {
		t.Log("Queue elements are not correct")
		t.Fail()
	}

	q.Enqueue(testNode2)

	// Queue should contain same elements in same order as shouldBe slice
	if !reflect.DeepEqual(q.elements, shouldBeWith2) {
		t.Log("Queue elements are not correct")
		t.Fail()
	}
}

func TestDequeue(t *testing.T) {
	q := NewQueue()

	// Buld some test nodes
	testNode1 := NewNode("test1")
	testNode2 := NewNode("test2")
	testNode3 := NewNode("test3")
	testNode4 := NewNode("test4")

	// Queue the nodes
	q.Enqueue(testNode1)
	q.Enqueue(testNode2)
	q.Enqueue(testNode3)
	q.Enqueue(testNode4)

	shouldBeWith3 := []*node{
		testNode2,
		testNode3,
		testNode4,
	}

	// Ignore error here
	el1, _ := q.Dequeue()

	if el1 != testNode1 {
		t.Logf("%s != %s\n", el1, testNode1)
		t.Fail()
	}

	if q.Length() != 3 {
		t.Logf("Queue should have Length %d, but has Length %d\n", 3, q.Length())
		t.Fail()
	}

	if !reflect.DeepEqual(q.elements, shouldBeWith3) {
		t.Log("Queue elements don't match")
		t.Fail()
	}
}

func TestDequeueErrEmptyQueue(t *testing.T) {
	q := NewQueue()

	el, err := q.Dequeue()
	if err == nil {
		t.Log("Queue is empty and should return error")
		t.Fail()
	}

	if el != nil {
		t.Log("Queue is empty and returned element should be nil")
		t.Logf("Instead return element %s", el)
		t.Fail()
	}
}

func TestContains(t *testing.T) {
	q := NewQueue()

	// Buld some test nodes
	testNode1 := NewNode("test1")
	testNode2 := NewNode("test2")
	testNode3 := NewNode("test3")
	testNode4 := NewNode("test4")

	// Queue the nodes
	q.Enqueue(testNode1)
	q.Enqueue(testNode2)
	q.Enqueue(testNode3)
	q.Enqueue(testNode4)

	if !q.Contains(testNode1) {
		t.Log("Contains should return true")
		t.Fail()
	}
}
