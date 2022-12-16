package repository

import (
	"nusatech/features/users"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.Repository {
	return &repoQuery{db: db}
}

func (rq *repoQuery) Insert(data users.UserCore) (users.UserCore, error) {
	var cnv User = FromCore(data)
	err := rq.db.Create(&cnv).Error
	if err != nil {
		log.Error("error on query register user", err.Error())
		return users.UserCore{}, err
	}

	res := ToCore(cnv)
	return res, nil
}

func (rq *repoQuery) GetLogin(input users.UserCore) (users.UserCore, error) {
	var data User
	err := rq.db.Preload("Balance").First(&data, "email = ?", input.Email).Error
	if err != nil {
		log.Error("error on query login user", err.Error())
		return users.UserCore{}, err
	}

	res := ToCore(data)
	return res, nil
}

func (rq *repoQuery) GetAll() ([]users.UserCore, error) {
	var data []User
	err := rq.db.Preload("Balance").Order("created_at desc").Find(&data).Error
	if err != nil {
		log.Error("error on query get all data user", err.Error())
		return nil, err
	}

	res := ToCoreArray(data)
	return res, nil
}

func (rq *repoQuery) Edit(data users.UserCore, id uint, email string) (users.UserCore, error) {
	var datas User
	var cnv User = FromCore(data)
	if err := rq.db.First(&datas, "id = ? AND email = ?", id, email).Error; err != nil {
		log.Error("error on query edit user", err.Error())
		return users.UserCore{}, err
	}

	if err := rq.db.Where("id = ?", id).Updates(&cnv).Error; err != nil {
		log.Error("error on query edit data user", err.Error())
		return users.UserCore{}, err
	}

	er := rq.db.Preload("Balance").First(&cnv, "id = ?", id).Error
	if er != nil {
		log.Error("error on get data after edit", er.Error())
		return users.UserCore{}, er
	}

	res := ToCore(cnv)
	return res, nil
}

func (rq *repoQuery) UpdateSendGrid(data users.UserCore) error {
	var cnv User = FromCore(data)
	if err := rq.db.Where("email = ?", cnv.Email).Updates(&cnv).Error; err != nil {
		return err
	}

	return nil
}
