package entities

type User struct {
	Name     string
	Password string
	Username string
}

func (u *User) CollectionName() string {
	return "sys_user"
}
