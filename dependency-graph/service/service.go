package service

import (
	"fmt"
	"log"
	"sort"

	"github.com/mradulrathore/onboarding-assignments/dependency-graph/domain/graph"
	"github.com/mradulrathore/onboarding-assignments/dependency-graph/domain/node"
)

var g graph.Graph

func GetParent(id int) ([]*node.Node, error) {
	n, err := getNodeById(id)
	if err != nil {
		return nil, err
	}
	return n.Parent, nil
}

func GetChild(id int) ([]*node.Node, error) {
	n, err := getNodeById(id)
	if err != nil {
		return nil, err
	}
	return n.Child, nil
}

func GetAncestors(id int) ([]int, error) {
	n, err := getNodeById(id)
	if err != nil {
		return nil, err
	}

	var ancestors []int
	cb := func(i int) {
		ancestors = append(ancestors, i)
	}
	ancestorsDFS(n, cb)

	return ancestors, nil
}

func ancestorsDFS(n *node.Node, visitCb func(int)) {
	visited := map[int]bool{}

	if n == nil {
		return
	}
	visited[n.Id] = true
	visitCb(n.Id)

	for _, ancestor := range n.Parent {
		if visited[ancestor.Id] {
			continue
		}
		ancestorsDFS(ancestor, visitCb)
	}
}

func GetDescendants(id int) ([]int, error) {
	n, err := getNodeById(id)
	if err != nil {
		return nil, err
	}

	var descendants []int
	cb := func(i int) {
		descendants = append(descendants, i)
	}
	descendantsDFS(n, cb)

	return descendants, nil
}

func descendantsDFS(n *node.Node, visitCb func(int)) {
	visited := map[int]bool{}

	if n == nil {
		return
	}
	visited[n.Id] = true
	visitCb(n.Id)

	for _, descendants := range n.Child {
		if visited[descendants.Id] {
			continue
		}
		descendantsDFS(descendants, visitCb)
	}
}

func AddNode(id int, name string, metaData map[string]string) {
	n := node.Node{
		Id:       id,
		Name:     name,
		MetaData: metaData,
	}
	index := SearchNode(&n)
	insertAt(index, &n)
}

func SearchNode(n *node.Node) int {
	if g.Nodes == nil {
		return 0
	}

	index := sort.Search(len(g.Nodes), func(i int) bool {
		return g.Nodes[i].Id >= n.Id
	})

	return index
}

//return index of id if it exists
func IdExist(id int) (int, bool) {
	if g.Nodes == nil {
		return -1, false
	}

	index := sort.Search(len(g.Nodes), func(i int) bool {
		return g.Nodes[i].Id >= id
	})
	if index == len(g.Nodes) {
		return -1, false
	}

	return index, g.Nodes[index].Id == id
}

func insertAt(index int, n *node.Node) {
	if index == len(g.Nodes) {
		g.Nodes = append(g.Nodes, n)
		return
	}

	g.Nodes = append(g.Nodes[:index+1], g.Nodes[index:]...)
	g.Nodes[index] = n
}

func AddEdge(id1, id2 int) error {
	n1, err := getNodeById(id1)
	if err != nil {
		return err
	}

	n2, err := getNodeById(id2)
	if err != nil {
		return err
	}

	if g.Edges == nil {
		g.Edges = make(map[*node.Node][]*node.Node)
	}

	n1.Child = append(n1.Child, n2)
	n2.Parent = append(n2.Parent, n1)
	g.Edges[n1] = append(g.Edges[n1], n2)

	return nil
}

var (
	IdNotExistMsg = "id doesn't exist"
)

func getNodeById(id int) (*node.Node, error) {
	index, exist := IdExist(id)
	if !exist {
		err := fmt.Errorf("%s", IdNotExistMsg)
		log.Println(err)
		return nil, err
	}

	n := g.Nodes[index]

	return n, nil
}

func Display() string {
	return g.String()
}