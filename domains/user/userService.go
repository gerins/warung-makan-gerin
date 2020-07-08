package user

import (
	"database/sql"
	"errors"
	"warung_makan_gerin/utils/token"
	"warung_makan_gerin/utils/validation"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/validator.v2"
)

type UserService struct {
	db       *sql.DB
	UserRepo UserRepository
}

type UserServiceInterface interface {
	GetUsers() (*[]User, error)
	HandleUserLogin(userLogin User) (*UserToken, error)
	HandleRegisterNewUser(d User) (*User, error)
	HandleUPDATEUser(id string, data User) (*User, error)
	HandleDELETEUser(id string) (*User, error)
}

func NewUserService(db *sql.DB) UserServiceInterface {
	return UserService{db, NewUserRepo(db)}
}

func (s UserService) GetUsers() (*[]User, error) {
	User, err := s.UserRepo.HandleGETAllUser()
	if err != nil {
		return nil, err
	}

	return User, nil
}

func (s UserService) HandleUserLogin(userLogin User) (*UserToken, error) {
	if err := validator.Validate(userLogin); err != nil {
		return nil, err
	}

	User, err := s.UserRepo.HandleUserLogin(userLogin.Username, "A")
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(userLogin.Password))
	if err != nil {
		return nil, errors.New("Username atau Password salah")
	}

	getToken := token.GenerateToken(User.Username, 3600)
	userToken := UserToken{*User, getToken}
	return &userToken, nil
}

func (s UserService) HandleRegisterNewUser(d User) (*User, error) {
	if err := validator.Validate(d); err != nil {
		return nil, err
	}

	_, err := s.UserRepo.HandleUserLogin(d.Username, "A")
	if err == nil {
		return nil, errors.New("Username sudah digunakan")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(d.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	d.Password = string(hash)

	result, err := s.UserRepo.HandlePOSTUser(d)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s UserService) HandleUPDATEUser(id string, data User) (*User, error) {
	return nil, errors.New("Under Construction")
	if err := validator.Validate(data); err != nil {
		return nil, err
	}

	if err := validation.ValidateInputNumber(id); err != nil {
		return nil, err
	}

	result, err := s.UserRepo.HandleUPDATEUser(id, data)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s UserService) HandleDELETEUser(id string) (*User, error) {
	return nil, errors.New("Under Construction")
	if err := validation.ValidateInputNumber(id); err != nil {
		return nil, err
	}

	_, err := s.UserRepo.HandleUserLogin(id, "A")
	if err != nil {
		return nil, err
	}

	result, err := s.UserRepo.HandleDELETEUser(id)
	if err != nil {
		return result, err
	}
	return result, nil
}
