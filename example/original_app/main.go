package main

import (
    "fmt"
    "owlet-go/example/original_app/api"
)

func main() {
    homepageFacade := api.HomepageFacade{}
    data := homepageFacade.GetData()
    fmt.Println(data)
}
