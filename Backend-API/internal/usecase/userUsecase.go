package usecase

import (
	"farm-integrated-web3/dto"
	"farm-integrated-web3/internal/repository"
)

type UserUsecase interface {
	CreateProfile(input *dto.CreateProfileRequest) error
	UpdateProfile(input *dto.UpdateProfileRequest) error
	UpdateRole(input *dto.UpdateRoleRequest) error
	ChangePassword(input *dto.UserChangePasswordRequest) error
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo}
}

func (u *userUsecase) ChangePassword(input *dto.UserChangePasswordRequest) error {
	valid, err := u.userRepo.CheckUserExist(0, input.Email)
	if err != nil {
		return err
	}

	if !valid {
		return err
	}

	if err := u.userRepo.ChangePassword(input.Email, input.NewPassword); err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) CreateProfile(input *dto.CreateProfileRequest) error {
	valid, err := u.userRepo.CheckUserExist(input.UserId, "")
	if err != nil {
		return err
	}

	if !valid {
		return err
	}

	if err := u.userRepo.CreateProfile(input); err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) UpdateProfile(input *dto.UpdateProfileRequest) error {
	valid, err := u.userRepo.CheckUserExist(input.UserId, "")
	if err != nil {
		return err
	}

	if !valid {
		return err
	}

	if err := u.userRepo.UpdateProfile(input); err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) UpdateRole(input *dto.UpdateRoleRequest) error {
	if err := u.userRepo.UpdateRole(input); err != nil {
		return err
	}
	return nil
}
