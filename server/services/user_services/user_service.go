package user_services

import (
	domain "skillsmart/homepage/src/domain/user"
	utils "skillsmart/homepage/src/utils/error"
)

// RETRIEVAL ------------------------

func GetUserByID(userId int64) (*domain.User, error) {
	if user, ok :=	domain.MockUsersById[userId]; ok {
		return user, nil
	}
	return nil, utils.ApplicationError.Info("User with given id does not exist")
}

func GetUserByEmail(email string) (*domain.User, error) {
	if user, ok := domain.MockUsersByEmail[email]; ok {
		return user, nil
	}
	return nil, utils.ApplicationError.Info("no user for given credentials found")
}
