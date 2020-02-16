package port

// AuthPort is
type AuthPort interface {
	Login(email, password string) (*string, error)
}
