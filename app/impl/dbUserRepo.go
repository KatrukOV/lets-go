package impl

import (
	"errors"
	"fmt"
	"github.com/andrii-minchekov/lets-go/domain/user"
	"golang.org/x/crypto/bcrypt"
)

type dbUserRepository struct {
	db Database
}

func (r dbUserRepository) CreateUser(user usr.User) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)

	fmt.Println("password is ", string(hashedPassword), "original is ", user.Password)

	if err != nil {

		return 0, err
	}

	stmt := "INSERT INTO users (name, email, password, created) VALUES ($1, $2, $3, CURRENT_DATE) RETURNING ID"

	result, err := r.db.Query(stmt, user.Name, user.Email, string(hashedPassword))
	if err != nil {
		return 0, err
	}

	if err := result.Next(); !err {
		return 0, errors.New("no id was generated")
	}
	var id int
	err = result.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r dbUserRepository) GetUserByEmail(email string) (*usr.User, error) {
	row := r.db.QueryRow("SELECT id, name, email, password FROM users WHERE email = $1", email)
	user := usr.User{}
	if err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password); err != nil {
		return nil, err
	}
	return &user, nil
}
