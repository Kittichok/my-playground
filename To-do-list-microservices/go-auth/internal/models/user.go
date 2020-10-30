package models

import "fmt"

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Tokens   []Token
}

func (u *User) AddToken(t Token) error {
	u.Tokens = append(u.Tokens, t)
	fmt.Println(u)
	fmt.Println(u.Tokens)
	err := DB.Model(&u).Update(User{
		Tokens: u.Tokens,
	}).Error
	return err
}
