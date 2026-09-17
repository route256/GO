package skus

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type listNode struct {
	element int64
	next    *listNode
}

var listHead, listTail *listNode

var nodesMap map[int64]*listNode

func IsSkusReady() bool {
	return len(nodesMap) > 0
}

func init() {

	defaultDataFileName := "init/skus.json"

	if _, err := os.Stat(defaultDataFileName); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("init/skus.json is absent ... unit-test? %v", err)
		return
	}

	InitWithDataFileName(defaultDataFileName)
}

func InitWithDataFileName(filename string) {

	var skus []int64
	fdata, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(fdata, &skus)

	nodesMap = make(map[int64]*listNode, len(skus))

	//zero element
	listHead = &listNode{
		element: 0,
	}
	listTail = listHead
	nodesMap[0] = listHead

	for _, v := range skus {

		node := &listNode{
			element: v,
		}
		nodesMap[v] = node

		listTail.next = node
		listTail = node
	}
}
