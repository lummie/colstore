package colstore

import (
	"fmt"
	"math/big"
	"unsafe"
)

// DenseColumn provides a Column <<interface>> compatible type that stores the ColType values in an Array and therefore is not suitable for sparsely populated data.
// In addition each index value is stored independently so a densely populated column that has repeating values is not stored efficiently.
type DenseColumn struct {
	store      []ColType
	startIndex uint
	endIndex   uint
	useTracker *big.Int
}

// ensure that all of the Column interface is implemented at compile time.
var _ ColTypeColumn = (*DenseColumn)(nil)

// NewDenseColumn create an instance of a DenseColumn for storing densely populated data
func NewDenseColumn(startIndex uint, capacity uint) *DenseColumn {
	// TODO need a check for startIndex + capacity > MAX uint

	dc := DenseColumn{
		startIndex: startIndex,
		store:      make([]ColType, 0, capacity),
		endIndex:   startIndex + capacity - 1,
		useTracker: big.NewInt(0),
	}
	return &dc
}

// EstimateUsage estimates the number of bytes of memory used by the store
func (d *DenseColumn) EstimateUsage() uint {
	// TODO add estimation data in store
	s := uint(unsafe.Sizeof(*d))                                // size of the DenseColumn in memory
	s += uint(len(d.useTracker.Bytes()))                        // number of bytes allocated for usage tracking
	s += uint(unsafe.Sizeof(d.store))                           // size of slice/array, not including contents
	s += uint(unsafe.Sizeof(new(ColType))) * uint(cap(d.store)) // size of ColType * allocated capacity of the array (as this is a Dense Column and memory is allocated whether needed or not)
	return s
}

// toLocalIndex converts the index to the local index taking into account the starting index supplied at construction time
func (d *DenseColumn) toLocalIndex(index uint) uint {
	return index - d.startIndex
}

// checkIndex checks that the supplied index is within the start and end range and Panics if not.
func (d *DenseColumn) checkIndex(index uint) {
	if !d.InRange(index) {
		panic(fmt.Sprintf("index out of range %v in [%v:%v]", index, d.startIndex, d.endIndex))
	}
}

// Get retrieves the value stored at index irrespective of if the value has been set before. If it has not been set then the las set value or the ColType default value will be returned.
// This method assumes that the callee has tested the index to be set already, otherwise use GetIfSet
func (d *DenseColumn) Get(index uint) ColType {
	d.checkIndex(index)
	index = d.toLocalIndex(index)
	return d.store[index]
}

// GetIfSet returns the stored value at index or an error if a value has not been stored.
func (d *DenseColumn) GetIfSet(index uint) (ColType, error) {
	d.checkIndex(index)
	if d.IsSet(index) {
		index = d.toLocalIndex(index)
		return d.store[index], nil
	}
	return 0, fmt.Errorf("Attempted to read an unset value at index %d", index)
}

// Set sets the stored value at index to be the supplied ColType value
func (d *DenseColumn) Set(index uint, value ColType) {
	d.checkIndex(index)
	index = d.toLocalIndex(index)
	d.useTracker.SetBit(d.useTracker, int(index), 1)
}

// Clear marks the value at the specified index to not be set.
func (d *DenseColumn) Clear(index uint) {
	d.useTracker.SetBit(d.useTracker, int(index), 0)
}

// IsSet returns a boolean true if the specified index has been set otherwise false
func (d *DenseColumn) IsSet(index uint) bool {
	index = d.toLocalIndex(index)
	return d.useTracker.Bit(int(index)) == 1
}

// RangeStart returns the starting index of this store
func (d *DenseColumn) RangeStart() uint {
	return d.startIndex
}

// RangeEnd returns the ending index of this store
func (d *DenseColumn) RangeEnd() uint {
	return d.endIndex
}

// InRange returns a boolean indicating if the index falls within the storage range
func (d *DenseColumn) InRange(index uint) bool {
	return index >= d.startIndex && index <= d.endIndex
}
