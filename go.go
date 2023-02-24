package owlet

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"golang.org/x/tools/go/ssa"
	"strings"
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

func GetFuncBody(fc * ssa.Function) string {
	sb := &strings.Builder{}
	_, _ = fc.WriteTo(sb)
	return sb.String()
}

func md5V(str string) string  {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func GetFuncBodyHash(fc *ssa.Function) string {
	body := GetFuncBody(fc)
	return md5V(body)
}