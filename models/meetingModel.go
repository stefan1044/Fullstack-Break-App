package models

import (
	"main/db"
)

type Meeting struct {
	Username string `json:"username"`
	Time     string `json:"time"`
}

type MeetingModel struct{}

func (i MeetingModel) Insert(meeting Meeting) (err error) {
	err = db.GetDb().QueryRow("INSERT INTO `meetings` (`Username`, `Time`) VALUES (\"" + meeting.Username + `", "` +
		meeting.Time + `")`).Err()

	return
}

func (i MeetingModel) SelectByUsername(username string) (meeting Meeting, err error) {
	err = db.GetDb().QueryRow("SELECT Username, Freetime FROM freetimes WHERE freetimes.Username = \""+
		username+"\"").Scan(&meeting.Username, &meeting.Time)

	return
}
