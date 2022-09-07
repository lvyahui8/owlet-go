package main

import (
    "github.com/stretchr/testify/assert"
    "go/token"
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

    edge = FindEdge(g, "main", "name")
    assert.NotNil(t, edge)
    _,ok = edge.Site.Common().Value.(*ssa.Function)
    args := edge.Site.Common().Args
    assert.True(t, ok && len(args) == 1)
    unOp,ok := args[0].(*ssa.UnOp)
    assert.True(t, ok && unOp.Op == token.MUL)

    edge = FindEdge(g, "main", "age")
    assert.NotNil(t, edge)
    _,ok = edge.Site.Common().Value.(*ssa.Function)
    args = edge.Site.Common().Args
    assert.True(t, ok && len(args) == 1)
    unOp,ok = args[0].(*ssa.UnOp)
    assert.True(t, ok && unOp.Op == token.MUL)

    edge = FindEdge(g, "main", "level")
    assert.NotNil(t, edge)
    _,ok = edge.Site.Common().Value.(*ssa.Function)
    args = edge.Site.Common().Args
    assert.True(t, ok && len(args) == 1)
    _,ok = args[0].(*ssa.Alloc)
    assert.True(t, ok)

    edge = FindEdge(g, "fc8$1", "fc9")
    assert.NotNil(t, edge)
    unOp,ok = edge.Site.Common().Value.(*ssa.UnOp)
    assert.True(t, ok)
    _,ok = unOp.X.(*ssa.FreeVar)
    assert.True(t, ok)
}
