//
// 認証を扱うドメイン
//
package domain

import "time"

type Token struct {
	id     uint
	userId uint
	token  string
	exp    time.Time
}

func (this *Token) Set(id uint, userId uint, token string, exp time.Time) {
	this.id = id
	this.userId = userId
	this.token = token
	this.exp = exp
}

func (this *Token) Id() uint {
	return this.id
}

func (this *Token) UserId() uint {
	return this.userId
}

func (this *Token) Token() string {
	return this.token
}

func (this *Token) Exp() time.Time {
	return this.exp
}
