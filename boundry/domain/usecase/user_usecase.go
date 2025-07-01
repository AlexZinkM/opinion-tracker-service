package usecase

import (
	"context"
	"opinion-tracker-service/boundry/dto"
)

type CreateUserUseCase interface {
	CreateUser(ctx context.Context, request dto.CreateUserRequestDTO) (*dto.CreateUserReplyDTO, error)
}
