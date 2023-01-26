package port

// AuthPort is
type AuthPort interface {
	Login(email, password string) (*string, error)
	TokenExists(userId uint, token *string) (bool, error)
}
