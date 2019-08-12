package server

import (
	"fmt"
	"testing"
)

type TestContext struct {
	Index int
}

func (c *TestContext) Next() {
	c.Index++
	handles[c.Index](c)
}

var handles = make([]func(c *TestContext), 0)

func TestUser(t *testing.T) {
	handles = append(handles, func(c *TestContext) {
		fmt.Println("A-1")
		c.Next()
		fmt.Println("A-2")
	})
	handles = append(handles, func(c *TestContext) {
		fmt.Println("B-1")
		c.Next()
		fmt.Println("B-2")
	})
	handles = append(handles, func(c *TestContext) {
		fmt.Println("hello world")
	})

	c := &TestContext{Index: -1}
	c.Next()
}
