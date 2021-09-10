package main

func newChain(id string, data string, previous *Chain) *Chain {
	return &Chain{id, data, previous}
}

type Chain struct {
	id       string
	data     string
	previous *Chain
}

func (c Chain) ID() string {
	return c.id
}

func (c Chain) Previous() *Chain {
	return c.previous
}
