package user

import (
	"errors"

	mysqlDB "github.com/rakeshnayak12/bookstore-users-api-golang/database/mysql"
)

var userDB = make(map[uint64]*User)

// Save saves an user to data base
func (u *User) Save() error {
	queryInsert := "INSERT INTO users(user_id, first_name, last_name, email) VALUES(?, ?, ?, ?)"
	stmt, stmtErr := mysqlDB.UserDB.Prepare(queryInsert)
	if stmtErr != nil {
		return stmtErr
	}
	defer stmt.Close()
	_, insertErr := stmt.Exec(u.UserID, u.FirstName, u.LastName, u.Email)
	if insertErr != nil {
		return insertErr
	}

	// if result := userDB[u.UserID]; result != nil {
	// 	if result.Email == u.Email {
	// 		return errors.New("User already exits")
	// 	}
	// 	return errors.New("User already exits")
	// }
	// userDB[u.UserID] = u
	return nil
}

// GetUser finds a user from the database
func (u *User) GetUser() error {

	result := userDB[u.UserID]
	if result == nil {
		return errors.New("User doesn't exists")
	}

	u.UserID = result.UserID
	u.FirstName = result.FirstName
	u.LastName = result.LastName
	u.Email = result.Email
	return nil
}
