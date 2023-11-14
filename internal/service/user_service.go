package service

import (
	"go-with-kafka/internal/model"
	"go-with-kafka/internal/repository"
)

type UserService interface {
	Insert(request model.DefaultPayload[model.UserRequest]) (*model.DefaultPayload[model.UserResponse], error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {

	return userService{
		userRepo: userRepo,
	}
}

func (s userService) Insert(request model.DefaultPayload[model.UserRequest]) (*model.DefaultPayload[model.UserResponse], error) {

	data := request.Data

	entity := model.UserModel{
		Username: data.Username,
	}

	userResponse, err := s.userRepo.Create(entity)
	if err != nil {
		return nil, err
	}

	response := model.UserResponse{
		Id:       userResponse.Id,
		UserId:   userResponse.UserId,
		Username: userResponse.Username,
	}

	return &model.DefaultPayload[model.UserResponse]{
		Data: response}, nil
}
