package tree

import (
	"fmt"
	"sort"
)

type Record struct {
	ID int
	Parent int
}

type Node struct {
	ID int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	nodes := make(map[int]*Node)
	var root *Node
	for _, rec := range records {
		node, ok := nodes[rec.ID]
		if !ok {
			node = &Node{ID: rec.ID}
			nodes[rec.ID] = node
		} else {
			return nil, fmt.Errorf("dup")
		}
		if rec.Parent != rec.ID {
			parent, ok := nodes[rec.Parent]
			if !ok {
				parent = &Node{ID: rec.Parent}
				nodes[rec.Parent] = parent
				parent.Children = append(parent.Children, node)
				sort.Slice(parent.Children, func(i, j int) bool {
					return parent.Children[i].ID < parent.Children[j].ID
				})
			}
		} else {
			root = node
		}
	}
	if root == nil && len(records) != 0 {
		return nil, fmt.Errorf("root")
	}
	return root, nil
}
