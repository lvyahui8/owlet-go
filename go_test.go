package owlet

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetGoFuncKey(t *testing.T) {
	prog := Program{}
	err := prog.Load(InitProgramArgs{
		Path:      "./example/sample",
		Algorithm: "vta",
	})
	require.Nil(t, err)
	for _, n := range prog.Graph.NodeMap {
		t.Logf("fcKey %s\n", GetGoFuncKey(n.Func))
	}
}
