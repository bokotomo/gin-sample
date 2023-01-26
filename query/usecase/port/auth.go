package port

// AuthPort is
type AuthPort interface {
	TokenExists(userId uint, token *string) (bool, error)
}
