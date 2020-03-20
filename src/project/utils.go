package project

func reverse(nodes []Node) []Node {
  nodesLen := len(nodes)
  newNodes := make([]Node, nodesLen)

  for i, j := 0, nodesLen - 1; j >= 0; i, j = i + 1, j - 1 {
    newNodes[i] = nodes[j]
  }

  return newNodes
}