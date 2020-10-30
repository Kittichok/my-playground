package models

type Token struct {
	UserID       uint
	AccessToken  string
	RefreshToken string
	IsActive     bool
}

func SaveToken(t Token) error {
	err := DB.Model(&t).Create(t).Error
	return err
}
