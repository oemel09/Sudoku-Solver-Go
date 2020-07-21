package main

type Cell struct {
	isFixed bool
	number uint8
	oldNumber uint8
}

func (c *Cell) resetNumber()  {
	c.oldNumber = c.number
	c.number = 0
}

func (c *Cell) resetOldNumber() {
	c.oldNumber = 0
}

func (c Cell) String() string {
	if c.number == 0 {
		return ""
	} else {
		return string(c.number)
	}
}
