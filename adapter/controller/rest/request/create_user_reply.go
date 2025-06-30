package requests

import (
	"opinion-tracker-service/boundry/dto"

	"github.com/google/uuid"
)

type CreateUserReply struct {
	Email   string    `json:"email"`
	UserKey uuid.UUID `json:"user_key"`
}

func FromDto(dto *dto.CreateUserReplyDTO) CreateUserReply {
	return CreateUserReply{
		Email:   dto.Email,
		UserKey: dto.UserKey,
	}
}
