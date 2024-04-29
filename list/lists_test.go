package list

import "testing"

func TestLinkedList(t *testing.T) {
	// Create a new linked list
	list := NewLinkedList()

	// Test appending elements
	list.Append(1)
	list.Append(2)
	list.Append(3)

	// Test Get method
	if val := list.Get(0); val != 1 {
		t.Errorf("Expected 1, got %d", val)
	}
	if val := list.Get(1); val != 2 {
		t.Errorf("Expected 2, got %d", val)
	}
	if val := list.Get(2); val != 3 {
		t.Errorf("Expected 3, got %d", val)
	}

	// Test Size method
	if size := list.Size(); size != 3 {
		t.Errorf("Expected size 3, got %d", size)
	}

	// Test Remove method
	list.Remove(1) // Remove element at index 1 (value 2)
	if size := list.Size(); size != 2 {
		t.Errorf("Expected size 2 after removal, got %d", size)
	}
	if val := list.Get(1); val != 3 {
		t.Errorf("Expected 3 after removal, got %d", val)
	}

	// Test Insert method
	list.Insert(1, 4) // Insert 4 at index 1
	if size := list.Size(); size != 3 {
		t.Errorf("Expected size 3 after insertion, got %d", size)
	}
	if val := list.Get(1); val != 4 {
		t.Errorf("Expected 4 after insertion, got %d", val)
	}

	// Test Clear method
	list.Clear()
	if size := list.Size(); size != 0 {
		t.Errorf("Expected size 0 after clearing, got %d", size)
	}
}

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
