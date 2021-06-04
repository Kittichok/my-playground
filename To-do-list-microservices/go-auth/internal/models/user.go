package models

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID       string `json:"id" gorm:"primary_key"`
	Username string `json:"username" binding:"required" gorm:"primary_key"`
	Password string `json:"password" binding:"required"`
	Salt     string `json:"-" binding:"required"`
	Tokens   []Token
}

func (u *User) AddToken(t Token) error {
	u.Tokens = append(u.Tokens, t)
	fmt.Println(u)
	fmt.Println(u.Tokens)
	err := DB.Model(&u).Update("Tokens", u.Tokens).Error
	return err
}

func (u User) MarshalJSON() ([]byte, error) {
	var tmp struct {
		ID       string `json:"id"`
		Username string `json:"username"`
	}
	tmp.Username = u.Username
	tmp.ID = u.ID
	return json.Marshal(&tmp)
}
