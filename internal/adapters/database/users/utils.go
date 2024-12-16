package users

import (
	entitiesUsers "server/internal/entities"
)

func mapUserDataToUserEntity(userData *userData) entitiesUsers.User {
	userEntity := entitiesUsers.User{
		UserID:    userData.UserID,
		FirstName: userData.FirstName,
		CreatedAt: userData.CreatedAt,
		UpdatedAt: userData.UpdatedAt,
		Password:  userData.Password,
		DeletedAt: nil,
	}

	if userData.EmailEncrypted != nil && len(*userData.EmailEncrypted) > 0 {
		userEntity.EmailEncrypted = userData.EmailEncrypted
	}

	if userData.EmailSearchable != nil && len(*userData.EmailSearchable) > 0 {
		userEntity.EmailSearchable = userData.EmailSearchable
	}

	if userData.PhoneEncrypted != nil && len(*userData.PhoneEncrypted) > 0 {
		userEntity.PhoneEncrypted = userData.PhoneEncrypted
	}

	if userData.PhoneSearchable != nil && len(*userData.PhoneSearchable) > 0 {
		userEntity.PhoneSearchable = userData.PhoneSearchable
	}

	return userEntity
}
