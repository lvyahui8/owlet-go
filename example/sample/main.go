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

func Api() {
    func() {
        func() {

        }()
    }()
    func() {

    }()
}

type User struct {

}

func (u*User) Ping() {
    func() {
        func(){

        }()
    }()
}

func (u *User) GetName() {

}

func main() {
    Homepage()
}
