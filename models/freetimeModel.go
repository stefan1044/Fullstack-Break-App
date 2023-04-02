package models

import (
	"main/db"
)

type Freetime struct {
	Username string `json:"username"`
	Freetime string `json:"freetime"`
}

type FreetimeModel struct{}

func (i FreetimeModel) Insert(freetime Freetime) (err error) {
	err = db.GetDb().QueryRow("INSERT INTO `freetimes` (`Username`, `Freetime`) VALUES (\"" + freetime.Username + `", "` +
		freetime.Freetime + `")`).Err()

	return
}

func (i FreetimeModel) SelectByUsername(username string) (freetime Freetime, err error) {
	err = db.GetDb().QueryRow("SELECT Username, Freetime FROM freetimes WHERE freetimes.Username = \""+
		username+"\"").Scan(&freetime.Username, &freetime.Freetime)

	return
}

func (i FreetimeModel) UpdateByUsername(username string, freetime Freetime) (err error) {
	err = db.GetDb().QueryRow("UPDATE freetimes SET" + " Username= \"" + freetime.Username + "\" , Freetime = \"" +
		freetime.Freetime + "\" WHERE Username = \"" + username + "\"").Err()

	return
}
