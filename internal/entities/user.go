package entities

type User struct {
	id   string
	name string
}

func (user *User) GetID() string {
	return user.id
}

func (user *User) SetID(id string) *User {
	user.id = id

	return user
}

func (user *User) GetName() string {
	return user.name
}

func (user *User) SetName(name string) *User {
	user.name = name

	return user
}
