package main

func Homepage() {
    defer func() {
        _ = recover()
    }()
    Service()
}

func Service() {
    // https://stackoverflow.com/questions/29518109/why-does-defer-recover-not-catch-panics
    // defer recover()
}

func main() {
    Homepage()
}
