package main

import (
    "github.com/stretchr/testify/assert"
    owlet "owlet-go"
    "owlet-go/utils"
    "testing"
)

func TestGraph(t *testing.T) {
    prog := owlet.Program{}
    err := prog.Load(owlet.InitProgramArgs{
        Path:      "./",
        Algorithm: "vta",
    })
    assert.Nil(t, err)
    painter := owlet.Painter{}
    err = painter.OutputGraphHtml(*prog.Graph, "")
    assert.Nil(t, err)
}

func TestRootDir(t *testing.T) {
    t.Logf("root path %s", utils.RootDir())
}
