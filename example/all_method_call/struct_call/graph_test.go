package struct_call

import (
    "all_method_call/testutil"
    "github.com/stretchr/testify/assert"
    "go/token"
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

    edge := testutil.FindEdge(g, "home", "name")
    assert.NotNil(t, edge)
    _,ok := edge.Site.Common().Value.(*ssa.Function)
    args := edge.Site.Common().Args
    assert.True(t, ok && len(args) == 1)
    unOp,ok := args[0].(*ssa.UnOp)
    assert.True(t, ok && unOp.Op == token.MUL)

    edge = testutil.FindEdge(g, "home", "age")
    assert.NotNil(t, edge)
    _,ok = edge.Site.Common().Value.(*ssa.Function)
    args = edge.Site.Common().Args
    assert.True(t, ok && len(args) == 1)
    unOp,ok = args[0].(*ssa.UnOp)
    assert.True(t, ok && unOp.Op == token.MUL)

    edge = testutil.FindEdge(g, "home", "level")
    assert.NotNil(t, edge)
    _,ok = edge.Site.Common().Value.(*ssa.Function)
    args = edge.Site.Common().Args
    assert.True(t, ok && len(args) == 1)
    _,ok = args[0].(*ssa.Alloc)
    assert.True(t, ok)

    edge = testutil.FindEdge(g, "home", "email")
    assert.NotNil(t, edge)
    _,ok = edge.Site.Common().Value.(*ssa.Function)
    args = edge.Site.Common().Args
    assert.True(t, ok && len(args) == 1)
    _,ok = args[0].(*ssa.Alloc)
    assert.True(t, ok)
}
