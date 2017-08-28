package lists

import (
	"testing"
)

func TestHQ(t *testing.T) {
	l := NewHierarchicalQueue(3, false)

	err := l.Enqueue("a", 2)
	err = l.Enqueue("b", 0)

	if err != nil {
		t.Error(err)
	}

	// if l.Dequeue() != "b" {
	// 	t.Error("dequeue failed")
	// }
	// if l.Dequeue() != "a" {
	// 	t.Error("dequeue failed")
	// }
}
