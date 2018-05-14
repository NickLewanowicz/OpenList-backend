package structs

import (
	"context"
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

// First resolves the Name field for User, it is all caps to avoid name clashes
func (u *User) First(ctx context.Context) *string {
	return &u.FirstName
}

//GetUser will get user with id
func (r *Resolver) GetUser(ctx context.Context, args struct{ Email string }) (*User, error) {
	return getUser(ctx, string(args.Email))
}

func getUser(ctx context.Context, Email string) (*User, error) {
	var user User
	fmt.Printf("Fetching user with ID '" + Email + "' from " + userTable)

	result, err := db.Query(fmt.Sprintf("SELECT id, first, last, email FROM %s where id='%s'", userTable, Email))
	didError(err)
	result.Next()

	err = result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	didError(err)
	return &user, nil
}
