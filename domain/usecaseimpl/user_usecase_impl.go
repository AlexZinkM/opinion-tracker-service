package usecase

import (
	"context"

	"opinion-tracker-service/boundry/dto"

	"github.com/google/uuid"
)

type CreateUserUseCaseImpl struct{}

func NewCreateUserUseCaseImpl() *CreateUserUseCaseImpl {
	return &CreateUserUseCaseImpl{}
}

func (u *CreateUserUseCaseImpl) CreateUser(ctx context.Context, request dto.CreateUserRequestDTO) (*dto.CreateUserReplyDTO, error) {

	return &dto.CreateUserReplyDTO{
		Email:   "example@example.com",
		UserKey: uuid.New(),
	}, nil
}
