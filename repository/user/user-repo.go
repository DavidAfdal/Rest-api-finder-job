package userrepo

import (
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(DB *gorm.DB) *UserRepo {
	return &UserRepo{DB: DB}
}

func (r *UserRepo) FindAll() {}

func (r *UserRepo) FindByPk() {}

func (r *UserRepo) Save() {}

func (r *UserRepo) Update() {}

func (r *UserRepo) Delete() {}