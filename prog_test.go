package owlet

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
	"strings"
	"testing"
	"time"
)

type LoadConf struct {
	RepoPath string
	RepoBasePkg string
}

func FullLoad(t *testing.T, conf LoadConf) map[string]*FcItem {
	return Load(t,conf,packages.LoadAllSyntax,true)
}


func FastLoad(t *testing.T, conf LoadConf) map[string]*FcItem {
	return Load(t,conf,packages.LoadAllSyntax,false,"./...")
}

func Load(t *testing.T,conf LoadConf,mode packages.LoadMode,full bool,patterns ... string) map[string]*FcItem  {
	t.Logf("load start .isFull %t",full)
	cfg := &packages.Config{
		Mode:  mode,
		Tests: false,
		// https://github.com/cloudwego/hertz
		Dir:   conf.RepoPath,
		Logf:  nil,
	}
	begin := time.Now().UnixMilli()
	initial, err := packages.Load(cfg, patterns...)
	t.Logf("load package cost time %d ms\n",time.Now().UnixMilli() - begin)
	require.Nil(t, err)
	t.Logf("initial pkg len %d\n", len(initial))
	err = CheckErrors(initial)
	require.Nil(t, err)
	var prog *ssa.Program
	var pkgs []*ssa.Package
	if full {
		prog ,pkgs = ssautil.AllPackages(initial, 0)
	} else {
		prog, pkgs = ssautil.Packages(initial, 0)
	}
	t.Logf("all pkg len %d\n", len(pkgs))
	begin = time.Now().UnixMilli()
	prog.Build()
	t.Logf("prog build cost time %d ms\n",time.Now().UnixMilli() - begin)
	allFunctions := ssautil.AllFunctions(prog)
	require.NotEmpty(t, allFunctions)
	t.Logf("func len %d\n", len(allFunctions))
	var repoFuncList []*ssa.Function
	var depFuncList[]*ssa.Function
	for fc := range allFunctions {
		if fc.Pkg != nil {
			if  strings.HasPrefix(fc.Pkg.Pkg.Path(),conf.RepoBasePkg) {
				repoFuncList = append(repoFuncList,fc)
				continue
			}
		}
		// pkg/mod/golang.org/x/tools@v0.6.0/go/ssa/create.go:113
		if strings.EqualFold(fc.Synthetic,"loaded from gc object file") {
		//if fc.Syntax() == nil {
			continue
		}
		depFuncList =  append(depFuncList,fc)
	}
	depFuncMap := toMap(depFuncList)
	t.Logf("dep func len %d\n",len(depFuncMap))
	t.Logf("repo func len %d\n", len(repoFuncList))
	t.Logf("load end")
	return toMap(repoFuncList)
}



type FcItem struct {
	hash string
	fc * ssa.Function
}

func toMap(fcList []*ssa.Function) map[string]*FcItem {
	m := make(map[string]*FcItem)
	for _, item := range fcList {
		m[GetGoFuncKey(item)] = func(fc *ssa.Function) *FcItem {
			defer func() {
				err := recover()
				if err != nil {
					fmt.Println(err)
				}
			}()
			return &FcItem {
				hash: GetFuncBodyHash(fc),
				fc : fc,
			}
		}(item)
	}
	return m
}

func TestFastLoad(t *testing.T) {
	//repoPath :="../hertz"
	//repoBasePkg := "github.com/cloudwego/hertz"
	repoPath :="../gin-master"
	repoBasePkg := "github.com/gin-gonic/gin"
	conf := LoadConf{
		RepoPath: repoPath,
		RepoBasePkg: repoBasePkg,
	}
	fullMap := FullLoad(t, conf)
	fastMap := FastLoad(t, conf)
	 // require.True(t, len(fcList2) == len(fcList1))
	 // fastLoad会把未使用的函数也算出来，可能比fullLoad反而多了一些仓库函数
	for srcKey, _ := range  fastMap{
		if _, ok := fullMap [srcKey]; !ok {
			t.Logf("fulMap %s not found\n",srcKey)
		}
	}

	for srcKey, _ := range fullMap {
		if _, ok := fastMap [srcKey]; !ok {
			t.Logf("fastMap %s not found\n",srcKey)
		}
	}

	for srcKey, srcItem := range fastMap {
		if tgtItem,ok := fullMap[srcKey]; ok {
			if !strings.EqualFold(srcItem.hash,tgtItem.hash) {
				// 无输出，拿到的函数body hash是一样的
				t.Logf("func key equals,but body not equals. fcKey %s\n",srcKey)
			}
		}
	}
}
