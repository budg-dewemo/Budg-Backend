package models

import (
	"BudgBackend/src/database"
	"fmt"
)

type IUser interface {
	GetUser(username string) User
}

// User is a user.
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	LastName string `json:"lastName"`
	Username string `json:"username"`
	Password string `json:"password"`
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

func (u *User) checkUserExists() bool {
	query := fmt.Sprintf("SELECT username as usernameDB FROM User WHERE username = '%s'", u.Username)
	rows, err := database.QueryDB(query)
	if err != nil {
		ErrorLogger.Println("Error checking if user exists: ", err)
	}
	i := 0
	var usernameDB string
	for rows.Next() {
		i++
		err = rows.Scan(&usernameDB)
	}
	switch i {
	case 0:
		return false
	default:
		if u.Username == usernameDB {
			return true
		} else {
			return false
		}
	}
}

func (u *User) checkActiveUser() bool {
	query := fmt.Sprintf("SELECT username FROM User WHERE username = '%s' AND active = true", u.Username)
	rows, err := database.QueryDB(query)
	if err != nil {
		ErrorLogger.Println("Error checking if user exists: ", err)
	}
	i := 0
	var usernameDB string
	for rows.Next() {
		i++
		err = rows.Scan(&usernameDB)
	}
	switch i {
	case 0:
		return false
	default:
		if u.Username == usernameDB {
			return true
		} else {
			return false
		}
	}
}

func (u *User) ValidateLogin() (bool, error) {
	if !u.checkUserExists() {
		return false, fmt.Errorf("User does not exist")
	}
	if !u.checkActiveUser() {
		return false, fmt.Errorf("User is not active")
	}
	query := fmt.Sprintf("SELECT id FROM User WHERE username = '%s' AND password = '%s'", u.Username, u.Password)
	rows, err := database.QueryDB(query)
	if err != nil {
		fmt.Println(err)
	}
	i := 0
	for rows.Next() {
		i++
		err = rows.Scan(&u.ID)
		if err != nil {
			fmt.Println(err)
		}
	}
	if i == 0 {
		WarningLogger.Println("Error validate login: ", u.Username)
		return false, fmt.Errorf("Password error")
	}
	if i == 1 {
		return true, nil
	}
	ErrorLogger.Println("Multiple users with username: ", u.Username)
	return false, fmt.Errorf("Multiple users with username: %s", u.Username)

}
