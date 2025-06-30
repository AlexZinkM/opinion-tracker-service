package dto

import "github.com/google/uuid"

type CreateUserReplyDTO struct {
	Email   string
	UserKey uuid.UUID
}
