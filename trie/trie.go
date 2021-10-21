package trie

import (
	"strings"
)

type HistoryTrie struct {
	root *Node
}

type Node struct {
	Value    string
	Children map[string]*Node
}

func NewTrie() *HistoryTrie {
	return &HistoryTrie{&Node{"", make(map[string]*Node)}}
}

func (h *HistoryTrie) Insert(entry string) {
	entries := strings.Split(entry, " ")
	current := h.root
	for i := 0; i < len(entries); i++ {
		entry := entries[i]
		_, exists := current.Children[entry]
		if !exists {
			current.Children[entry] = &Node{entry, make(map[string]*Node)}
		}
		current = current.Children[entry]
		current.Value = entry
	}
}

func (h *HistoryTrie) find(entry string) *Node {
	entries := strings.Split(entry, " ")
	current := h.root
	for i := 0; i < len(entries); i++ {
		key := entries[i]
		if current.Children[key] != nil {
			current = current.Children[key]
			if i == len(entries)-1 {
				return current
			}
		}
	}
	return nil
}

func (h *HistoryTrie) AutoComplete(entry string) map[string]struct{} {
	results := make(map[string]struct{})
	match := h.find(entry)
	if match != nil {
		h.getResults(match, entry, results)
	}
	return results
}

func (h *HistoryTrie) getResults(node *Node, prefix string, results map[string]struct{}) {
	if len(results) > 0 {
		prefix = prefix + " " + node.Value
	}
	if len(node.Children) == 0 {
		results[prefix] = struct{}{}
		return
	}
	for _, v := range node.Children {
		h.getResults(v, prefix, results)
	}
}
