package list

import "testing"

func TestBinarySearch(t *testing.T) {
	list := NewArrayList()

	list.push(5) // 0
	list.push(7) //1
	list.push(9) //2

	if list.get(2) != 9 {
		t.Errorf("get(2) failed, expected: %d, got: %d", 9, list.get(2))
	}

	removed := list.delete(1)
	if removed != 7 {
		t.Errorf("removeAt(1) failed, expected: %d, got: %d", 7, removed)
	}
	if list.length != 2 {
		t.Errorf("length() failed, expected: %d, got: %d", 2, list.length)
	}

	list.push(11)
	removed = list.delete(1)
	if removed != 9 {
		t.Errorf("removeAt(1) failed, expected: %d, got: %d", 9, removed)
	}
	removed = list.delete(9)
	if removed != -1 {
		t.Errorf("remove(9) failed, expected: %d, got: %d", -1, removed)
	}

}
