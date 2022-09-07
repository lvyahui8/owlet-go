package async_call

import (
    "all_method_call/testutil"
    "github.com/stretchr/testify/assert"
    "golang.org/x/tools/go/ssa"
    owlet "owlet-go"
    "testing"
)

func TestCallSite(t *testing.T) {
    prog := owlet.Program{}
    err := prog.Load(owlet.InitProgramArgs{
        Path: "./",
        Algorithm: "vta",
    })
    assert.Nil(t, err)
    g := prog.Graph
    edge := testutil.FindEdge(g, "fc1", "fc2")
    assert.NotNil(t, edge)
    _,ok := edge.Site.Common().Value.(*ssa.Function)
    assert.True(t, ok)

    edge = testutil.FindEdge(g, "fc1", "fc3")
    assert.NotNil(t, edge)
    _,ok = edge.Site.Common().Value.(*ssa.Function)
    assert.True(t, ok)
}

