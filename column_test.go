package colstore_test

import (
	"math/big"
	"testing"

	"github.com/lummie/colstore"
)

func TestNewColumnBlockConstructorAndBounds(t *testing.T) {
	var boundsTests = []struct {
		offset         *big.Int
		size           uint64
		expectedLBound *big.Int
		expectedUBound *big.Int
	}{
		{big.NewInt(0), 10, big.NewInt(0), big.NewInt(9)},
		{big.NewInt(0), 1000, big.NewInt(0), big.NewInt(999)},
		{big.NewInt(1), 1000, big.NewInt(1), big.NewInt(1000)},
		{big.NewInt(100), 10, big.NewInt(100), big.NewInt(109)},
		{big.NewInt(1000), 100000000000, big.NewInt(1000), big.NewInt(100000000999)},
		{big.NewInt(1000), 100000000000, big.NewInt(1000), big.NewInt(100000000999)},
	}

	for index, test := range boundsTests {
		c := colstore.NewColumnBlock(test.offset, test.size)
		if c.LBound().Cmp(test.expectedLBound) != 0 {
			t.Errorf("boundsTest LBound(%d): expected %v, actual %v", index, test.expectedLBound.Uint64(), c.LBound().Uint64())
		}
		if c.UBound().Cmp(test.expectedUBound) != 0 {
			t.Errorf("boundsTest UBound(%d): expected %v, actual %v", index, test.expectedUBound.Uint64(), c.UBound().Uint64())
		}

	}
}
