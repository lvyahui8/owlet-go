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
		hash := GetFuncBodyHash(n.Func)
		t.Logf("hash %s,fcKey %s\n",hash, GetGoFuncKey(n.Func))
	}
}
