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

func home() {
    user := User{}
    user.name()
    user.age()

    member := Member{}
    member.level()
}