package usecase

import (
	"context"
	"farm-integrated-web3/dto"
	"farm-integrated-web3/internal/repository"
)

type UserUsecase interface {
	Me(ctx context.Context, id uint) (*dto.GetUser, error)
	CreateProfile(ctx context.Context, input *dto.CreateProfileRequest) error
	UpdateProfile(ctx context.Context, input *dto.UpdateProfileRequest) error
	UpdateRole(ctx context.Context, input *dto.UpdateRoleRequest) error
	ChangePassword(ctx context.Context, input *dto.UserChangePasswordRequest) error
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo}
}

func (u *userUsecase) ChangePassword(ctx context.Context, input *dto.UserChangePasswordRequest) error {
	if err := u.userRepo.ChangePassword(ctx, input.UserId, input.Email, input.NewPassword); err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) CreateProfile(ctx context.Context, input *dto.CreateProfileRequest) error {
	if err := u.userRepo.CreateProfile(ctx, input); err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) UpdateProfile(ctx context.Context, input *dto.UpdateProfileRequest) error {
	if err := u.userRepo.UpdateProfile(ctx, input); err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) UpdateRole(ctx context.Context, input *dto.UpdateRoleRequest) error {
	if err := u.userRepo.UpdateRole(ctx, input); err != nil {
		return err
	}
	return nil
}

func (u *userUsecase) Me(ctx context.Context, id uint) (*dto.GetUser, error) {
	return u.userRepo.Me(ctx, id)
}
