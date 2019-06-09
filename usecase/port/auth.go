package port

type AuthPort interface {
	Login(email string, password string) (*string, error)
}
