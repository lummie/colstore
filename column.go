package colstore

import (
	"fmt"
	"math/big"
)

// ColType todo
type ColType int

// Columnable todo
type Columnable interface {
	Get(index big.Int) *ColType
	Set(index big.Int, value ColType)
	Remove(index big.Int)
}

// ColumnBlock todo
type ColumnBlock struct {
	offset *big.Int
	size   uint64
	root   *columnItem
}

//TODO: Make a block uint based and always zero indexed, add a wrapper struct for storing the block and offsets

// NewColumnBlock allocates and returns a new ColumnBlock with a lower bound of offset and an upper bound of offset+size
func NewColumnBlock(offset *big.Int, size uint64) *ColumnBlock {
	n := ColumnBlock{
		offset: offset,
		size:   size,
	}
	return &n
}

// Set sets the stored value at external index to value
func (c *ColumnBlock) Set(index *big.Int, value ColType) {
	diff := big.NewInt(0).Sub(index, c.offset)

	inBounds := (diff.Int64() >= 0) // check lower boundary
	inBounds = inBounds && (diff.Uint64() < c.size-1)

	if !inBounds {
		panic(fmt.Sprintf("index [%v] out of range [%v..%v]", index.String(), c.LBound().String(), c.UBound().String()))
	}

	c.setInternal(diff.Uint64(), value)
}

// setInternam iterates from the root item to find the insertion point for the new value
func (c *ColumnBlock) setInternal(internalIndex uint64, value ColType) {
	current := c.root

}

// LBound returns the lower boundary of the indexes this block represents
func (c *ColumnBlock) LBound() *big.Int {
	return c.offset
}

// UBound returns the upper boundary of the indexes this block represents
func (c *ColumnBlock) UBound() *big.Int {
	r := big.NewInt(0)
	r.Add(c.offset, big.NewInt(0).SetUint64(c.size-1))
	return r
}
