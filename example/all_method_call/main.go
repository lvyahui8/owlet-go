package main

func fc1() {
    fc2()
}

func fc2() {
}

func fc3(fc func()) {
    fc()
}

func fc4()  {

}

func fc5() {
    fc6()()
}

func fc6() func () {
    return fc7
}

func fc7() {

}

type User struct{}

func (u User) name() {
}

func (u User) age() {
}

func (u * User) phone() {

}

type Member struct{}

func (m * Member) level() {

}

func fc8(fc func()) func() {
    return func() {
        fc()
    }
}

func fc9() {

}

func main() {
    fc1()
    fc3(fc4)
    fc5()
    user := User{}
    user.name()
    user.age()

    member := Member{}
    member.level()

    fc8(fc9)()
}
