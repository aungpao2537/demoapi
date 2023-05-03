package models

type User struct {
	id   string
	name string
}

var users map[string]User = map[string]User{}

func NewUser(id, name string) User {
	u := User{
		id:   id,
		name: name,
	}
	users[id] = u
	return u
}

func GetUser(id string) *User {
	u := users[id]
	if u.id == "" {
		return nil
	}
	return &u
}

func (u User) GetID() string {
	return u.id
}

func (u User) GetName() string {
	return u.name
}

func (u User) ToMap() map[string]string {
	return map[string]string{
		"id":   u.id,
		"name": u.name,
	}
}
