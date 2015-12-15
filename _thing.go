package colstore

type ValueType int

const MaxUint = ^uint(0)

type rangedItem struct {
	StartIndex uint
	EndIndex   uint
	Value      ValueType
}

type RangedBlock struct {
	data []*rangedItem
}

func NewRangedBlock(capacity int) *RangedBlock {
	n := RangedBlock{
		data: make([]*rangedItem, 0, capacity),
	}
	n.data[0] = &rangedItem{
		StartIndex: 0,
		EndIndex:   MaxUint,
		Value:      0,
	}
	return &n
}

func (r *RangedBlock) Capacity() int {
	return cap(r.data)
}

func (r *RangedBlock) Set(index uint, value ValueType) {
	item := r.findItem(index)
	if item == nil {
		panic("Index not found in RangeBlock.data")
	}

	// if the value is the same then do nothing, else split the item
	if item.Value == value {
		return
	}

	newItem := &rangedItem{
		StartIndex: index,
		EndIndex:   index,
		Value:      value,
	}
	afterItem := &rangedItem{
		StartIndex: index + 1,
		EndIndex:   item.EndIndex,
		Value:      item.Value,
	}
	item.EndIndex = index - 1

	r.data = append(r.data, newItem, afterItem)
}

func (r *RangedBlock) findItem(index uint) (found *rangedItem) {
	for _, item := range r.data {
		if index >= item.StartIndex && index <= item.EndIndex {
			found = item
			return
		}
	}
	return
}
