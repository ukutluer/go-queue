package queue

import "testing"

func TestQueuePeek(t *testing.T) {
	q := NewQueue[string]()
	q.Enqueue("1")
	q.Enqueue("2")
	q.Enqueue("3")
	expectedValue := "1"
	actualValue, _ := q.Peek()
	if actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestQueueEnqueue(t *testing.T) {
	q := NewQueue[string]()
	q.Enqueue("1")
	expectedValue := "1"
	actualValue, _ := q.Peek()
	if actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}

func TestQueueDequeue(t *testing.T) {
	q := NewQueue[string]()
	q.Enqueue("1")
	q.Enqueue("2")
	expectedValue := "1"
	actualValue, _ := q.Peek()
	if actualValue != expectedValue {
		t.Errorf("Got %v expected %v", actualValue, expectedValue)
	}
}
