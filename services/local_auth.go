package services

import (
	"example/models"
	"example/utils/encrypt"
	"example/utils/mysql"
	"time"
)

type RegisterSchema struct {
	FirstName string `json:"firstName" validate:"required,alpha,min=5,max=128"`
	LastName  string `json:"lastName" validate:"required,alpha,min=5,max=128"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,alphanum,min=6,max=128"`
}

type User struct {
	Token      string `json:"token"`
	ID         int    `json:"id"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	FacebookID string `json:"FacebookID"`
	GoogleID   string `json:"GoogleID"`
	Avatar     string `json:"avatar"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

type LoginSchema struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,alphanum,min=6,max=128"`
}

func Register(body RegisterSchema) (User, error) {
	hashPassword, err := encrypt.HashPassword(body.Password)
	if err != nil {
		return User{}, err
	}
	newUser := models.User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Password:  hashPassword,
	}
	err = mysql.DB.Create(&newUser).Error
	if err != nil {
		return User{}, err
	}
	token := GenerateToken(newUser.ID)
	createdUser := User{
		Token:      token,
		ID:         newUser.ID,
		FirstName:  newUser.FirstName,
		LastName:   newUser.LastName,
		Email:      newUser.Email,
		FacebookID: newUser.FacebookID,
		GoogleID:   newUser.GoogleID,
		Avatar:     newUser.Avatar,
		CreatedAt:  newUser.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  newUser.UpdatedAt.Format(time.RFC3339),
	}
	return createdUser, nil
}

func Login(body LoginSchema) (User, error) {
	user := models.User{}
	err := mysql.DB.Model(&models.User{}).Where("email = ?", body.Email).First(&user).Error
	if err != nil {
		return User{}, err
	}
	if err = encrypt.CheckPassword(body.Password, user.Password); err != nil {
		return User{}, err
	}
	token := GenerateToken(user.ID)
	result := User{
		Token:      token,
		ID:         user.ID,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		FacebookID: user.FacebookID,
		GoogleID:   user.GoogleID,
		Avatar:     user.Avatar,
		CreatedAt:  user.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  user.UpdatedAt.Format(time.RFC3339),
	}
	return result, nil
}
