package colstore_test

import (
	"testing"

	"github.com/lummie/colstore"
)

func TestColumnItem(t *testing.T) {
	root := colstore.NewColumnItem(10, -10)
	if root.InRange(1) != -1 {
		t.Errorf("Expected Range 1 to be in item %v", root)
	}
}
