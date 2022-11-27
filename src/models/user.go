package models

import (
	"BudgBackend/src/database"
	"BudgBackend/src/hashing"
	"fmt"
	"strconv"
)

type IUser interface {
	GetUser(username string) User
}

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	AvatarLink string `json:"avatar"`
	Username   string `json:"username"`
	Password   string `json:"password"`
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
	query := fmt.Sprintf("SELECT id,password FROM User WHERE username = '%s'", u.Username)
	rows, err := database.QueryDB(query)
	if err != nil {
		ErrorLogger.Println(err.Error())
	}
	i := 0
	var hashFromBD string
	for rows.Next() {
		i++
		err = rows.Scan(&u.ID, &hashFromBD)
		if err != nil {
			ErrorLogger.Println(err.Error())
		}
	}
	// si obtengo 0 resultados, no existe el usuario
	if i == 0 {
		WarningLogger.Println("Error validate login: ", u.Username)
		return false, fmt.Errorf("Password error")
	}
	// si obtengo 1 resultado, valido el hash
	if i == 1 {
		if u.validateHash(hashFromBD) {
			return true, nil
		} else {
			WarningLogger.Println("Error validate login: ", u.Username)
			return false, fmt.Errorf("Password error")
		}
	}
	ErrorLogger.Println("Multiple users with username: ", u.Username)
	return false, fmt.Errorf("Multiple users with username: %s", u.Username)

}
func (u *User) createAvatar() string {
	return fmt.Sprintf("https://ui-avatars.com/api/?name=%s+%s?length=2", u.Username, u.LastName)
}

func (u *User) CreateUser() (int64, error) {
	if u.checkUserExists() {
		return 0, fmt.Errorf("User already exists")
	}
	pswHashed, errHash := hashing.HashPassword(u.Password)
	if errHash != nil {
		return 0, fmt.Errorf("Error hashing password")
	}
	u.Password = pswHashed
	query := fmt.Sprintf("INSERT INTO User (name, last_name, username, password, email, avatar, active) VALUES ('%s', '%s', '%s', '%s', '%s', '%s',1)", u.Name, u.LastName, u.Username, u.Password, u.Email, u.createAvatar())
	id, err := database.InsertDB(query)
	if err != nil {
		ErrorLogger.Println("Error creating user: ", err)
		return 0, fmt.Errorf("Error creating user: ", err)
	}
	u.ID = int(id)
	_, errBudget := u.createBudget()
	if errBudget != nil {
		ErrorLogger.Println("Error creating budget: ", err)
		return 0, fmt.Errorf("Error creating budget: ", err)
	}
	return id, nil
}

func (u *User) GetUser() (User, error) {

	query := fmt.Sprintf("SELECT id, name, last_name as lastName, email, avatar, username FROM User WHERE id = %s", strconv.Itoa(u.ID))
	rows, err := database.QueryDB(query)
	if err != nil {
		ErrorLogger.Println("Error getting user: ", err)
		return User{}, fmt.Errorf("Error getting user: ", err)
	}
	for rows.Next() {
		err = rows.Scan(&u.ID, &u.Name, &u.LastName, &u.Email, &u.AvatarLink, &u.Username)
		if err != nil {
			ErrorLogger.Println("Error getting user: ", err)
			return User{}, fmt.Errorf("Error getting user: ", err)
		}
	}

	return *u, nil
}
func (u *User) createBudget() (int64, error) {
	query := fmt.Sprintf("INSERT INTO Budget (user_id, name, amount, start_date, end_date, current_budget) VALUES (%d, 'Budget', 0, now(),DATE_ADD(now(), INTERVAL 1 MONTH),1)", u.ID)

	id, err := database.InsertDB(query)
	if err != nil {
		ErrorLogger.Println("Error creating budget: ", err)
		return 0, fmt.Errorf("Error creating budget: ", err)
	}
	return id, nil
}

func (u *User) validateHash(hashFromBD string) bool {
	return hashing.CheckPasswordHash(u.Password, hashFromBD)
}
