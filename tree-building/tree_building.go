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
	declared := make(map[int]bool)
	var root *Node
	for _, rec := range records {
		if declared[rec.ID] {
			return nil, fmt.Errorf("dup")
		}
		declared[rec.ID] = true
		node, ok := nodes[rec.ID]
		if !ok {
			node = &Node{ID: rec.ID}
			nodes[rec.ID] = node
		}
		if rec.Parent > rec.ID {
			return nil, fmt.Errorf("min")
		}
		if rec.Parent != rec.ID {
			parent, ok := nodes[rec.Parent]
			if !ok {
				parent = &Node{ID: rec.Parent}
				nodes[rec.Parent] = parent
			}
			parent.Children = append(parent.Children, node)
			sort.Slice(parent.Children, func(i, j int) bool {
				return parent.Children[i].ID < parent.Children[j].ID
			})
		} else {
			root = node
		}
	}
	for i := 0; i < len(declared); i++ {
		if _, ok := declared[i]; !ok {
			return nil, fmt.Errorf("noncontiguous")
		}
	}
	if root == nil && len(records) != 0 {
		return nil, fmt.Errorf("root")
	}
	seen := map[int]bool{}
	if root != nil && (!DFS(root, seen) || len(seen) != len(nodes)) {
		return nil, fmt.Errorf("cycle or disjoint")
	}
	return root, nil
}

func DFS(node *Node, seen map[int]bool) bool {
	if seen[node.ID] {
		return false
	}
	seen[node.ID] = true
	for _, c := range node.Children {
		if !DFS(c, seen) {
			return false
		}
	}
	return true
}
