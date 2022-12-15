package service

import (
	"nusatech/features/users"
	"nusatech/utils/middlewares"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	qry users.Repository
}

func New(repo users.Repository) users.Service {
	return &userService{qry: repo}
}

func (us *userService) Create(data users.UserCore) (users.UserCore, error) {
	generate, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return users.UserCore{}, err
	}
	data.Password = string(generate)

	res, err := us.qry.Insert(data)
	if err != nil {
		return users.UserCore{}, err
	}

	return res, nil
}

func (us *userService) Login(input users.UserCore) (users.UserCore, error) {
	res, err := us.qry.GetLogin(input)
	if err != nil {
		return users.UserCore{}, err
	}

	er := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(input.Password))
	if er != nil {
		return users.UserCore{}, er
	}

	res.Token = middlewares.GenerateToken(res.ID)

	return res, nil
}

func (us *userService) ShowAll() ([]users.UserCore, error) {
	res, err := us.qry.GetAll()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (us *userService) Update(data users.UserCore, id uint) (users.UserCore, error) {
	if data.Password != "" {
		generate, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
		if err != nil {
			return users.UserCore{}, err
		}
		data.Password = string(generate)
	}

	res, err := us.qry.Edit(data, id)
	if err != nil {
		return users.UserCore{}, err
	}

	return res, nil
}
