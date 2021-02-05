package services

import (
	"github.com/tpmanc/gateway/models"
	"github.com/tpmanc/gateway/repositories"
	"github.com/tpmanc/gateway/requests"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceInterface interface {
	Login(r *requests.LoginRequest) *models.User
	Signup(r *requests.SignupRequest) *models.User
}

type userService struct {
	rep repositories.UserRepositoryInterface
}

func (s *userService) Login(r *requests.LoginRequest) *models.User {
	user := s.rep.GetUserByLogin(r.Login)
	if user == nil {
		return nil
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Пароль не совпадает!!
		return nil
	}

	return user
}

func (s *userService) Signup(r *requests.SignupRequest) *models.User {
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)

	model := &models.User{
		Username: r.Login,
		Password: string(passwordHash),
	}


	s.rep.Create(model)

	return model
}

func GetUserService(rep repositories.UserRepositoryInterface) UserServiceInterface {
	return &userService{
		rep: rep,
	}
}
