package Utils

import (
	"web-api/src/Models/DatabaseModels"
	"web-api/src/Models/ViewModels"
)

func UserDbToViewModel(user DatabaseModels.User) ViewModels.User {
	return ViewModels.User{
		UserID:          user.ID,
		Username:        user.Username,
		Email:           user.Email,
		ProfileImageUrl: user.ProfileImageUrl,
	}
}
