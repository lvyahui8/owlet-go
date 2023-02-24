package owlet

import (
	"fmt"
	"golang.org/x/tools/go/ssa"
)

func GetGoFuncKey(fc  *ssa.Function) string {
	recv := fc.Signature.Recv()
	recvKey := ""
	pkgPath := ""
	if recv != nil {
		recvKey = "(" +  recv.Type().String() +")"
	}
	if fc.Pkg  != nil{
		pkgPath = fc.Pkg.Pkg.Path()
	}
	parentKey := ""
	if fc.Parent() != nil {
		parentKey = "[" + GetGoFuncKey(fc.Parent()) + "]"
	}
	return fmt.Sprintf("%s%s%s#%s",pkgPath,recvKey, parentKey,fc.Name())
}
