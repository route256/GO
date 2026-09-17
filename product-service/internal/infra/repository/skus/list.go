package skus

// List -
func List(skip int64, limit int32) []int64 {

	if limit <= 0 {
		return []int64{}
	}

	var (
		node  *listNode
		found bool
	)
	if node, found = nodesMap[skip]; !found && skip != 0 {
		return []int64{}
	}

	if skip == 0 {
		node = listHead
	}

	answer := make([]int64, 0, limit)
	i := 0

	for {
		node = node.next
		if node == nil || i == int(limit) {
			break
		}

		answer = append(answer, node.element)
		i++
	}

	return answer
}
