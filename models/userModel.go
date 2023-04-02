package models

import (
	"fmt"
	"main/db"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserModel struct{}


func (i UserModel) Insert(user User) (err error) {

	err = db.GetDb().QueryRow("INSERT INTO `users` (`Username`, `Password`)" +
		" VALUES (\"" + user.Username + "\" , \"" + user.Password + "\" )").Err()
	//fmt.Println("INSERT INTO `items` (`Id`,`Action`,`Completed`)" +
	//	" VALUES (" + strconv.Itoa(int(item.Id)) + "," + "\"" + string(item.Action) +
	//	"\"" + "," + boolString + ")")

	return
}

func (i UserModel) SelectAll() (users []User, err error) {
	rows, err := db.GetDb().Query("SELECT Username, Password FROM users")
	defer rows.Close()
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}

	for i := 0; rows.Next(); i++ {
		var temp User
		if errRow := rows.Scan(&temp.Username, &temp.Password); errRow != nil {
			var zero []User
			return zero, err
		}
		users = append(users, temp)
	}

	return
}

func (i UserModel) SelectByUsername(username string) (user User, err error) {
	err = db.GetDb().QueryRow("SELECT Username, Password FROM users WHERE users.Username = \""+
		username+"\"").Scan(&user.Username, &user.Password)

	return
}

func (i UserModel) UpdateByUsername(username string, user User) (err error) {

	err = db.GetDb().QueryRow("UPDATE users SET" + " Username= \"" + user.Username + "\" , Password = \"" +
		user.Password + "\" WHERE Username = \"" + username + "\"").Err()

	return
}
