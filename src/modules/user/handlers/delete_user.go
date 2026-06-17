package handlers

import "banc-api/src/modules/user/domain/repositories"

// DeleteUserid is a handler wrapper for user deletion operations.
type DeleteUserid struct {
	repo repositories.UserRepository
}

func DeleteUserHandler(repo repositories.UserRepository) *DeleteUserid {
	if repo == nil {
		return nil
	}

	return &DeleteUserid{repo}
}
