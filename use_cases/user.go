package use_cases

type User struct{}

func NewUser() *User {
	return &User{}
}

func (u *User) Get() {

}
