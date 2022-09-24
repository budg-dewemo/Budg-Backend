package models

type IUser interface {
	GetUser(username string) User
}

// User is a user.
type User struct {
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	LastName string    `json:"lastName"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Balance  []Balance `json:"balance"`
}

var Users = []User{
	User{
		Username: "user1",
		Password: "password1",
	},
	User{
		Username: "user2",
		Password: "password2",
	},
}

func (u *User) GetUser(username string) User {
	for _, user := range Users {
		if user.Username == username {
			return user
		}
	}
	return User{}
}
