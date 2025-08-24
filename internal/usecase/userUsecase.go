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

	ChangePassword(input *dto.UserChangePassword) error
	DeleteAccount(useriD uint) error
	Logout(userId uint) error
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo}
}

func (u *userUsecase) ChangePassword(input *dto.UserChangePassword) error {
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

func (u *userUsecase) Logout(userId uint) error {
	valid, err := u.userRepo.CheckUserExist(userId, "")
	if err != nil {
		return err
	}

	if !valid {
		return helper.ErrUserNotFound
	}

	return u.userRepo.DeleteToken(userId)

}

func (u *userUsecase) DeleteAccount(userId uint) error {
	valid, err := u.userRepo.CheckUserExist(userId, "")
	if err != nil {
		return err
	}

	if !valid {
		return err
	}

	if err := u.userRepo.DeleteAccount(userId); err != nil {
		return err
	}

	if err := u.userRepo.DeleteToken(userId); err != nil {
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
