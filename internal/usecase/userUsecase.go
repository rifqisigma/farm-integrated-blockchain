package usecase

import (
	"farm-integrated-web3/dto"
	"farm-integrated-web3/internal/repository"
	"farm-integrated-web3/utils/helper"
)

type UserUsecase interface {
	CreateConsumerProfile(input *dto.CreateConsumerProfile) error
	CreateFarmerProfile(input *dto.CreateFarmerProfile) error
	CreateDistributorProfile(input *dto.CreateDistributorProfile) error
	CreateRetailerProfile(input *dto.CreateRetailerProfile) error

	UpdateConsumerProfile(input *dto.UpdateConsumerProfile) error
	UpdateFarmerProfile(input *dto.UpdateFarmerProfile) error
	UpdateDistributorProfile(input *dto.UpdateDistributorProfile) error
	UpdateRetailerProfile(input *dto.UpdateRetailerProfile) error

	ChangePassword(input *dto.UserNewPassword) error
	DeleteAccount(email string) error
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo}
}

func (u *userUsecase) ChangePassword(input *dto.UserNewPassword) error {
	valid, err := u.userRepo.CheckUserExist(0, input.Email)
	if err != nil {
		return err
	}

	if !valid {
		return err
	}
	if input.ConfirmNewPassword != input.NewPassword {
		return helper.ErrBadRequest
	}

	if err := u.userRepo.ChangePassword(input.Email, input.NewPassword); err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) DeleteAccount(email string) error {
	valid, err := u.userRepo.CheckUserExist(0, email)
	if err != nil {
		return err
	}

	if !valid {
		return err
	}

	if err := u.userRepo.DeleteAccount(email); err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) CreateConsumerProfile(input *dto.CreateConsumerProfile) error {
	valid, err := u.userRepo.CheckUserExist(input.UserId, "")
	if err != nil {
		return err
	}

	if !valid {
		return err
	}

	if err := u.userRepo.CreateConsumerProfile(input); err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) CreateFarmerProfile(input *dto.CreateFarmerProfile) error {
	valid, err := u.userRepo.CheckUserExist(input.UserId, "")
	if err != nil {
		return err
	}

	if !valid {
		return err
	}

	if err := u.userRepo.CreateFarmerProfile(input); err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) CreateDistributorProfile(input *dto.CreateDistributorProfile) error {
	valid, err := u.userRepo.CheckUserExist(input.UserId, "")
	if err != nil {
		return err
	}

	if !valid {
		return err
	}

	if err := u.userRepo.CreateDistributorProfile(input); err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) CreateRetailerProfile(input *dto.CreateRetailerProfile) error {
	valid, err := u.userRepo.CheckUserExist(input.UserId, "")
	if err != nil {
		return err
	}

	if !valid {
		return err
	}

	if err := u.userRepo.CreateRetailerProfile(input); err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) UpdateConsumerProfile(input *dto.UpdateConsumerProfile) error {
	valid, err := u.userRepo.CheckUserExist(input.UserId, "")
	if err != nil {
		return err
	}

	if !valid {
		return err
	}

	if err := u.userRepo.UpdateConsumerProfile(input); err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) UpdateFarmerProfile(input *dto.UpdateFarmerProfile) error {
	valid, err := u.userRepo.CheckUserExist(input.UserId, "")
	if err != nil {
		return err
	}

	if !valid {
		return err
	}

	if err := u.userRepo.UpdateFarmerProfile(input); err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) UpdateDistributorProfile(input *dto.UpdateDistributorProfile) error {
	valid, err := u.userRepo.CheckUserExist(input.UserId, "")
	if err != nil {
		return err
	}

	if !valid {
		return err
	}

	if err := u.userRepo.UpdateDistributorProfile(input); err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) UpdateRetailerProfile(input *dto.UpdateRetailerProfile) error {
	valid, err := u.userRepo.CheckUserExist(input.UserId, "")
	if err != nil {
		return err
	}

	if !valid {
		return err
	}

	if err := u.userRepo.UpdateRetailerProfile(input); err != nil {
		return err
	}

	return nil
}
