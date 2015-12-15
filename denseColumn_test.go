package colstore

import "testing"

func TestRangeStart(t *testing.T) {
	c := NewDenseColumn(0, 1000)
	if c.RangeStart() != 0 {
		t.Error("Range Start should be zero")
	}
}

func TestRangeEnd(t *testing.T) {
	c := NewDenseColumn(0, 1000)
	if c.RangeEnd() != 999 {
		t.Errorf("Range End should be 999 not %v", c.RangeEnd())
	}
}

func TestInRange(t *testing.T) {
	c := NewDenseColumn(100, 900)
	if c.InRange(99) {
		t.Errorf("%d should not be in range %d : %d", 99, c.RangeStart(), c.RangeEnd())
	}
	if !c.InRange(100) {
		t.Errorf("%d should be in range %d : %d", 100, c.RangeStart(), c.RangeEnd())
	}
	if !c.InRange(500) {
		t.Errorf("%d should be in range %d : %d", 500, c.RangeStart(), c.RangeEnd())
	}
	if !c.InRange(999) {
		t.Errorf("%d should be in range %d : %d", 999, c.RangeStart(), c.RangeEnd())
	}
	if c.InRange(1000) {
		t.Errorf("%d should not be in range %d : %d", 1000, c.RangeStart(), c.RangeEnd())
	}
}

func TestRanges(t *testing.T) {
	var tests = []struct {
		s uint
		c uint
	}{
		{0, 10},
		{1, 10},
		{100, 10},
		{1000, 100},
	}

	for no, test := range tests {
		c := NewDenseColumn(test.s, test.c)
		if v := c.RangeStart(); v != test.s {
			t.Errorf("Test #%v : Expected %v to be %v", no, v, test.s)
		}
		if v := c.RangeEnd(); v != test.s+test.c-1 {
			t.Errorf("Test #%v : Expected %v to be %v", no, v, test.s+test.c-1)
		}
	}

}

func TestGetSetTrackingIsCorrect(t *testing.T) {
	c := NewDenseColumn(0, 1000)
	if c.IsSet(0) != false {
		t.Error("Index zero should not be set")
	}
	c.Set(0, 0)
	if c.IsSet(0) == false {
		t.Error("Index zero should be set")
	}
	if c.IsSet(999) != false {
		t.Error("Index 999 should not be set")
	}
	c.Set(999, 0)
	if c.IsSet(999) == false {
		t.Error("Index 999 should be set")
	}
}

func TestSetAll(t *testing.T) {
	c := NewDenseColumn(0, 10)
	for i := uint(0); i < 10; i++ {
		c.Set(i, 0)
	}
}

func TestIndexesBeforeRange(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	c := NewDenseColumn(20, 100)
	c.Set(19, 0)
}

func TestIndexesAfterRange(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	c := NewDenseColumn(0, 100)
	c.Set(100, 0)
}

func TestEstimateUsage(t *testing.T) {

	c := NewDenseColumn(0, 100)
	var baseSize uint = 72 + 100*8 // 72 + 8 * capacity

	if s := c.EstimateUsage(); s != baseSize {
		t.Errorf("Expected %v to be %v", s, baseSize)
	}

	var tests = []struct {
		index         uint
		expectedBytes uint
	}{
		{0, baseSize + 1},
		{1, baseSize + 1},
		{2, baseSize + 1},
		{3, baseSize + 1},
		{4, baseSize + 1},
		{5, baseSize + 1},
		{6, baseSize + 1},
		{7, baseSize + 1},
		{8, baseSize + 2},
		{9, baseSize + 2},
		{10, baseSize + 2},
		{11, baseSize + 2},
		{12, baseSize + 2},
		{13, baseSize + 2},
		{14, baseSize + 2},
		{15, baseSize + 2},
		{16, baseSize + 3},
		{64, baseSize + 9},
	}

	for no, test := range tests {
		c.Set(test.index, 0)
		if s := c.EstimateUsage(); s != test.expectedBytes {
			t.Errorf("Test #%v : Expected %v to be %v", no, s, test.expectedBytes)
		}
	}

}
