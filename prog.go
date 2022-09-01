package owlet

import (
    "errors"
    "golang.org/x/tools/go/callgraph"
    "golang.org/x/tools/go/callgraph/cha"
    "golang.org/x/tools/go/callgraph/rta"
    "golang.org/x/tools/go/callgraph/vta"
    "golang.org/x/tools/go/packages"
    "golang.org/x/tools/go/pointer"
    "golang.org/x/tools/go/ssa"
    "golang.org/x/tools/go/ssa/ssautil"
)

type Program struct {
    Args  InitProgramArgs
    Graph *callgraph.Graph
}

type Algo string

type InitProgramArgs struct {
    Path      string `validate:"required"`
    Algorithm string `validate:"oneof=cha rta vta pta"`
}

func (p *Program) Load(args InitProgramArgs) error {
    p.Args = args
    cfg := &packages.Config{
        Mode:  packages.LoadAllSyntax,
        Tests: false,
        Dir:   args.Path,
        Logf:  nil,
    }
    initial, err := packages.Load(cfg, "./...")
    if err != nil {
        return err
    }
    err = CheckErrors(initial)
    if err != nil {
        return err
    }
    prog, pkgs := ssautil.AllPackages(initial, 0)
    prog.Build()
    mainPkgs := ssautil.MainPackages(pkgs)
    switch args.Algorithm {
    case "cha":
        p.Graph = cha.CallGraph(prog)
    case "rta":
        var roots []*ssa.Function
        for _, mainPkg := range mainPkgs {
            roots = append(roots, mainPkg.Func("main"))
            roots = append(roots, mainPkg.Func("init"))
        }
        res := rta.Analyze(roots, true)
        p.Graph = res.CallGraph
    case "vta":
        p.Graph = vta.CallGraph(ssautil.AllFunctions(prog), cha.CallGraph(prog))
    case "pta":
        result, err := pointer.Analyze(&pointer.Config{
            Mains:          mainPkgs,
            BuildCallGraph: true,
        })
        if err != nil {
            return err
        }
        p.Graph = result.CallGraph
    }

    return nil
}

func CheckErrors(pkgs []*packages.Package) error {
    eMsg := ""
    packages.Visit(pkgs, nil, func(pkg *packages.Package) {
        if len(pkg.Errors) > 0 {
            eMsg += "\npackage " + pkg.PkgPath + " contain errors"
            for _, err := range pkg.Errors {
                eMsg += "\n\t" + err.Error()
            }
        }
    })
    if eMsg != "" {
        return errors.New(eMsg)
    } else {
        return nil
    }
}
