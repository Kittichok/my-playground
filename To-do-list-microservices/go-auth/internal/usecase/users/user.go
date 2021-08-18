package users

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/kittichok/go-auth/internal/models"
	"github.com/kittichok/go-auth/internal/repository"
	"github.com/kittichok/go-auth/internal/utils"
)

type RespUser struct {
	AccessToken string `json:"AccessToken"`
	TokenType string `json:"TokenType"`
	ExpiresIn int64 `json:"ExpiresIn"`
}

type UseCase interface {
	// GetUsers(key string) (models.User)
	SignIn(models.User) (*RespUser, error)
	SignUp(models.User) error
}

type UserUsecase struct {
	uRep repository.IUserRepository
}

func NewUserUseCase(uRep repository.IUserRepository) UseCase {
	return UserUsecase{uRep}
}

func (u UserUsecase) SignUp(user models.User) error {
	existUser := u.uRep.FindByName(user.Username)
	if existUser != nil {
		return models.ErrExists
	}

	salt, err := utils.GenerateRandomBytes(utils.SaltSize)
	if err != nil {
		return models.ErrInternal
	}

	hash := utils.HashPassword([]byte(user.Password), salt)
	user.Password = string(hash)
	user.Salt = string(salt)
	user.ID = uuid.Must(uuid.NewRandom()).String()

	err = u.uRep.Add(user)
	if err != nil {
		// u.trace.issue(err)
		return err
	}
	return nil
}

func (u UserUsecase) SignIn(user models.User) (*RespUser, error) {
	existUser := u.uRep.FindByName(user.Username)
	if existUser == nil {
		return nil, models.ErrInvalidUser
	}

	user.Password = string(utils.HashPassword([]byte(user.Password), []byte(existUser.Salt)))
	pwdMatchUser, err := u.uRep.Find(user)
	if pwdMatchUser == nil {
		return nil, models.ErrInvalidUser
	}

	currentTime := time.Now().AddDate(0, 0, 2).Unix()
	expiresAt := int64(currentTime)
	claims := &jwt.StandardClaims{
		ExpiresAt: expiresAt,
		Issuer:    "auth.todo-list",
		Id: existUser.ID,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//TODO use env
	mySigningKey := []byte("AllYourBase")
	tokenString, _ := accessToken.SignedString(mySigningKey)

	t := models.Token{
		AccessToken: tokenString,
		IsActive:    true,
		UserID:      existUser.ID,
	}

	err = u.uRep.AddToken(t)
	if err != nil {
		return nil, err
	}

	var r RespUser
	r.AccessToken = tokenString
	r.TokenType = "Bearer"
	r.ExpiresIn = expiresAt
	return &r, nil
}