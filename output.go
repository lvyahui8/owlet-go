package owlet

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

type Painter struct {
}

type Obj map[string]interface{}

func (p *Painter) OutputGraphHtml(graph * Graph, graphName string) error {
    var res = make(map[string]interface{})
    var nodes []Obj
    var edges []Obj
    for _, node := range graph.NodeMap {
        pkgPath := ""
        if  node.Func.Pkg != nil {
            pkgPath =  node.Func.Pkg.Pkg.Path()
        }
        nodes = append(nodes, Obj{
            "id":    node.ID,
            "label": node.Func.Name(),
            "pkg": pkgPath ,
        })
        for _, edge := range node.Out {
            edges = append(edges, Obj{
                "source":edge.CallerID,
                "target": edge.CalleeID,
            })
        }
    }
    res["nodes"] = nodes
    res["edges"] = edges
    bits, err := json.Marshal(res)
    if err != nil {
        return err
    }
    var outputPath string
    if graphName == "" {
        outputPath = "./output/"
    } else {
        outputPath = fmt.Sprintf("./output/%s/", graphName)
    }
    err = os.MkdirAll(outputPath, os.ModePerm)
    if err != nil {
        return err
    }
    _ = CopyFile(fmt.Sprintf("%s/resources/graph.html", RootDir()), fmt.Sprintf("%s/graph.html", outputPath))
    _ = CopyFile(fmt.Sprintf("%s/resources/graph.js", RootDir()), fmt.Sprintf("%s/graph.js", outputPath))
    return ioutil.WriteFile(fmt.Sprintf("%s/graph.json", outputPath), bits, os.ModePerm)
}
