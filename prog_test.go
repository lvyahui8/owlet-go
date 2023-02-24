package owlet

import (
	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa/ssautil"
	"strings"
	"testing"
	"time"
)

func TestFastLoad(t *testing.T) {
	//repoPath :="../hertz"
	//repoBasePkg := "github.com/cloudwego/hertz"
	repoPath :="../gin-master"
	repoBasePkg := "github.com/gin-gonic/gin"

	cfg := &packages.Config{
		Mode:  packages.LoadAllSyntax,
		Tests: false,
		// https://github.com/cloudwego/hertz
		Dir:   repoPath,
		Logf:  nil,
	}
	begin := time.Now().UnixMilli()
	initial, err := packages.Load(cfg, "./...")
	t.Logf("load package cost time %d ms\n",time.Now().UnixMilli() - begin)
	require.Nil(t, err)
	t.Logf("initial pkg len %d\n", len(initial))
	err = CheckErrors(initial)
	require.Nil(t, err)

	prog, pkgs := ssautil.Packages(initial, 0)
	t.Logf("all pkg len %d\n", len(pkgs))
	begin = time.Now().UnixMilli()
	prog.Build()
	t.Logf("prog build cost time %d ms\n",time.Now().UnixMilli() - begin)
	allFunctions := ssautil.AllFunctions(prog)
	require.NotEmpty(t, allFunctions)
	t.Logf("func len %d\n", len(allFunctions))
	repoFuncCnt := 0
	for fc := range allFunctions {
		if fc.Pkg != nil {
			if  strings.HasPrefix(fc.Pkg.Pkg.Path(),repoBasePkg) {
				repoFuncCnt ++
			}
		}
	}
	t.Logf("repo func len %d\n",repoFuncCnt)
}
