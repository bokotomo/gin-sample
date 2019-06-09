package port

type UserPort interface {
	CreateUser(email string, password string) (*string, error)
}
