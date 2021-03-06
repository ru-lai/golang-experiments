package main

import (
	"errors"
	"log"
)

type Node struct {
	Value int
	Data  string
	Left  *Node
	Right *Node
}

func (n *Node) Insert(value int, data string) error {
	if n == nil {
		return errors.New("Cannot insert val into a nil tree")
	}

	switch {
	case value == n.Value:
		return nil

	case value < n.Value:
		if n.Left == nil {
			n.Left = &Node{Value: value, Data: data}
			return nil
		}
		return n.Left.Insert(value, data)

	case value > n.Value:
		if n.Right == nil {
			n.Right = &Node{Value: value, Data: data}
			return nil
		}
		return n.Right.Insert(value, data)
	}
	return nil
}

func (n *Node) Find(value int) (string, bool) {
	if n == nil {
		return "", false
	}

	switch {
	case value == n.Value:
		return n.Data, true

	case value < n.Value:
		return n.Left.Find(value)

	default:
		return n.Right.Find(value)
	}
}

func main() {
	n := Node{Value: 5, Data: "data"}
	n.Insert(6, "sixes")
	six, _ := n.Find(6)
	log.Println(six)
	log.Println(n.Right)
}
