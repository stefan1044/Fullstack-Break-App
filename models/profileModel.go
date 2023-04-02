package models

import (
	"fmt"
	"main/db"
)

type Person struct {
	Firstname   string `json:"firstname"`
	Lastname    string `json:"password"`
	Username    string `json:"username"`
	Departament string `json:"departament"`
	Officename  string `json:"officename"`
	Teamname    string `json:"teamname"`
}

type PersonModel struct{}

func (i PersonModel) Insert(person Person) (err error) {

	err = db.GetDb().QueryRow("INSERT INTO `profiles` (`Firstname`, `Lastname`, `Username`, `Departament`, `Officename`, `Teamname`)" +
		` VALUES ("` + person.Firstname + `", ` + `"` + person.Lastname + `", ` + `"` + person.Username + `", ` + `"` +
		person.Departament + `", ` + `"` + person.Officename + `", ` + `"` + person.Teamname + `" )`).Err()

	return
}

func (i PersonModel) SelectAll() (persons []Person, err error) {
	rows, err := db.GetDb().Query("SELECT Firstname, Lastname, Username, Departament, Officename, Teamname FROM profiles")
	defer rows.Close()
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}

	for i := 0; rows.Next(); i++ {
		var temp Person
		if errRow := rows.Scan(&temp.Firstname, &temp.Lastname, &temp.Username, &temp.Departament, &temp.Officename, &temp.Teamname); errRow != nil {
			var zero []Person
			return zero, err
		}
		persons = append(persons, temp)
	}

	return
}

func (i PersonModel) SelectByFirstname(firstname string) (person Person, err error) {
	err = db.GetDb().QueryRow("SELECT Firstname, Lastname, Username, Departament, Officename, Teamname FROM profiles WHERE profiles.Firstname = \""+
		firstname+"\"").Scan(&person.Firstname, &person.Lastname, &person.Username, &person.Departament, &person.Officename, &person.Teamname)

	return
}
func (i PersonModel) SelectByLastname(lastname string) (person Person, err error) {
	err = db.GetDb().QueryRow("SELECT Firstname, Lastname, Username, Departament, Officename, Teamname FROM profiles WHERE profiles.Firstname = \""+
		lastname+"\"").Scan(&person.Firstname, &person.Lastname, &person.Username, &person.Departament, &person.Officename, &person.Teamname)

	return
}
func (i PersonModel) SelectByUsername(username string) (person Person, err error) {
	err = db.GetDb().QueryRow("SELECT Firstname, Lastname, Username, Departament, Officename, Teamname FROM profiles WHERE profiles.Username = \""+
		username+"\"").Scan(&person.Firstname, &person.Lastname, &person.Username, &person.Departament, &person.Officename, &person.Teamname)

	return
}

func (i PersonModel) UpdateByFirstrname(firstname string, person Person) (err error) {

	err = db.GetDb().QueryRow("UPDATE profiles SET" + " Firstname= \"" + person.Firstname + "\" , Password = \"" +
		person.Lastname + "\" , Username = \"" + person.Username + "\" , Departament = \"" + person.Departament +
		"\" , Officename = \"" + person.Officename + "\" , Teamname = \"" + person.Teamname + "\" WHERE Firstname = \"" + firstname + "\"").Err()

	return
}
func (i PersonModel) UpdateByLastname(lastname string, person Person) (err error) {

	err = db.GetDb().QueryRow("UPDATE profiles SET" + " Firstname= \"" + person.Firstname + "\" , Password = \"" +
		person.Lastname + "\" , Username = \"" + person.Username + "\" , Departament = \"" + person.Departament +
		"\" , Officename = \"" + person.Officename + "\" , Teamname = \"" + person.Teamname + "\" WHERE Firstname = \"" + lastname + "\"").Err()

	return
}
