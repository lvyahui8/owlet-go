package testutil

import owlet "owlet-go"

func FindNode(graph * owlet.Graph,name string) *owlet.Node {
    for _, node := range graph.NodeMap {
        if node.Func.Name() == name {
            return node
        }
    }
    panic("not found")
}

func FindEdge(graph * owlet.Graph,fc1 ,fc2 string) *owlet.Edge {
    n1 := FindNode(graph, fc1)
    n2 := FindNode(graph, fc2)
    return n1.Out[n2.ID]
}
