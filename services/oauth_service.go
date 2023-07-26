package services

import (
	"errors"
	"example/models"
	"example/structs"
	"example/utils/mysql"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
	"gorm.io/gorm"
)

func LoginGoogleCallback(context *gin.Context) (structs.User, error) {
	googleUser, err := gothic.CompleteUserAuth(context.Writer, context.Request)
	if err != nil {
		return structs.User{}, err
	}
	existUser := models.User{}
	err = mysql.DB.Model(&models.User{}).Where("email = ?", googleUser.Email).First(&existUser).Error
	if err == nil {
		token := GenerateToken(existUser.ID)
		user := structs.User{
			Token:      token,
			ID:         existUser.ID,
			FirstName:  existUser.FirstName,
			LastName:   existUser.LastName,
			Email:      existUser.Email,
			FacebookID: existUser.FacebookID,
			GoogleID:   googleUser.UserID,
			Avatar:     existUser.Avatar,
			CreatedAt:  existUser.CreatedAt.Format(time.RFC3339),
			UpdatedAt:  existUser.UpdatedAt.Format(time.RFC3339),
		}
		if existUser.GoogleID == "" {
			go mysql.DB.Updates(models.User{
				ID:       existUser.ID,
				GoogleID: googleUser.UserID,
			})
		}
		if existUser.Avatar == "" {
			go mysql.DB.Updates(models.User{
				ID:     existUser.ID,
				Avatar: googleUser.AvatarURL,
			})
		}
		return user, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return structs.User{}, err
	}
	newUser := models.User{
		FirstName: googleUser.FirstName,
		LastName:  googleUser.LastName,
		Email:     googleUser.Email,
		GoogleID:  googleUser.UserID,
		Avatar:    googleUser.AvatarURL,
	}
	err = mysql.DB.Create(&newUser).Error
	if err != nil {
		return structs.User{}, err
	}
	token := GenerateToken(newUser.ID)
	createdUser := structs.User{
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

func LoginFacebookCallback(context *gin.Context) (structs.User, error) {
	facebookUser, err := gothic.CompleteUserAuth(context.Writer, context.Request)
	if err != nil {
		return structs.User{}, err
	}
	existUser := models.User{}
	err = mysql.DB.Model(&models.User{}).Where("facebook_id = ?", facebookUser.UserID).First(&existUser).Error
	if err == nil {
		token := GenerateToken(existUser.ID)
		user := structs.User{
			Token:      token,
			ID:         existUser.ID,
			FirstName:  existUser.FirstName,
			LastName:   existUser.LastName,
			Email:      existUser.Email,
			FacebookID: existUser.FacebookID,
			GoogleID:   existUser.GoogleID,
			Avatar:     existUser.Avatar,
			CreatedAt:  existUser.CreatedAt.Format(time.RFC3339),
			UpdatedAt:  existUser.UpdatedAt.Format(time.RFC3339),
		}
		if existUser.Avatar == "" {
			go mysql.DB.Updates(models.User{
				ID:     existUser.ID,
				Avatar: facebookUser.AvatarURL,
			})
		}
		return user, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return structs.User{}, err
	}
	newUser := models.User{
		FirstName:  facebookUser.FirstName,
		LastName:   facebookUser.LastName,
		Email:      facebookUser.Email,
		FacebookID: facebookUser.UserID,
		Avatar:     facebookUser.AvatarURL,
	}
	err = mysql.DB.Create(&newUser).Error
	if err != nil {
		return structs.User{}, err
	}
	token := GenerateToken(newUser.ID)
	createdUser := structs.User{
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
