package compiler

import "strconv"

type Counter struct {
	Number uint16
}

func (c *Counter) Count() (n uint16) {
	n = c.Number
	c.Number += 1
	return n
}

func (c *Counter) String() string {
	return strconv.Itoa(int(c.Number))
}
