package list

import "testing"

func TestBinarySearch(t *testing.T) {
	list := ArrayList.New()

	list.append(5)
	list.append(7)
	list.append(9)

	if list.get(2) != 9 {
		t.Errorf("get(2) failed, expected: %d, got: %d", 9, list.get(2))
	}

	removed := list.removeAt(1)
	if removed != 7 {
		t.Errorf("removeAt(1) failed, expected: %d, got: %d", 7, removed)
	}
	if list.length() != 2 {
		t.Errorf("length() failed, expected: %d, got: %d", 2, list.length())
	}

	list.append(11)
	removed = list.removeAt(1)
	if list.removeAt(1) != 9 {
		t.Errorf("removeAt(1) failed, expected: %d, got: %d", 9, removed)
	}
	removed = list.removeAt(9)
	if removed != -1 {
		t.Errorf("remove(9) failed, expected: %d, got: %d", -1, removed)
	}

	if list.length() != 0 {
		t.Errorf("length() failed, expected: %d, got: %d", 0, list.length())
	}

	list.prepend(5)
	list.prepend(7)
	list.prepend(9)

	if list.get(2) != 5 {
		t.Errorf("get(2) failed, expected: %d, got: %d", 5, list.get(2))
	}
	if list.get(0) != 9 {
		t.Errorf("get(0) failed, expected: %d, got: %d", 9, list.get(0))
	}
	if list.remove(9) != 9 {
		t.Errorf("remove(9) failed, expected: %d, got: %d", 9, list.remove(9))
	}
	if list.length() != 2 {
		t.Errorf("length() failed, expected: %d, got: %d", 2, list.length())
	}
	if list.get(0) != 7 {
		t.Errorf("get(0) failed, expected: %d, got: %d", 7, list.get(0))
	}
}
