package param_call

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

    edge := testutil.FindEdge(g, "fc3", "fc4")
    assert.NotNil(t, edge)
    _,ok := edge.Site.Common().Value.(*ssa.Parameter)
    assert.True(t, ok)

    edge = testutil.FindEdge(g, "fc8$1", "fc9")
    assert.NotNil(t, edge)
    unOp,ok := edge.Site.Common().Value.(*ssa.UnOp)
    assert.True(t, ok)
    _,ok = unOp.X.(*ssa.FreeVar)
    assert.True(t, ok)
}
