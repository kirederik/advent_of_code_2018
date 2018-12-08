package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Missing file as the first argument")
	}

	answer1 := PartI(os.Args[1])
	fmt.Println("Part I: ", answer1)
}

func readAll(filepath string) ([]int, error) {
	f, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var arr []int
	for _, p := range strings.Split(string(f), " ") {
		v, _ := strconv.Atoi(p)
		arr = append(arr, v)
	}

	return arr, nil
}

func PartI(filepath string) int {
	data, _ := readAll(filepath)

	tree := Node{}
	a, _ := sumMetadata(data, &tree, 0)
	fmt.Printf("Part II: root.value = %+v\n", tree.childs[0].value)
	return a
}

type Node struct {
	value  int
	childs []*Node
}

func sumMetadata(data []int, tree *Node, start int) (int, int) {
	if len(data) == 0 {
		return 0, 0
	}

	nChilds := data[start]
	nMetadata := data[start+1]

	node := &Node{}
	tree.childs = append(tree.childs, node)

	read := 2

	var sum int
	for i := 0; i < nChilds; i++ {
		s, r := sumMetadata(data, node, start+read)
		sum += s
		read += r
	}

	var metadataSum int
	for i := 0; i < nMetadata; i++ {
		sum += data[i+read+start]
		metadataSum += data[i+read+start]
	}

	if len(node.childs) == 0 {
		node.value = metadataSum
	} else {
		for i := 0; i < nMetadata; i++ {
			entry := data[i+read+start]
			if entry <= len(node.childs) {
				node.value += node.childs[entry-1].value
			}
		}
	}

	read += nMetadata

	return sum, read
}
