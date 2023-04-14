package schema

import (
	"fmt"
  "gorm.io/gorm"
  "golang.org/x/crypto/bcrypt"
	"os"
	"errors"
)

// GORM provides a defined gorm.Model struct which includes sensible defaults,
// have a look in the future: https://gorm.io/docs/models.html
type User struct {
	gorm.Model
	ID				 uint
	Name			 string
	Password	 string
}
func (*User) ComparePassword(pass string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(pass))
	if err != nil {
		return errors.New(fmt.Sprintf("Login failure detected for userID/name %s/%s\n",
			User.ID,
			User.Name,
		))
	}
	return nil
}

func NewUser(name string, pass string) User {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failure in password hash: %s\n", err)
		return nil
	}
	return User{
		Name: name,
		Password: hashedPass,
	}
}
