package colstore

import (
	"testing"
)

func TestColumnItemInRange(t *testing.T) {
	tests := []struct {
		value    uint
		expected int
	}{
		{1, -1},
		{2, -1},
		{3, -1},
		{4, -1},
		{5, -1},
		{6, -1},
		{7, -1},
		{8, -1},
		{9, -1},
		{10, 0},
		{11, 1},
		{13, 1},
		{14, 1},
		{15, 1},
		{16, 1},
		{17, 1},
	}
	root := NewColumnItem(10, -10)

	for _, test := range tests {
		if root.InRange(test.value) != test.expected {
			t.Errorf("Expected Range %d to be in item %v to return %d", test.value, root, test.expected)
		}
	}
}

func TestColumnItemInRangeAtZero(t *testing.T) {
	tests := []struct {
		value    uint
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 1},
		{4, 1},
		{5, 1},
		{6, 1},
		{7, 1},
		{8, 1},
		{9, 1},
		{10, 1},
	}
	root := NewColumnItem(0, -0)

	for _, test := range tests {
		if root.InRange(test.value) != test.expected {
			t.Errorf("Expected Range %d to be in item %v to return %d", test.value, root, test.expected)
		}
	}
}

func TestColumnItemInsertionConsecutive(t *testing.T) {
	// item added at 10
	root := NewColumnItem(10, 10)

	// insert before 10 - 9,8,7,6
	root.Set(9, 9)
	if root.p.index != 9 || root.p.value != 9 {
		t.Errorf("Expected item %v to be at index %d", root.p, 9)
	}

	root.Set(8, 8)
	if root.p.p.index != 8 || root.p.p.value != 8 {
		t.Errorf("Expected item %v to be at index %d", root.p, 8)
	}

	root.Set(7, 7)
	if root.p.p.p.index != 7 || root.p.p.p.value != 7 {
		t.Errorf("Expected item %v to be at index %d", root.p, 7)
	}

	root.Set(6, 6)
	if root.p.p.p.p.index != 6 || root.p.p.p.p.value != 6 {
		t.Errorf("Expected item %v to be at index %d", root.p, 6)
	}

	// insert after 10 - 11,12
	root.Set(11, 11)
	if root.n.index != 11 || root.n.value != 11 {
		t.Errorf("Expected item %v to be at index %d", root.n, 11)
	}

	root.Set(12, 12)

	if root.n.n.index != 12 || root.n.n.value != 12 {
		t.Errorf("Expected item %v to be at index %d", root.n, 12)
	}

}

func TestColumnItemInsertionEven(t *testing.T) {
	// item added at 10
	root := NewColumnItem(10, 10)

	// insert before 10 - 8, 6, 4, 2
	root.Set(8, 8)
	if root.p.index != 8 || root.p.value != 8 {
		t.Errorf("Expected item %v to be at index %d", root.p, 8)
	}

	root.Set(6, 6)
	if root.p.p.index != 6 || root.p.p.value != 6 {
		t.Errorf("Expected item %v to be at index %d", root.p, 6)
	}

	root.Set(4, 4)
	if root.p.p.p.index != 4 || root.p.p.p.value != 4 {
		t.Errorf("Expected item %v to be at index %d", root.p, 4)
	}

	root.Set(2, 2)
	if root.p.p.p.p.index != 2 || root.p.p.p.p.value != 2 {
		t.Errorf("Expected item %v to be at index %d", root.p, 2)
	}

	// insert after 10 - 12, 14
	root.Set(12, 12)
	if root.n.index != 12 || root.n.value != 12 {
		t.Errorf("Expected item %v to be at index %d", root.n, 12)
	}

	root.Set(14, 14)

	if root.n.n.index != 14 || root.n.n.value != 14 {
		t.Errorf("Expected item %v to be at index %d", root.n, 14)
	}

}

func BenchmarkSequentialForward(b *testing.B) {
	root := NewColumnItem(0, 0)
	println(b.N)
	for n := 1; n < b.N; n++ {
		root.Set(uint(n), columnItemType(n))
	}
}

func BenchmarkSequentialForwardDynamicRoot(b *testing.B) {
	root := NewColumnItem(0, 0)
	println(b.N)
	for n := 1; n < b.N; n++ {
		root = root.Set(uint(n), columnItemType(n))
	}
}

func BenchmarkSequentialForwardDynamicRootLots(b *testing.B) {
	root := NewColumnItem(0, 0)
	println(b.N)
	for n := 1; n < 4000000000; n++ {
		root = root.Set(uint(n), columnItemType(n))
	}
}
