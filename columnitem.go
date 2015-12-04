package colstore

// generic type that will be generated bu go generate
type columnItemType int

// columnItem stores a value and the number of occurrences of that item it also stores a previous and next item in the lined list
type columnItem struct {
	p, n  *columnItem
	index uint
	count uint
	value columnItemType
}

// NewColumnItem allocates a new column item.
// This should only be called once to create the root item as new items are automatically added to the linked list by the set method.
func NewColumnItem(index uint, value columnItemType) *columnItem {
	n := columnItem{
		index: index,
		value: value,
		count: 1,
	}
	return &n
}

//  index   count   l   h
//  0       5       0   4
//  5       2       5   6
//  10      5       10  14
//

func (c *columnItem) Set(index uint, value columnItemType) *columnItem {
	current := c

	var isInRange int
	// attempt to find the item that holds the range of the index we want to add
loop:
	for {
		isInRange = c.InRange(index)
		switch {
		case isInRange == -1:
			if current.p != nil {
				current = current.p // move to previous
			} else {
				break loop
			}
		case isInRange == +1:
			if current.n != nil {
				current = current.n //move to next
			} else {
				break loop
			}
		default:
			{

				break loop
			}
		}
	}

	// current should now point to an item holding the range
	// if isInRange != 0 then this is either the first or last item that is not in the range.
	n := NewColumnItem(index, value)

	switch {
	case isInRange == -1:
		current.insertBefore(n)
	case isInRange == +1:
		current.insertAfter(n)
	default:
		panic("Expected to insert into the same item")
	}
	return n
}

/*func (c *columnItem) findNearest(index uint) *columnItem {
	current := c
	var isInRange int
	for {
		isInRange = c.InRange(index)
		switch {
		case isInRange == -1:
			if current.p != nil {
				current = current.p // move to previous
			} else {
				break loop
			}
		case isInRange == +1:
			if current.n != nil {
				current = current.n //move to next
			} else {
				break loop
			}
		default:
			{

				break loop
			}
		}
	}
}*/

// InRange returns -1 if the index is lower, 0 if in the range and +1 if after the range covered by the columnItem
func (c *columnItem) InRange(index uint) int {
	switch {
	case index < c.index:
		return -1
	case index > c.index+c.count-1:
		return +1
	default:
		return 0
	}
}

// Adds an item before the current item
func (c *columnItem) insertBefore(n *columnItem) {
	o := c.p //store previous

	// o : o.n = n  (old previous)
	// n : n.o = o, n.n = c  (new)
	// c : c.o = n (current)
	if o != nil {
		o.n = n
	}
	n.p = o
	n.n = c
	c.p = n
}

// Adds an item before the next item
func (c *columnItem) insertAfter(n *columnItem) {
	o := c.n // store next

	// c: c.n = n (current)
	// n: n.p = c, n.n = o  (new)
	// o: o.p = n (old next)
	c.n = n
	n.p = c
	n.n = o
	if o != nil {
		o.p = n
	}
}

// removes the current item and joins the previous and next together
func (c *columnItem) remove() {
	p := c.p
	n := c.n

	// link previous to next
	if p != nil {
		p.n = n
	}

	// link next to previous
	if n != nil {
		n.p = p
	}

	// clear current links
	c.p = nil
	c.n = nil
}
