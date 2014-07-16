package server

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"labix.org/v2/mgo"
	"os"
)

type Feedback struct {
	Name  string
	Email string
	Text  string
}

// Establish connection to MongoDB and write feedback to collection
func WriteFeedback(name string, email string, text string) int {
	// We like only correct e-mails
	if !govalidator.IsEmail(email) {
		fmt.Printf("'%v' is not an email!\n", email)
		return NOT_AN_EMAIL
	}
	// Length of name can't be less than five letters and more than thirty two letters
	if !govalidator.IsByteLength(name, 5, 32) {
		fmt.Printf("Name '%v' has invalid length!\n", name)
		return INVALID_NAME
	}
	// Length of text should be between  32 and 1024 letters
	if !govalidator.IsByteLength(text, 32, 1024) {
		fmt.Printf("Feedback '%v' has invalid length!\n", text)
		return INVALID_NAME
	}
	url := os.Getenv("DATABASE")
	fmt.Printf("DB_URL is %v\n", url)
	sess, err := mgo.Dial(url)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		return DB_UNAVAILABLE
	}
	defer sess.Close()
	sess.SetMode(mgo.Monotonic, true)

	c := sess.DB("asaskevich").C("feedback")
	err = c.Insert(&Feedback{name, email, text})
	if err != nil {
		fmt.Printf("Can't write to mongo, go error %v\n", err)
		return CANT_WRITE_TO_DB
	}
	return OK
}
