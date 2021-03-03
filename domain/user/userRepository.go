package usr

type Repository interface {
	CreateUser(user User) (int, error)
	GetUserByEmail(email string) (*User, error)
}
