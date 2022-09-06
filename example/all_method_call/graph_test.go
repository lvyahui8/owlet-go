package main

import (
    "github.com/stretchr/testify/assert"
    "golang.org/x/tools/go/ssa"
    owlet "owlet-go"
    "testing"
)

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

func TestCallType(t *testing.T) {
    prog := owlet.Program{}
    err := prog.Load(owlet.InitProgramArgs{
        Path: "./",
        Algorithm: "vta",
    })
    assert.Nil(t, err)
    painter := owlet.Painter{}
    g := prog.Graph
    err = painter.OutputGraphHtml(g, "")
    assert.Nil(t, err)
    edge := FindEdge(g, "fc1", "fc2")
    assert.NotNil(t, edge)
    _,ok := edge.Site.Common().Value.(*ssa.Function)
    assert.True(t, ok)

    edge = FindEdge(g, "fc3", "fc4")
    assert.NotNil(t, edge)
    _,ok = edge.Site.Common().Value.(*ssa.Parameter)
    assert.True(t, ok)

    edge = FindEdge(g, "fc5", "fc7")
    assert.NotNil(t, edge)
    _,ok = edge.Site.Common().Value.(*ssa.Call)
    assert.True(t, ok)
}
