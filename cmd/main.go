package main

import (
    "flag"
    "fmt"
    "github.com/go-playground/validator/v10"
    "os"
    "owlet-go"
)

type Args struct {
    owlet.InitProgramArgs
}

func main() {
    args := &Args{}
    flag.StringVar(&args.Path, "path", "", "code path")
    flag.StringVar(&args.Algorithm, "algo", "pta", "cha|rta|vta|pta")
    flag.Parse()
    validate := validator.New()
    var err error
    err = validate.Struct(args)
    if err != nil {
        _, _ = fmt.Fprintln(os.Stderr, err)
        flag.PrintDefaults()
        os.Exit(1)
    }
    program := &owlet.Program{}
    err = program.Load(args.InitProgramArgs)
    if err != nil {
        _, _ = fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
    fmt.Printf("load func len %d\n", len(program.Graph.NodeMap))
}
