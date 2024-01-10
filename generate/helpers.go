package main

import (
	"golang.org/x/net/html"
	"strings"
)

type FindOpts struct {
	level   int
	results int

	Criteria   func(node *html.Node) bool
	MaxLevel   int
	MaxResults int
}

func (opts *FindOpts) ResetCounters() {
	opts.level = 0
	opts.results = 0
}

func (opts *FindOpts) IsMaxLevel() bool {
	return opts.MaxLevel != 0 && opts.level > opts.MaxLevel
}

func (opts *FindOpts) IncLevel() {
	opts.level += 1
}

func (opts *FindOpts) DecLevel() {
	opts.level -= 1
}

func (opts *FindOpts) IsMaxResults() bool {
	return opts.MaxResults != 0 && opts.results >= opts.MaxResults
}

func (opts *FindOpts) IncResultsCount() {
	opts.results += 1
}

func getNodeText(node *html.Node) (text string) {
	findOpts := FindOpts{
		Criteria: func(node *html.Node) bool {
			return node.Type == html.TextNode
		},
	}

	if nodes := findAllNodes(node, &findOpts); nodes != nil {
		for _, n := range nodes {
			text += strings.Trim(n.Data, " ") + " "
		}
	}

	return strings.TrimRight(text, " ")
}

func findNextNode(node *html.Node, opts *FindOpts) *html.Node {
	opts.MaxResults = 1

	nodes := findAllNodes(node, opts)
	if nodes == nil || len(nodes) == 0 {
		return nil
	}

	return nodes[0]
}

func findAllNodes(node *html.Node, opts *FindOpts) (nodes []*html.Node) {
	if opts.IsMaxLevel() || opts.IsMaxResults() {
		return
	}

	opts.IncLevel()
	defer opts.DecLevel()

	nodes = make([]*html.Node, 0)

	if opts.Criteria(node) {
		nodes = append(nodes, node)
		opts.IncResultsCount()
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		foundNodes := findAllNodes(child, opts)
		if foundNodes == nil {
			return
		}

		nodes = append(nodes, foundNodes...)
	}

	return
}

func getNodeAttributes(node *html.Node) (attrs map[string]string) {
	attrs = make(map[string]string)

	if node != nil {
		for _, attr := range node.Attr {
			attrs[attr.Key] = attr.Val
		}
	}

	return
}
