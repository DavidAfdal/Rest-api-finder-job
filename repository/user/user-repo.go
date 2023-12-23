package userrepo

import (
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

type UserRepoMethod interface {
	FindAll()
	FindByPk()
	Save()
	Update()
	Delete()
}

func NewUserRepo(DB *gorm.DB) UserRepoMethod {
	return &UserRepo{DB: DB}
}

func (r *UserRepo) FindAll() {}

func (r *UserRepo) FindByPk() {}

func (r *UserRepo) Save() {}

func (r *UserRepo) Update() {}

func (r *UserRepo) Delete() {}