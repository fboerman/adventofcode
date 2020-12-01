package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func load_input(fname string) (entries []string){
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		entries = append(entries, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

type Node struct {
	// the node this node orbits around
	parent *Node
	// all its connections
	neighbors []*Node
	// the name of this node, corresponds to the global map key
	name string
	// previous node for the BFS
	previous *Node
}

func BFS(q []*Node, target string) bool{
	//execute BFS

	for len(q) != 0 {
		// pop the front of the queue for FIFO effect
		current := q[0]
		q = q[1:]
		// check all children
		for _, c := range current.neighbors {
			if c.name == target {
				// we have found santa!
				c.previous = current
				return true
			}
			if c.previous == nil {
				//c.visited = true
				c.previous = current
				q = append(q, c)
			}
		}
	}
	// no path found!
	return false
}

func main() {
	// load the input
	entries := load_input("day6_input.txt")

	// map that holds all the nodes, the keys are the names
	// the paths of the nodes are held as pointers by the nodes themselves
	m := make(map[string]*Node)

	// go through all entries and create a node for each, insert it into the map
	for _, entrie := range entries {
		// split the line at the orbit sign
		parts := strings.Split(entrie, ")")
		parent := parts[0]
		child := parts[1]

		// check if the parent node exists if not create it
		var parent_node *Node = nil
		if node, exists := m[parent]; !exists {
			parent_node = &Node{nil, []*Node{}, parent, nil}
			m[parent] = parent_node
		} else {
			parent_node = node
		}

		// retrieve the child from the map, create if does not exist
		if child_node, exists := m[child]; !exists {
			child_node = &Node{parent_node, []*Node{}, child, nil}
			m[child] = child_node
		} else {
			child_node.parent = parent_node
		}
	}

	// now that the db is build start counting
	// for each item count until there is no parent anymore
	total_count := 0
	for _, node := range m {
		local_count := 0
		var selected_node *Node = node.parent
		for selected_node != nil {
			local_count++
			selected_node = selected_node.parent
		}
		//fmt.Println(node.name, " has ", local_count, " orbits")
		total_count += local_count
	}

	fmt.Println("Part 1: ", total_count)

	// for part 2, do a pathfinding from YOU to SAN
	// unweighted graph so to find shortest path use simple breadth first search

	// first rebuild the graph that each node holds all its connections
	for _, node := range m {
		if node.parent != nil {
			node.neighbors = append(node.neighbors, node.parent)
			node.parent.neighbors = append(node.parent.neighbors, node)
		}
	}

	start := "YOU"
	stop := "SAN"

	root := m[start]
	var q []*Node
	root.previous = root
	q = append(q, root)
	found := BFS(q, stop)

	if !found {
		fmt.Println("Could not find path!")
		return
	}

	var path []string
	selected := m[stop]
	for selected.name != start {
		path = append(path, selected.name)
		selected = selected.previous
	}


	fmt.Println("Part 2: ")
	fmt.Println(path)
	fmt.Println(len(path)-2)

}