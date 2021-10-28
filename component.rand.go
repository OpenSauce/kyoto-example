package main

import (
	"math/rand"
	"strconv"
)

type ComponentRand struct {
	Content string
}

func (c *ComponentRand) Init() {
	c.Content = strconv.Itoa(rand.Intn(1000))
}
