package structs

import (
	"fmt"
)

//User struct
type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first"`
	LastName  string `json:"last"`
	Email     string `json:"email"`
	Auth      string `json:"auth"`
	Token     string `json:"token"`
}

func (u User) insertSQL() {
	fmt.Printf("    - Inserting user '" + u.ID + "' into database. ")
	_, err = db.Exec(fmt.Sprintf("INSERT INTO %s VALUES (%s,%s,%s,%s,%s,%s)", userTable, u.ID, u.FirstName, u.LastName, u.Email, u.Auth, u.Token))
	didError(err)
}

func (u User) updateSQL() {
	fmt.Printf("    - Updating user '" + u.ID + "' in database. ")
	_, err = db.Exec(fmt.Sprintf("UPDATE %s SET first='%s' last='%s' email='%s' auth='%s' token='%s' WHERE id='%s'", userTable, u.FirstName, u.LastName, u.Email, u.Auth, u.Token, u.ID))
	didError(err)
}
