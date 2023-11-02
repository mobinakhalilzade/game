package userservice

import (
	"fmt"
	"go-cast/game/entity"
	"go-cast/game/pkg/phonenumber"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	Register(user entity.User) (entity.User, error)
}

type Service struct {
	repo Repository
}

type RegisterRequest struct {
	Name        string
	phoneNumber string
}

type RegisterResponse struct {
	User entity.User
}

func (s Service) register(request RegisterRequest) (RegisterResponse, error) {
	// TODO - we should verify phone number by verification code

	//validate phone number
	if !phonenumber.IsValid(request.phoneNumber) {
		return RegisterResponse{}, fmt.Errorf("phone number is not valid")
	}

	//check uniqueness of the phone number
	if isUnique, err := s.repo.IsPhoneNumberUnique(request.phoneNumber); err != nil || !isUnique {
		if err != nil {
			return RegisterResponse{}, fmt.Errorf("unexpected error:%w", err)
		}

		if !isUnique {
			return RegisterResponse{}, fmt.Errorf("phone number is not unique")
		}
	}

	//validate name
	if len(request.Name) < 3 {
		return RegisterResponse{}, fmt.Errorf("name length should be greater than 3")
	}

	//save in database
	user := entity.User{
		ID:          0,
		Name:        request.Name,
		PhoneNumber: request.phoneNumber,
	}
	createdUser, err := s.repo.Register(user)
	if err != nil {
		return RegisterResponse{}, fmt.Errorf("unexpected error:%w", err)
	}

	//return created user
	return RegisterResponse{User: createdUser}, nil
}
