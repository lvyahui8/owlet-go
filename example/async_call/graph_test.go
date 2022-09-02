package main

import (
    "github.com/stretchr/testify/assert"
    owlet "owlet-go"
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
    err = painter.OutputGraphHtml(prog.Graph, "")
    assert.Nil(t, err)
}
