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

func newColumnItem(index uint, value columnItemType) *columnItem {
	n := columnItem{
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

func (c *columnItem) set(index uint, value) {

    if index < c.low() {
        // before this item
    } else if index > c.high() {
        // after this item
        } else {
            // in the range of this item
        }
}


func (c *columnItem) low() uint {
	return c.index
}

func (c *columnItem) high() uint {
	return c.index + c.count - 1
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
	o := c.p // store next

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
