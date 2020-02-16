package domain

import "time"

// Token 認証を扱うドメイン
type Token struct {
	id     uint
	userId uint
	token  string
	exp    time.Time
}

// Set is
func (t *Token) Set(id uint, userId uint, token string, exp time.Time) {
	t.id = id
	t.userId = userId
	t.token = token
	t.exp = exp
}

// Id is
func (t *Token) Id() uint {
	return t.id
}

// UserId is
func (t *Token) UserId() uint {
	return t.userId
}

// Token is
func (t *Token) Token() string {
	return t.token
}

// Exp is
func (t *Token) Exp() time.Time {
	return t.exp
}
