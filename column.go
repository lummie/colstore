package colstore

// ColType TODO
type ColType int

// ColTypeColumn provides the basic interface for setting and getting values in a column.
// The type that implements this interface has a fixed index range that the boundaries can be queried using RangeStart() and RangeEnd() or the index tested with InRange()
// A specific index does not have to have a value stored against that index and this can be tested with IsSet
// Get() will return the value stored or the type's default value if the index has not been set
// GetIfSet() will return the value or an error if the value has not been set
// Set() sets the value at the specified index.
type ColTypeColumn interface {
	Get(index uint) ColType
	GetIfSet(index uint) (ColType, error)
	Set(index uint, value ColType)
	Clear(index uint)
	IsSet(index uint) bool
	RangeStart() uint
	RangeEnd() uint
	InRange(index uint) bool
	EstimateUsage() uint
}
