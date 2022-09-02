package main

import (
    "sync"
    "time"
)

func Homepage() {
    group := sync.WaitGroup{}
    group.Add(2)
    var handlers []func() = []func(){Profile,Asset}
    for _,handler := range handlers {
        go func(f func()) {
            defer group.Done()
            f()
        }(handler)
    }
    group.Wait()
    go Logging()
}

func Profile() {
    time.Sleep(time.Duration(5) * time.Second)
}

func Asset() {
    time.Sleep(time.Duration(5) * time.Second)
}

func Logging() {

}

func main() {
    Homepage()
}
