package services

import (
	"example/models"
	"example/structs"
	"example/utils/encrypt"
	"example/utils/mysql_util"
	"time"
)

func Register(body structs.RegisterSchema) (structs.User, error) {
	hashPassword, err := encrypt.HashPassword(body.Password)
	if err != nil {
		return structs.User{}, err
	}
	newUser := models.User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Password:  hashPassword,
	}
	if err := mysql_util.DB.Create(&newUser).Error; err != nil {
		return structs.User{}, err
	}
	token := GenerateToken(newUser.ID)
	createdUser := structs.User{
		Token:      token,
		ID:         newUser.ID,
		FirstName:  newUser.FirstName,
		LastName:   newUser.LastName,
		Email:      newUser.Email,
		FacebookId: newUser.FacebookId,
		GoogleId:   newUser.GoogleId,
		Avatar:     newUser.Avatar,
		CreatedAt:  newUser.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  newUser.UpdatedAt.Format(time.RFC3339),
	}
	return createdUser, nil
}

func Login(body structs.LoginSchema) (structs.User, error) {
	user := models.User{}
	err := mysql_util.DB.Model(&models.User{}).Where("email = ?", body.Email).First(&user).Error
	if err != nil {
		return structs.User{}, err
	}
	if err = encrypt.CheckPassword(body.Password, user.Password); err != nil {
		return structs.User{}, err
	}
	token := GenerateToken(user.ID)
	result := structs.User{
		Token:      token,
		ID:         user.ID,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		FacebookId: user.FacebookId,
		GoogleId:   user.GoogleId,
		Avatar:     user.Avatar,
		CreatedAt:  user.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  user.UpdatedAt.Format(time.RFC3339),
	}
	return result, nil
}
