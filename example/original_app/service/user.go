package service

type UserService struct {
}

type User struct {
	Id       uint64
	Username string
	Email    string
}

func (us UserService) GetLoggedUser() User {
	return User{
		Id:       1,
		Username: "yahui",
		Email:    "yahui@email.com",
	}
}
