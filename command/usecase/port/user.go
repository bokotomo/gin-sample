package port

// UserPort is
type UserPort interface {
	CreateUser(email, password string) (*string, error)
}
