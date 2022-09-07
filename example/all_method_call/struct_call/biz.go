package struct_call


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

func (m * Member) email() {

}

func home() {
    user := User{}
    user.name()
    user.age()

    member := Member{}
    member.level()

    m2 := &Member{}
    m2.email()
}