package owlet

import (
    "encoding/json"
    "fmt"
    "golang.org/x/tools/go/callgraph"
    "io/ioutil"
    "os"
    "owlet-go/utils"
    "strconv"
)

type Painter struct {
}

type Obj map[string]interface{}

func (p *Painter) OutputGraphHtml(graph callgraph.Graph, graphName string) error {
    var res = make(map[string]interface{})
    var nodes []Obj
    var edges []Obj
    for _, node := range graph.Nodes {
        nodes = append(nodes, Obj{
            "id":    strconv.Itoa(node.ID),
            "label": node.Func.Name(),
            "pkg":   node.Func.Pkg.Pkg.Path(),
        })
        for _, edge := range node.Out {
            edges = append(edges, Obj{
                "source": strconv.Itoa(edge.Caller.ID),
                "target": strconv.Itoa(edge.Callee.ID),
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
    _ = utils.CopyFile(fmt.Sprintf("%s/resources/graph.html", utils.RootDir()), fmt.Sprintf("%s/graph.html", outputPath))
    _ = utils.CopyFile(fmt.Sprintf("%s/resources/graph.js", utils.RootDir()), fmt.Sprintf("%s/graph.js", outputPath))
    return ioutil.WriteFile(fmt.Sprintf("%s/graph.json", outputPath), bits, os.ModePerm)
}
